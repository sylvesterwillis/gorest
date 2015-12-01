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

/*
Parameters: 
    headers: A map of header fields which map to their respective header value.
    url: The url to the endpoint of course.

Returns:
    A map of strings to interfaces or a nil, error tuple. I chose to use interfaces 
    in this case because I would like to keep the library as generic as possible 
    so it's best to not define a JSON stucture.
*/

func Get(headers map[string]string, url string) (interface{}) {
    var responseBody interface{}

    client := &http.Client{}

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    for key, val := range headers {
        req.Header.Add(key, val)
    }

    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()
    rawBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    return json.Unmarshal(rawBody, &responseBody)
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