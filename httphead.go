package main

import (
    "fmt"
    "github.com/antonlindstrom/httpheadgo/httphead/config"
    "github.com/antonlindstrom/httpheadgo/httphead/docker"
    "log"
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

// Initiate a check
func InitiateCheck() int {
    c := config.GetConfig()

    log.Printf("Checking %s (%s)\n", c["Name"], c["Url"])
    protocol, statusCode := HeadCheck(c["Url"])

    log.Printf("Status: %s %d\n", protocol, statusCode)

    if statusCode >= 200 && statusCode < 400 {
        log.Printf("OK - %s\n", c["Name"])
    } else {
        log.Printf("FAIL - %s\n", c["Name"])
        if c["Docker"] == "enabled" {
            docker.BootContainer(c["DockerImage"])
        }
    }

		return statusCode
}

// Main function
func main() {
    log.Printf("Serving /check on port 8080\n")

    http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
        checkResult := InitiateCheck()

        fmt.Fprintf(w, "Status: %d", checkResult)
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}
