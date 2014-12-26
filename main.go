package main

import "net/http"
import "os"

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    http.ListenAndServe(":" + port, http.FileServer(http.Dir(".")))
}
