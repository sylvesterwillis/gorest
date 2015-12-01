package main

import (
    "gorest/gorest"
    "fmt"
    "encoding/json"
)

func main() {
    var headers = make(map[string]string)
    var url = "http://jsonplaceholder.typicode.com/posts/1"

    headers["Content-Type"] = "application/json"

    jsonData, err := json.Marshal(gorest.Get(headers, url))

    if err != nil {
        panic(err)
    }

    fmt.Println(string(jsonData))
}