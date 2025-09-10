package db

import (
    "database/sql"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
    "os"

    _ "modernc.org/sqlite"
)

var DB *sql.DB

type Student struct {
    ID    int    `json:"id"`
    Nome  string `json:"nome"`
    CPF   string `json:"cpf"`
    Email string `json:"email"`
    Idade int    `json:"idade"`
}

func Init() {
    // Configura saída legível no terminal
    log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

    // Garante que a pasta exista
    err := os.MkdirAll("students", os.ModePerm)
    if err != nil {
        log.Fatal().Err(err).Msg("Erro ao criar pasta")
    }

    DB, err = sql.Open("sqlite", "students/students.db")
    if err != nil {
        log.Fatal().Err(err).Msg("Erro ao abrir o banco")
    }

    log.Info().Msg("Banco de dados SQLite inicializado com sucesso")

    // Cria a tabela se não existir
    _, err = DB.Exec(`
        CREATE TABLE IF NOT EXISTS students (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            nome TEXT NOT NULL,
            cpf TEXT NOT NULL,
            email TEXT NOT NULL,
            idade INTEGER NOT NULL
        );
    `)
    if err != nil {
        log.Fatal().Err(err).Msg("Erro ao criar tabela")
    }

    log.Info().Msg("Tabela 'students' verificada/criada com sucesso")
}

func AddStudent(s Student) {
    _, err := DB.Exec("INSERT INTO students (nome, cpf, email, idade) VALUES (?, ?, ?, ?)",
        s.Nome, s.CPF, s.Email, s.Idade)
    if err != nil {
        log.Error().Err(err).Msg("Erro ao adicionar estudante")
    } else {
        log.Info().Msgf("Estudante adicionado: %s", s.Nome)
    }
}

func GetAllStudents() ([]Student, error) {
    rows, err := DB.Query("SELECT id, nome, cpf, email, idade FROM students")
    if err != nil {
        log.Error().Err(err).Msg("Erro ao buscar estudantes")
        return nil, err
    }
    defer rows.Close()

    var students []Student
    for rows.Next() {
        var s Student
        if err := rows.Scan(&s.ID, &s.Nome, &s.CPF, &s.Email, &s.Idade); err != nil {
            log.Error().Err(err).Msg("Erro ao escanear estudante")
            return nil, err
        }
        students = append(students, s)
    }
    log.Info().Msgf("Total de estudantes encontrados: %d", len(students))
    return students, nil
}

func GetStudentByID(id string) (Student, error) {
    var s Student
    err := DB.QueryRow("SELECT id, nome, cpf, email, idade FROM students WHERE id = ?", id).
        Scan(&s.ID, &s.Nome, &s.CPF, &s.Email, &s.Idade)
    if err != nil {
        log.Error().Err(err).Msgf("Estudante com ID %s não encontrado", id)
    }
    return s, err
}

func UpdateStudent(id string, nome string) error {
    _, err := DB.Exec("UPDATE students SET nome = ? WHERE id = ?", nome, id)
    if err != nil {
        log.Error().Err(err).Msgf("Erro ao atualizar estudante com ID %s", id)
    } else {
        log.Info().Msgf("Estudante com ID %s atualizado para nome: %s", id, nome)
    }
    return err
}

func DeleteStudent(id string) error {
    _, err := DB.Exec("DELETE FROM students WHERE id = ?", id)
    if err != nil {
        log.Error().Err(err).Msgf("Erro ao deletar estudante com ID %s", id)
    } else {
        log.Info().Msgf("Estudante com ID %s deletado com sucesso", id)
    }
    return err
}
  