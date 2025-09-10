<<<<<<< HEAD
# api-students
API to 'Golang do zero' course students 

Routes:
- GET /students - List all students
- POST /alunos - Creat students
- GET /students/:id - get infos from a specific student
- PUT /students/:id - Update student
- DELETE /students/:id - Delete student 

struct student 
- Name (string)
- Cpf (int)
- Email (string)
- Age (int)
- Active (bool)


=======
# API-students-
API- to ' Golang do zero ' course students
>>>>>>> d77e8bdc1bc3478afb1720a71013f5fb7520f580
# API Students

# 🧠 API Students

Uma API simples em Go usando o framework Echo para gerenciar estudantes. Ideal para estudos, testes de integração e como base para projetos maiores.

---

## 🚀 Como rodar

### Pré-requisitos

- Go instalado ([instalar aqui](https://golang.org/doc/install))
- SQLite instalado (opcional, para visualizar o banco diretamente)

### Passos

```bash
# Instalar dependências
go mod tidy

# Rodar a aplicação
go run main.go 

 

📦 Tecnologias
Go 


SQLite

Driver: modernc.org/sqlite (100% Go, sem CGO)

Echo framework 

📄 Autor
William Viegas 

