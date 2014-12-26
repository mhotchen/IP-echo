package main

import "net/http"
import "os"
import "strings"
import "fmt"

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    http.HandleFunc("/", IpResponse)
    http.ListenAndServe(":" + port, nil)
}

func IpResponse(response http.ResponseWriter, request *http.Request) {
    ip := request.Header.Get("X-Forwarded-For")
    if ip == "" {
        ip = request.RemoteAddr
    }

    colonPosition := strings.Index(ip, ":")
    if colonPosition > -1 {
        ip = ip[:colonPosition]
    }

    commaPosition := strings.Index(ip, ",")
    if commaPosition > -1 {
        ip = ip[:commaPosition]
    }

    fmt.Println(ip)
    response.Write([]byte(ip))
}
