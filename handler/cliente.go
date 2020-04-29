package handler

import (
	"log"
	"net/http"

	"github.com/jeffprestes/curso-go-web/lib/contx"
	"github.com/jeffprestes/curso-go-web/repo"
)

//IndexCliente abre a pagina de gerenciamento de clientes
func IndexCliente(ctx *contx.Context) {
	clientes, err := repo.GetClientes()
	if err != nil {
		log.Println("IndexCliente - repo.GetClientes - Error: ", err.Error())
		ctx.NativeHTML(http.StatusInternalServerError, "erro")
		return
	}
	ctx.Data["clientes"] = clientes
	ctx.NativeHTML(http.StatusOK, "index")
}
