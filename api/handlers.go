package api

import (
    "net/http"
    "strconv"

    "github.com/labstack/echo/v4"
    "github.com/WilliamViega/API-students/db"
)

// ListStudents godoc
// @Summary Lista todos os estudantes
// @Description Retorna todos os estudantes cadastrados
// @Tags estudantes
// @Produce json
// @Success 200 {array} db.Student
// @Router /students [get]
func (api *API) ListStudents(c echo.Context) error {
    students, err := db.GetAllStudents()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, students)
}

// CreateStudent godoc
// @Summary Cria um novo estudante
// @Description Adiciona um estudante ao banco de dados
// @Tags estudantes
// @Accept json
// @Produce json
// @Param student body db.Student true "Dados do estudante"
// @Success 201 {object} map[string]string
// @Router /students [post]
func (api *API) CreateStudent(c echo.Context) error {
    student := db.Student{}
    if err := c.Bind(&student); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados inválidos"})
    }

    db.AddStudent(student)
    return c.JSON(http.StatusCreated, map[string]string{"message": "Estudante criado com sucesso"})
}

// GetStudent godoc
// @Summary Busca um estudante por ID
// @Description Retorna os dados de um estudante específico
// @Tags estudantes
// @Produce json
// @Param id path int true "ID do estudante"
// @Success 200 {object} db.Student
// @Failure 404 {object} map[string]string
// @Router /students/{id} [get]
func (api *API) GetStudent(c echo.Context) error {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
    }

    student, err := db.GetStudentByID(id)
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Estudante não encontrado"})
    }
    return c.JSON(http.StatusOK, student)
}

// UpdateStudent godoc
// @Summary Atualiza o nome de um estudante
// @Description Modifica o nome de um estudante existente
// @Tags estudantes
// @Accept json
// @Produce json
// @Param id path int true "ID do estudante"
// @Param student body db.Student true "Novo nome do estudante"
// @Success 200 {object} map[string]string
// @Router /students/{id} [put]
func (api *API) UpdateStudent(c echo.Context) error {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
    }

    var s db.Student
    if err := c.Bind(&s); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados inválidos"})
    }

    err = db.UpdateStudent(id, s.Nome)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, map[string]string{"message": "Estudante atualizado"})
}

// DeleteStudent godoc
// @Summary Remove um estudante
// @Description Deleta um estudante do banco de dados
// @Tags estudantes
// @Produce json
// @Param id path int true "ID do estudante"
// @Success 200 {object} map[string]string
// @Router /students/{id} [delete]
func (api *API) DeleteStudent(c echo.Context) error {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
    }

    err = db.DeleteStudent(id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, map[string]string{"message": "Estudante deletado com sucesso"})
}
