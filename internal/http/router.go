package http

import (
    "net/http"
)

func NewRouter() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/", homeHandler)
    return mux
}
