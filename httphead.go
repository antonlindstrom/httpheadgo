package main

import (
    "net/http"
    "fmt"
)

// Run HEAD request against url, fail will result
// in code 599
func HeadCheck(url string) (string, int) {
    resp, err := http.Head(url)

    if err != nil {
        return "HTTP/1.1", 599
    }

    return resp.Proto, resp.StatusCode
}

// Main function
func main() {
    url := "http://example.com"

    protocol, statusCode := HeadCheck(url)

    fmt.Printf("Host: %s\n", url)
    fmt.Printf("Status: %s %d\n", protocol, statusCode)

    if statusCode >= 200 && statusCode < 300 {
        fmt.Printf("%s OK\n", url)
    } else {
        fmt.Printf("%s FAIL\n", url)
    }
}
