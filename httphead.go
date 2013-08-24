package main

import (
    "fmt"
    "github.com/antonlindstrom/httpheadgo/httphead"
    "net/http"
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
    name, url := config.GetConfig()

    fmt.Printf("Checking %s (%s)\n", name, url)
    protocol, statusCode := HeadCheck(url)

    fmt.Printf("Status: %s %d\n", protocol, statusCode)

    if statusCode >= 200 && statusCode < 300 {
        fmt.Printf("OK - %s\n", name)
    } else {
        fmt.Printf("FAIL - %s\n", name)
    }
}
