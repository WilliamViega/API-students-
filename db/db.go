package db

import (
    "database/sql"
    "log"
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
    // Garante que a pasta exista
    err := os.MkdirAll("students", os.ModePerm)
    if err != nil {
        log.Fatal("Erro ao criar pasta:", err)
    }

    DB, err = sql.Open("sqlite", "students/students.db") 
    if err != nil {
        log.Fatal("Erro ao abrir o banco:", err)
    }

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
        log.Fatal("Erro ao criar tabela:", err)
    }
}

func AddStudent(s Student) {
    _, err := DB.Exec("INSERT INTO students (nome, cpf, email, idade) VALUES (?, ?, ?, ?)",
        s.Nome, s.CPF, s.Email, s.Idade)
    if err != nil {
        log.Println("Erro ao adicionar estudante:", err)
    }
}

func GetAllStudents() ([]Student, error) {
    rows, err := DB.Query("SELECT id, nome, cpf, email, idade FROM students")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var students []Student
    for rows.Next() {
        var s Student
        if err := rows.Scan(&s.ID, &s.Nome, &s.CPF, &s.Email, &s.Idade); err != nil {
            return nil, err
        }
        students = append(students, s)
    }
    return students, nil
}

func GetStudentByID(id string) (Student, error) {
    var s Student
    err := DB.QueryRow("SELECT id, nome, cpf, email, idade FROM students WHERE id = ?", id).
        Scan(&s.ID, &s.Nome, &s.CPF, &s.Email, &s.Idade)
    return s, err
}

func UpdateStudent(id string, nome string) error {
    _, err := DB.Exec("UPDATE students SET nome = ? WHERE id = ?", nome, id)
    return err
}

func DeleteStudent(id string) error {
    _, err := DB.Exec("DELETE FROM students WHERE id = ?", id)
    return err
}
