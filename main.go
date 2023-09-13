package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/harisaginting/gwyn/frontend"
	repo "github.com/harisaginting/gwyn/repositories"
	router "github.com/harisaginting/gwyn/routers"
	"github.com/harisaginting/gwyn/utils/helper"
	"github.com/harisaginting/gwyn/utils/log"
)

const projectDirName = "gwyn" //  project name

func main() {
	runGin()
}

func runGin() {
	helper.LoadEnv(projectDirName)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	log.Info(ctx, fmt.Sprintf("%s START", projectDirName))

	port := helper.MustGetEnv("PORT")
	app := gin.New()
	ginConfig(ctx, app)

	// route
	app.GET("/healthcheck", healthcheck)
	app.NoRoute(lostInSpce)
	// FRONTEND
	app.Static("/static", "./frontend/asset")
	// template
	app.LoadHTMLGlob("./frontend/page/*.html")

	plain := app.Group("")

	// API
	router.Api(plain)
	// PAGE
	frontend.Page(plain)

	// handling server gracefully shutdown
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: app,
	}
	// Initializing the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Info(ctx, fmt.Sprintf("listen: %s", port))
		}
	}()
	// Listen for the interrupt signal.
	<-ctx.Done()
	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Warn(ctx, "shutting down gracefully, press Ctrl+C again to force 🔴")
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Warn(ctx, "Server forced to shutdown 🔴: ", err)
	}
	log.Warn(ctx, "Server shutdown 🔴")
}

func ginConfig(ctx context.Context, app *gin.Engine) {
	app.Use(gin.Logger())

	// DB CONNECTION
	db := repo.Connection()
	// repo.Migration(db)
	app.Use(repo.Inject(db))

	// get default url request
	app.UseRawPath = true
	app.UnescapePathValues = true
	// cors configuration
	config := cors.DefaultConfig()
	config.AddAllowHeaders("Authorization", "x-source")
	config.AddAllowHeaders("X-Frame-Options", "*")
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"OPTIONS", "PUT", "POST", "GET", "DELETE"}
	app.Use(cors.New(config))

	// error recorvery
	app.Use(gin.CustomRecovery(panicHandler))
}

func lostInSpce(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":        404,
		"data":          nil,
		"error_message": "No Route Found",
	})
}

func healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"port":         os.Getenv("PORT"),
		"service_name": os.Getenv("APP_NAME"),
		"mode":         os.Getenv("MODE"),
		"db_host":      os.Getenv("DB_HOST"),
	})
}

// Custom Recovery Panic Error
func panicHandler(c *gin.Context, err interface{}) {
	ctx := c.Request.Context()
	newerr := helper.ForceError(err)
	log.Error(ctx, newerr, "Panic Error 🔴")
	c.JSON(500, gin.H{
		"status":        500,
		"error_message": err,
	})
}

// test input from cmd
// wont work with hot reload
func inputTest() {
	fmt.Println("Enter Input 1 : ")
	var command1 string
	// Taking input from console
	fmt.Scanln(&command1)

	fmt.Println("Enter Input 2 : ")
	var command2 string
	fmt.Scanln(&command2)

	samplePrintInput(command1, command2)
}
func samplePrintInput(command1, command2 string) {
	fmt.Printf("command1:%s, command2:%s\n", command1, command2)
}
