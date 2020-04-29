package repo

import (
	"github.com/jeffprestes/curso-go-web/conf"
	"github.com/jeffprestes/curso-go-web/model"
)

//GetClientes busca todos os clientes no banco de dados
func GetClientes() (clientes []model.Cliente, err error) {
	db, err := conf.GetDB()
	if err != nil {
		return
	}
	err = db.Select(&clientes, "SELECT * FROM clientes")
	if err != nil {
		return
	}
	return
}
