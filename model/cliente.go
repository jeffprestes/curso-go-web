package model

//Cliente representa o modelo da entidade Cliente
type Cliente struct {
	ID     int    `db:"cli_id" form:"formID"`
	Nome   string `db:"cli_nome" form:"formNome" binding:"Required"`
	Cidade string `db:"cli_cidade" form:"formCidade" binding:"Required"`
	UF     string `db:"cli_uf" form:"formUF" binding:"Required"`
	CEP    string `db:"cli_cep" form:"formCEP" binding:"Required"`
}
