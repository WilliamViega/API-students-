package api

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/WilliamViega/API-students/db"
)

type API struct {
    Echo *echo.Echo
}

func NewServer() *API {
    db.Init()
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    return &API{
        Echo: e,
    }
}

func (api *API) ConfigureRoute() {
    api.Echo.GET("/students", api.listStudents)
    api.Echo.POST("/students", api.createStudent)
    api.Echo.GET("/students/:id", api.getStudent)
    api.Echo.PUT("/students/:id", api.updateStudent)
    api.Echo.DELETE("/students/:id", api.deleteStudent)
}

func (api *API) Start() error {
    return api.Echo.Start(":8081")
}

func (api *API) listStudents(c echo.Context) error {
    students, err := db.GetAllStudents()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, students)
}

func (api *API) createStudent(c echo.Context) error {
    student := db.Student{}
    if err := c.Bind(&student); err != nil {
        return err
    }

    db.AddStudent(student)
    return c.String(http.StatusOK, "Student created")
}

func (api *API) getStudent(c echo.Context) error {
    id := c.Param("id")
    student, err := db.GetStudentByID(id)
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Estudante não encontrado"})
    }
    return c.JSON(http.StatusOK, student)
}

func (api *API) updateStudent(c echo.Context) error {
    id := c.Param("id")
    var s db.Student
    if err := c.Bind(&s); err != nil {
        return err
    }
    err := db.UpdateStudent(id, s.Nome)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.String(http.StatusOK, "Estudante atualizado")
}

func (api *API) deleteStudent(c echo.Context) error {
    id := c.Param("id")
    err := db.DeleteStudent(id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.String(http.StatusOK, "Estudante deletado")
}
