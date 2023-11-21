package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
	"github.com/ZubairARooghwall/GoVueracleAlchemy/routes"
	
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	_ "github.com/godror/godror"
)

func main() {
	db, err := sql.Open("godror", `user="C##USERNAME" password="Login123" connectString="dbhost:1521/orclpdb1"`)
	if err != nil {
		log.Fatal("Error connecting to the database: %v", err)
	}
	defer db.Close()

	router := gin.Default()

	router.Use(cors.Default())

	-, currentFile, -, - := runtime.Caller(0)
	staticPath := filepath.Join(filepath.Dir(currentFile), "static")
	router.Static("/", staticPath)

	userRepo := repository.NewUserRepository(db)
	userController := controllers.NewUserController(userRepo)

	initRoutes(router, userController)

	router.Run(":8080")
}

func initRouter(router *gin.Engine, userController *controller.UserController) {
	router.GET("/users", userController.ListAllUsers)

	router.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join("static", "index.html"))
	})
}
