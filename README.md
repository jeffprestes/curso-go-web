# Curso de Go Web - MVC

## Objetivo

Este aplicação serve de material para o Curso de Go Web - MVC. Este curso ensina como criar uma aplicaçao Web em Go usando o padrão MVC sobre o framework [Go-Macaron](https://go-macaron.com) e seguindo o modelo (template) do [Mercurius](https://github.com/novatrixtech/mercurius) que já traz muitos artefatos prontos e economiza muito tempo.

## Preparação antes do curso

Usamos um banco de dados MySQL ou MariaDB nesse exemplo. Por favor crie um banco chamado curso*go_web e um usuário chamado sysapp com a senha 'senhaSimples1234' . Depois execute os scripts SQL que estão nos diretórios */lib/auth* e depois em */dados-tabela\_

## Para executar a aplicação

Basta executar `go run main.go` do diretório raiz da aplicação

## Para compilar a aplicação

`go build` para ele compilar no diretório local e para sua arquitetura de desenvolvimento. Para um server Linux, por exemplo, você deve executar `GOARCH=amd64 GOOS=linux go build -o curso-go-web-linux-binary`

## Para acessar a aplicação

Acesse em seu navegador http://localhost:8080

Caso queira mudar a porta acesse o arquivo _conf/app.ini_ e mude a variável **http_port** ou ainda defina através da criação da variável de ambiente **PORT** com o valor desejado para a porta.
