package main

import (
    "net/http"
    "github.com/WilliamViega/API-students-/db"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    db.Init()
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/students", listStudents)
    e.POST("/students", createStudent)
    e.GET("/students/:id", getStudent)
    e.PUT("/students/:id", updateStudent)
    e.DELETE("/students/:id", deleteStudent)

    e.Logger.Fatal(e.Start(":8081"))
}

func listStudents(c echo.Context) error {
    students, err := db.GetAllStudents()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, students)
} 


func createStudent(c echo.Context) error {
    student := db.Student{}
    if err := c.Bind(&student); err != nil {
        return err
    }

    db.AddStudent(student)
    return c.String(http.StatusOK, "Student created")
}

func getStudent(c echo.Context) error {
    id := c.Param("id")
    student, err := db.GetStudentByID(id)
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Estudante não encontrado"})
    }
    return c.JSON(http.StatusOK, student)
}

func updateStudent(c echo.Context) error {
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

func deleteStudent(c echo.Context) error {
    id := c.Param("id")
    err := db.DeleteStudent(id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.String(http.StatusOK, "Estudante deletado")
}
 