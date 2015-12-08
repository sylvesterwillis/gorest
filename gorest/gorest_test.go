package gorest

import "testing"

func TestGet(t *testing.T) {
    var headers = make(map[string]string)
    var url = "http://jsonplaceholder.typicode.com/posts/1"

    headers["Content-Type"] = "application/json"

    _, err := Get(headers, url)

    if err != nil {
        t.Error(err)
    }
}
