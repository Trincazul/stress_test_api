package main

import (
    "fmt"
    "net/http"
    "time"
)

func testServer(url string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()
    elapsed := time.Since(start)
    fmt.Printf("Tempo de resposta: %s\n", elapsed)
}

func main() {
    url := "http://localhost:8080"
    for i := 0; i < 10; i++ {
        testServer(url)
    }

}
