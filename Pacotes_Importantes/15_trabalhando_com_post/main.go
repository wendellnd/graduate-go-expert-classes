package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	c := http.Client{
		Timeout: time.Second,
	}

	reqBody := bytes.NewBuffer([]byte(`{"name":"Wendell"}`))

	resp, err := c.Post("https://google.com", "application/json", reqBody)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.CopyBuffer(os.Stdin, resp.Body, nil)
}
