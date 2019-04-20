package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	scfgo "github.com/serverlessplus/go"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

const (
	port = 1216
)

var handler *scfgo.Handler

func init() {
	// start your server
	r := gin.Default()
	r.GET("/go-gin-example", func(c *gin.Context) {
		c.Data(200, "text/html", []byte("hello world"))
	})
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", scfgo.Host, port))
	if err != nil {
		fmt.Printf("failed to listen on port %d: %v\n", port, err)
		// panic to force the runtime to restart
		panic(err)
	}
	go http.Serve(l, r)

	// setup handler
	types := make(map[string]struct{})
	types["image/png"] = struct{}{}
	handler = scfgo.NewHandler(port).WithBinaryMIMETypes(types)
}

func entry(ctx context.Context, req *scfgo.APIGatewayRequest) (*scfgo.APIGatewayResponse, error) {
	return handler.Handle(ctx, req)
}

func main() {
	cloudfunction.Start(entry)
}
