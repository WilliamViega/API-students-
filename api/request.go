package api

type CreateStudentRequest struct {
    Nome  string `json:"nome"`
    CPF   string `json:"cpf"`
    Email string `json:"email"`
    Idade int    `json:"idade"`
}

func (r CreateStudentRequest) IsValid() (bool, string) {
    if r.Nome == "" {
        return false, "Nome é obrigatório"
    }
    if len(r.CPF) < 11 {
        return false, "CPF inválido"
    }
    if r.Email == "" {
        return false, "Email é obrigatório"
    }
    if r.Idade <= 0 {
        return false, "Idade deve ser maior que zero"
    }
    return true, ""
}
 