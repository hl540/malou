package web_server

import (
	v1 "github.com/hl540/malou/proto/v1"
)

type WebServer struct {
	v1.UnimplementedMalouWebServer
}
