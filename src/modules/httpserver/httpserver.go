package httpserver

import (
	"fmt"
	"LinkLobby-Go/src/modules/response"
	"net/http"
	"strings"
)

func handler(resp http.ResponseWriter, req *http.Request) {
	if !strings.Contains(req.Header["User-Agent"],"PCL2") {
		resp.WriteHeader(response.Forbidden)
		resp.Header().Set("Content-Type", "application/json")
		resp.Header().Set("X-Linklobby-Api-Location")
	}
	switch req.URL.Path {
	case "/api":

	case "/api/auth/login":
		go Authization(req, resp)
	case "/api/lobbies/create":
		go createLobby(req, resp)
	case "/api/v1":
	}

}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("HTTP Server bind at 20643")
	fmt.Println("Socket Server bind at 20644")
	if err := http.ListenAndServe(":20643", nil); err != nil {
		fmt.Println("服务器启动错误:", err)
	}
}
