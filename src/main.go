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
// APP_PORT application internal port
// APP_CACHE_DIR application cache directory, mainly using for let's encrypt certificates
// DB_CONN database connection string, must contain full connection data like username, password. db name, etc
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

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/login")
	})

	r.GET("/login", LoginPageHandler)
	r.GET("/register", RegisterPageHandler)
	r.GET("/forgot-password", ForgotPasswordPageHandler)

	r.POST("/login", LoginHandler)
	r.POST("/register", RegisterHandler)
	r.POST("/forgot-password", ForgotPasswordHandler)

	API := r.Group("/API")
	{
		API.POST("/login", APILoginHandler)
		API.POST("/register", APIRegisterHandler)
		API.POST("/forgot-password", APIForgotPasswordHandler)
		API.POST("/refresh-token", APIRefreshTokenHandler)
	}

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

// CheckErr is errors handler
func checkErr(err error, c *gin.Context) {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// LoginPageHandler login page
func LoginPageHandler(c *gin.Context) {

}

// RegisterPageHandler registration page
func RegisterPageHandler(c *gin.Context) {

}

// ForgotPasswordPageHandler reset password page, first step
func ForgotPasswordPageHandler(c *gin.Context) {

}

// ResetPasswordPageHandler reset password page, second step
func ResetPasswordPageHandler(c *gin.Context) {

}

// LoginHandler login user via web interface
func LoginHandler(c *gin.Context) {

}

// RegisterHandler register new user via web interface
func RegisterHandler(c *gin.Context) {

}

// ForgotPasswordHandler reset fogotten password via web interface, step one
func ForgotPasswordHandler(c *gin.Context) {

}

// ResetPasswordHandler reset fogotten password via web interface, step one
func ResetPasswordHandler(c *gin.Context) {

}

// APILoginHandler login user via API
func APILoginHandler(c *gin.Context) {

}

// APIRegisterHandler register new user via API
func APIRegisterHandler(c *gin.Context) {

}

// APIForgotPasswordHandler reset fogotten password via API, first step
func APIForgotPasswordHandler(c *gin.Context) {

}

// APIResetPasswordHandler reset fogotten password via API, step two
func APIResetPasswordHandler(c *gin.Context) {

}

// APIRefreshTokenHandler refresh API token
func APIRefreshTokenHandler(c *gin.Context) {

}
