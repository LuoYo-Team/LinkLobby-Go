package server

import (
    "fmt"
    "net/http"
)

// helloHandler 是处理 /hello 路径请求的处理函数
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "你好，世界！")
}

func main() {
    // 将 /hello 路径的请求注册到 helloHandler 处理函数
    http.HandleFunc("/hello", helloHandler)

    // 在端口 8080 启动 HTTP 服务器
    fmt.Println("服务器启动于端口 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("服务器启动错误:", err)
    }
}