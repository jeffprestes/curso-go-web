package repo

import (
	"log"

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

//GetClientePeloID busca o cliente no banco de dados pelo seu ID
func GetClientePeloID(ID int) (cliente model.Cliente, err error) {
	db, err := conf.GetDB()
	if err != nil {
		return
	}
	err = db.Get(&cliente, "SELECT * FROM clientes WHERE cli_id=?", ID)
	if err != nil {
		return
	}
	return
}

//AtualizaCliente atualiza dados do cliente no banco de dados
func AtualizaCliente(form model.Cliente) (err error) {
	db, err := conf.GetDB()
	if err != nil {
		return
	}
	_, err = db.Exec("UPDATE clientes SET cli_nome=?, cli_cidade=?, cli_uf=?, cli_cep=? WHERE cli_id=?", form.Nome, form.Cidade, form.UF, form.CEP, form.ID)
	if err != nil {
		log.Println("AtualizaCliente - db.Exec - Erro: ", err.Error())
		return
	}
	log.Printf("form %+v\n", form)
	return
}

//InsereCliente insere novo cliente no banco de dados
func InsereCliente(form model.Cliente) (novoCliente model.Cliente, err error) {
	db, err := conf.GetDB()
	if err != nil {
		return
	}
	res, err := db.Exec("INSERT INTO clientes (cli_nome, cli_cidade, cli_uf, cli_cep) VALUES (?, ?, ?, ?)", form.Nome, form.Cidade, form.UF, form.CEP)
	if err != nil {
		return
	}
	novoID, err := res.LastInsertId()
	if err != nil {
		return
	}
	novoCliente.ID = int(novoID)
	novoCliente.CEP = form.CEP
	novoCliente.Cidade = form.Cidade
	novoCliente.Nome = form.Nome
	novoCliente.UF = form.UF
	return
}
