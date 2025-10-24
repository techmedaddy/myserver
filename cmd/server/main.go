package main

import (
    "log"
    "net/http"

    apphttp "github.com/techmedaddy/myserver/internal/http"
    "github.com/techmedaddy/myserver/pkg/utils"
)

func main() {
    logger := utils.NewLogger()
    router := apphttp.NewRouter()

    logger.Info("Server started on :8083")

    if err := http.ListenAndServe(":8083", router); err != nil {
        logger.Error(err)
        log.Fatal(err)
    }
}
