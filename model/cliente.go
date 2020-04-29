package model

//Cliente representa o modelo da entidade Cliente
type Cliente struct {
	ID     int    `db:"cli_id"`
	Nome   string `db:"cli_nome"`
	Cidade string `db:"cli_cidade"`
	UF     string `db:"cli_uf"`
	CEP    string `db:"cli_cep"`
}
