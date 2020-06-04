package main

import (
    "fmt"
    "net/http"
    "strconv"
    "os"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    nice := solution()
    fmt.Fprintf(w, "Nice: ", nice)
}

func solution() int {
    a, err := strconv.Atoi(os.Getenv(""))
    b, err := strconv.Atoi(os.Getenv(""))
    c, err := strconv.Atoi(os.Getenv(""))
    d := a+b+c
    if err == nil {
        return(d)
    } else {
        return(1)
    }
}
