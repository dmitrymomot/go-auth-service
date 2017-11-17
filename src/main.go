package main

import (
	"log"
	"net/http"
	"os"

	limit "github.com/aviddiviner/gin-limit"
	xss "github.com/dvwright/xss-mw"
	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/acme/autocert"
)

// APP_ENV application environment
// APP_HOST application host name
const (
	AppEnv          string = "APP_ENV"
	AppHost         string = "APP_HOST"
	AppPort         string = "APP_PORT"
	CacheDir        string = "APP_CACHE_DIR"
	DBConnectString string = "DB_CONN"
)

// Environment options
const (
	DevEnv  string = "development"
	ProdEnv string = "production"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(limit.MaxAllowed(10))
	var xssMdlwr xss.XssMw
	r.Use(xssMdlwr.RemoveXss())
	r.Use(nice.Recovery(recoveryHandler))

	db, err := gorm.Open("mysql", os.Getenv(DBConnectString))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	if db.HasTable(&User{}) == false {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&User{})
	}
	if db.HasTable(&Role{}) == false {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Role{})
	}
	db.AutoMigrate(&User{}, &Role{})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": http.StatusText(http.StatusNotFound)})
	})

	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{"error": http.StatusText(http.StatusNotImplemented)})
	})

	if os.Getenv(AppEnv) == ProdEnv {
		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(os.Getenv(AppHost)),
			Cache:      autocert.DirCache(CacheDir),
		}
		log.Fatal(autotls.RunWithManager(r, &m))
	} else {
		r.Run(os.Getenv(AppPort))
	}
}

// recoveryHandler creates json response with error
func recoveryHandler(c *gin.Context, err interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err})
}
