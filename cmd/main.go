package main

import (
	_ "github.com/lib/pq"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-comment/internal/app"
)

func main() {
	app.Run()
}
