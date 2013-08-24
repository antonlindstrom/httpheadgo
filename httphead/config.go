package config

import (
    "encoding/json"
    "io/ioutil"
)

// Load File and return content
func LoadFile(path string) ([]byte) {
    bytes, err := ioutil.ReadFile(path)

    if err != nil {
        return nil
    }

    return bytes
}

// Parse JSON content and return map
func ParseJson(file []byte) (map[string]string) {
    data := make(map[string]string)
    err  := json.Unmarshal(file, &data)

    if err != nil {
        return nil
    }

    return data
}

// Get Config and return name and url
func GetConfig() (string, string) {
    file := LoadFile("./app.json")
    data := ParseJson(file)

    return data["Name"], data["Url"]
}
