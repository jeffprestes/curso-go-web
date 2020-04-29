package handler

import (
	"net/http"

	"github.com/jeffprestes/curso-go-web/lib/contx"
)

//IndexCliente abre a pagina de gerenciamento de clientes
func IndexCliente(ctx *contx.Context) {
	ctx.NativeHTML(http.StatusOK, "index")
}
