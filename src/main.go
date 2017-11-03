package main

import (
    "log"

    "github.com/gin-gonic/autotls"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/acme/autocert"
)

const APP_ENV = "APP_ENV"
const APP_HOST = "APP_HOST"

const (
    DevEnv  string = "development"
    ProdEnv string = "production"
)

func main() {
    r := gin.Default()

    // Ping handler
    r.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })

    if os.Getenv(APP_ENV) == ProdEnv {
        m := autocert.Manager{
            Prompt:     autocert.AcceptTOS,
            HostPolicy: autocert.HostWhitelist(os.Getenv(APP_HOST)),
            Cache:      autocert.DirCache("/var/www/.cache"),
        }
        log.Fatal(autotls.RunWithManager(r, &m))
    } else {
        r.Run(":80")
    }
}
