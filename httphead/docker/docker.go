package docker

import (
    "bytes"
    "fmt"
    "github.com/antonlindstrom/httpheadgo/httphead/config"
    "io/ioutil"
    "net/http"
)

// Boot container (creates container and starts it)
func BootContainer(image string) {
    fmt.Printf("Creating container (image: %s)\n", image)
    id := CreateContainer(image)

    fmt.Printf("Starting container (id: %s)\n", id)
    StartContainer(id)
}

// Start container
func StartContainer(id string) ([]byte, error) {
    path    := fmt.Sprintf("/containers/%s/start", id)
    payload := "{}"

    body,err := PostRequest(path,payload)

    if err != nil {
        fmt.Printf("Error: %s\n", err)
        return nil, err
    }

    return body, err
}

// Create a container in docker
func CreateContainer(image string) (string) {
    path := "/containers/create"
    img  := fmt.Sprintf("{\"Image\":\"%s\"}", image)

    body,_  := PostRequest(path, img)
    data    := config.ParseJson(body)

    return data["Id"]
}

// Post Request and return bytes
func PostRequest(path, postData string) ([]byte, error) {
    buf  := bytes.NewBufferString(postData)
    resp, err := http.Post(buildUrl(path), "application/json", buf)

    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}

// Routines not exported below

// Build Url from parts
func buildUrl(path string) (string) {
    base := dockerBaseUrl()

    return fmt.Sprintf("%s%s", base, path)
}

// Docker Base Url
func dockerBaseUrl() (string) {
    c := config.GetConfig()
    return c["DockerUrl"]
}
