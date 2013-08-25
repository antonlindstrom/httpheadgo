package main

import (
    "fmt"
    "github.com/antonlindstrom/httpheadgo/httphead/config"
    "github.com/antonlindstrom/httpheadgo/httphead/docker"
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
    c := config.GetConfig()

    fmt.Printf("Checking %s (%s)\n", c["Name"], c["Url"])
    protocol, statusCode := HeadCheck(c["Url"])

    fmt.Printf("Status: %s %d\n", protocol, statusCode)

    if statusCode >= 200 && statusCode < 300 {
        fmt.Printf("OK - %s\n", c["Name"])
    } else {
        fmt.Printf("FAIL - %s\n", c["Name"])
        if c["Docker"] == "enabled" {
            docker.BootContainer(c["DockerImage"])
        }
    }
}
