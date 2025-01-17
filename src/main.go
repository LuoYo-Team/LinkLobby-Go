package main

import (
	"LinkLobby-Go/src/modules/response"
	"fmt"
    "LinkLobby-Go/src/modules/base"
)

func main() {
	fmt.Println(response.NoContent)
    var datamap = make(map[string]any)
    datamap["user"] = "shimoranla"
    var data = base.GetString(datamap)
    fmt.Println(data)
}
