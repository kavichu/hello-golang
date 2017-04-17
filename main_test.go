package main

import (
    "net/http"
    // "net/http/httptest"
    "testing"
    "fmt"
)

func Test_main(t *testing.T) {
    resp, err := http.Get("http://web:8080/")
    if err != nil {
        t.Fatal(err)
    }
    fmt.Println(resp.StatusCode == 200)
}