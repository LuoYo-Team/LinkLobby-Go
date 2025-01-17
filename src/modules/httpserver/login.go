package httpserver

import (
	"fmt"
	"strings"
	"LinkLobby-Go/src/modules/response"
	"LinkLobby-Go/src/modules/token"
	"net/http"
)

func Authization(req Request,resp ResponseWriter) *string {
	if (len(req.body) == 19 && len(strings.Split(req.body, "-")) == 4) {
		return createToken()
	} else {
		return fmt.Sprintf(response.Forbidden)
	}
}
