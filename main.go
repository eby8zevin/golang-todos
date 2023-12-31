package main

import (
	"github.com/eby8zevin/golang-todos/controller"
	"github.com/eby8zevin/golang-todos/database"
	"github.com/eby8zevin/golang-todos/middleware"
	"github.com/labstack/echo"
)

func main() {
	db := database.InitDb()
	defer db.Close()
	
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.Use(middleware.AuthMiddleware)

	controller.GetAllTodosController(e, db)
	controller.CreateTodoController(e, db)
	controller.DeleteTodoController(e, db)
	controller.UpdateTodoController(e, db)
	controller.CheckTodoController(e, db)

	controller.RegisterUserController(e, db)
	controller.LoginUserController(e, db)
	controller.CreateScopeController(e, db)
	controller.DeleteScopeController(e, db)
	controller.AssignScopeController(e, db)

	e.Start(":8080")
}