package logger

import (
    "archive/zip"
    "bufio"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "sync"
    "time"
)

type Logger struct {
    mu            sync.Mutex
    file          *os.File
    writer        *bufio.Writer
    developerMode bool
}

// NewLogger 创建一个新的 Logger 实例，打开日志文件并设置缓冲写入器
func NewLogger(filePath string, developerMode bool) (*Logger, error) {
    file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return nil, err
    }

    return &Logger{
        file:          file,
        writer:        bufio.NewWriter(file),
        developerMode: developerMode,
    }, nil
}

// lock 锁定互斥锁
func (l *Logger) lock() {
    l.mu.Lock()
}

// unlock 解锁互斥锁
func (l *Logger) unlock() {
    l.mu.Unlock()
}

// write 写入日志条目到缓冲写入器
func (l *Logger) write(module string, level string, message string) {
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    logEntry := fmt.Sprintf("%s [%s] [%s]: %s\n", timestamp, level, module, message)
    l.writer.WriteString(logEntry)
    l.writer.Flush()
}

// Info 记录信息日志
func (l *Logger) Info(module string, message string) {
    l.lock()
    defer l.unlock()
    l.write(module, "INFO", message)
}

// Debug 记录调试日志
func (l *Logger) Debug(module string, message string) {
    if !l.developerMode {
        return
    }
    l.lock()
    defer l.unlock()
    l.write(module, "DEBUG", message)
}

// Warn 记录警告日志
func (l *Logger) Warn(module string, message string) {
    l.lock()
    defer l.unlock()
    l.write(module, "WARN", message)
}

// Error 记录错误日志
func (l *Logger) Error(module string, message string) {
    l.lock()
    defer l.unlock()
    l.write(module, "ERROR", message)
}

// Export 导出日志文件为 .zip 文件
func (l *Logger) Export(zipFilePath string) error {
    l.lock()
    defer l.unlock()

    // 刷新缓冲区到文件
    if err := l.writer.Flush(); err != nil {
        return err
    }

    // 创建 .zip 文件
    zipFile, err := os.Create(zipFilePath)
    if err != nil {
        return err
    }
    defer zipFile.Close()

    zipWriter := zip.NewWriter(zipFile)
    defer zipWriter.Close()

    // 添加日志文件到 .zip
    logFileWriter, err := zipWriter.Create(filepath.Base(l.file.Name()))
    if err != nil {
        return err
    }

    logFile, err := os.Open(l.file.Name())
    if err != nil {
        return err
    }
    defer logFile.Close()

    if _, err := io.Copy(logFileWriter, logFile); err != nil {
        return err
    }

    return nil
}

// Close 关闭日志文件
func (l *Logger) Close() error {
    l.lock()
    defer l.unlock()
    if err := l.writer.Flush(); err != nil {
        return err
    }
    return l.file.Close()
}

//
func (l *Logger) Crash(reason string) {
    if (developerMode){
        panic(reason)
    }else {
        return
    }
}

// GetLogger 获取 Logger 实例
func GetLogger(filePath string, developerMode bool) (*Logger, error) {
    return NewLogger(filePath, developerMode)
}