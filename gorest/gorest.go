package gorest
/*
    This library is meant to help provide common functionality to send/retrieve JSON
    data from a provided REST endpoint.
*/

import (
    "net/http"
    "io/ioutil"
    "encoding/json"
)

func panicError(err error) {
    if err != nil {
        panic(err)
    }
}

/*
Parameters: 
    headers: A map of header fields which map to their respective header value.
    url: The url to the endpoint of course.

Returns:
    A map of strings to interfaces. I chose to use interfaces in this case because 
    I would like to keep the library as generic as possible so it's best to not
    define a JSON stucture.
*/

func Get(headers map[string]string, url string) (interface{}) {
    var responseBody interface{}

    client := &http.Client{}

    req, err := http.NewRequest("GET", url, nil)
    panicError(err)

    for key, val := range headers {
        req.Header.Add(key, val)
    }

    resp, err := client.Do(req)
    panicError(err)

    defer resp.Body.Close()
    rawBody, err := ioutil.ReadAll(resp.Body)
    panicError(err)

    err = json.Unmarshal(rawBody, &responseBody)
    panicError(err)

    return responseBody
}

/*

Parameters: 
    headers: A map of header fields which map to their respective header value.
             The header map should contain a 'Application-Type' key with a
             corresponding value.
    url: The url to the endpoint of course.
    data: The data to be sent to the client

Returns:
    A map of strings to interfaces following the same reasoning as explaned within the get function.
*/
// func post(map[string]string headers, string url, []byte data) map[string]interface{} {
//     client := &http.Client{}

//     req, err := http.NewRequest("POST", url, data)

//     if err != nil {
//         panic()
//     }
// }