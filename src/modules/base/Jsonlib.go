package base

import (
	"encoding/json"
)

// GetJson parses a JSON string and returns a map and an error if any
func GetJson(content string) any {
	var data map[string]any
	if err := json.Unmarshal([]byte(content), &data); err != nil {
		return nil
	}
	return data
}

func GetString(data map[string]any) any {
    jsonData, err := json.Marshal(data)
    if err != nil {
        return nil
    }
    return string(jsonData)
}
