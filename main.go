package main

import (
	"log"
	"sadlyy/handler"
	"sadlyy/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/sdly?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	//input := user.LoginInput{
	//	Email: "tegarp00@gmail.com",
	//	Password: "password",
	//}

	//user, err := userService.Login(input)
	//if err != nil {
	//	fmt.Println("terjadi kesalahan")
	//	fmt.Println(err.Error())
	//}

	//fmt.Println(user.Email)
	//fmt.Println(user.Name)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)

	router.Run()

	// input dari user
	// handler, mapping input dari user -> struct input
	// service : melakukan mapping dari struct input ke struct User
	// repository
	// db
}
