package main

import (
	"os"
	"github.com/valyala/fasthttp"
)

func main() {
	var port string = os.Getenv("PORT")

	fs := &fasthttp.FS{
		Root:               "./docs",
		IndexNames:         []string{"index.html"},
		GenerateIndexPages: true,
		Compress:           true,
		AcceptByteRange:    false,
	}

	fasthttp.ListenAndServe(
		":" + port,
		fs.NewRequestHandler(),
	)
}