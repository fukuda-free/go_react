package main

import (
        "net/http"
)

// func main() {
//         http.ListenAndServe("", http.NotFoundHandler())
// }

// func main() {
//         http.ListenAndServe("", http.FileServer(http.Dir(".")))
// }

// func main() {
//         hh := func(w http.ResponseWriter, rq *http.Request) {
//                 w.Write([]byte("Hello, This is GO-server!!"))
//         }

//         http.HandleFunc("/hello", hh)

//         http.ListenAndServe("", nil)
// }


func main() {
        msg := `<html><body>
                <h1>Hello</h1>
                <p>This is GO-server!!</p>
                </body></html>`
        hh := func(w http.ResponseWriter, rq *http.Request) {
                w.Write([]byte(msg))
        }

        http.HandleFunc("/hello", hh)

        http.ListenAndServe("", nil)
}
