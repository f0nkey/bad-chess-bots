package main

import (
	"log"
	"net/http"
	"strings"
)

const dir = "./docs/."

func main() {
	fs := http.FileServer(http.Dir(dir))
	log.Print("Serving " + dir + " on http://localhost:3131")
	http.ListenAndServe(":3131", http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Cache-Control", "no-cache")
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("Content-Type", "application/wasm")
		}
		fs.ServeHTTP(resp, req)
	}))
}
