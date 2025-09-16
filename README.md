# 🧠 API Students

API desenvolvida em Go para o curso **"Golang do Zero"** ministrado por **Stephanie Cardoso** (Globo).  
Essa aplicação gerencia estudantes usando o framework Echo e banco de dados SQLite.

---

## 📚 Rotas disponíveis

- `GET /students` — Lista todos os estudantes
- `POST /students` — Cria um novo estudante
- `GET /students/:id` — Busca um estudante específico por ID
- `PUT /students/:id` — Atualiza o nome de um estudante
- `DELETE /students/:id` — Remove um estudante
- `GET /students?active=<true/false>` — Filtra estudantes ativos/inativos

---

## 🧬 Estrutura do estudante

```go
type Student struct {
    Nome   string
    CPF    int
    Email  string
    Idade  int
    Active bool
}

=======
# API-students-
API- to ' Golang do zero ' course students
>>>>>>> d77e8bdc1bc3478afb1720a71013f5fb7520f580
# API Students


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

Echo Framework

SQLite com driver modernc.org/sqlite (100% Go, sem CGO)

Swagger para documentação interativa dos endpoints 

## 📸 Exemplo de uso


Abaixo, uma prévia da API funcionando via Swagger:

![Swagger interface](assets/print.jpeg)

📄 Autor
Desenvolvido por William Viegas 
Baseado no curso de Stephanie Cardoso, desenvolvedora na Globo 🚀

