package api

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/WilliamViega/API-students/db"
   echoSwagger "github.com/swaggo/echo-swagger" 
   _ "github.com/WilliamViega/API-students/docs" 
   
)

type API struct {
    Echo *echo.Echo
}

// Inicializa o servidor e configura middlewares
func NewServer() *API {
    db.Init() // Inicializa o banco de dados

    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    return &API{Echo: e}
}

// Define todas as rotas da API
func (api *API) ConfigureRoute() {
    api.Echo.GET("/students", api.ListStudents)
    api.Echo.POST("/students", api.CreateStudent)
    api.Echo.GET("/students/:id", api.GetStudent)
    api.Echo.PUT("/students/:id", api.UpdateStudent)
    api.Echo.DELETE("/students/:id", api.DeleteStudent)
    api.Echo.GET("/swagger/*", echoSwagger.WrapHandler) 

}


// Inicia o servidor na porta 8081
func (api *API) Start() error {
    return api.Echo.Start(":8081")
}
 