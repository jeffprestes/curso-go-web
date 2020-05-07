# Curso de Go Web - MVC

## Objetivo

Este aplicação serve de material para o Curso de Go Web - MVC. Este curso ensina como criar uma aplicaçao Web em Go usando o padrão MVC sobre o framework [Go-Macaron](https://go-macaron.com) e seguindo o modelo (template) do [Mercurius](https://github.com/novatrixtech/mercurius) que já traz muitos artefatos prontos e economiza muito tempo.

## Preparação antes do curso

Usamos um banco de dados MySQL ou MariaDB nesse exemplo. Por favor crie um banco chamado **curso_go_web** e um usuário chamado sysapp com a senha 'senhaSimples1234' . Depois execute os scripts SQL que estão nos diretórios _/lib/auth_ e depois em \*/dados-tabela\_

## Para executar a aplicação

Basta executar `go run main.go` do diretório raiz da aplicação. Caso Go reclame da ausencia de alguma dependencia execute `go install` para que _go modules_ baixem as dependencias ou execute `go get ./...` para que Go baixe todas as dependencias em seu GOPATH.

Se estiver num servidor Linux e quiser iniciar a aplicação como um serviço, clone este projeto dentro do diretório `/srv`, depois copie o arquivo curso-go-web.service para `/etc/systemd/system/`, daí execute os seguintes comandos: `sudo systemctl daemon-reload && sudo systemctl start curso-go-web && sudo systemctl status curso-go-web && sudo systemctl enable curso-go-web` . Se tudo estiver OK com as permissões a aplicação deverá subir na porta 8080 e ser recarregada quando o servidor for reiniciado. Se quiser rodar na porta 80 edite o arquivo `conf/app.ini`

## Para compilar a aplicação

`go build` para ele compilar no diretório local e para sua arquitetura de desenvolvimento. Para um server Linux, por exemplo, você deve executar `GOARCH=amd64 GOOS=linux go build -o curso-go-web-linux-binary`

## Para acessar a aplicação

Acesse em seu navegador http://localhost:8080

Caso queira mudar a porta acesse o arquivo _conf/app.ini_ e mude a variável **http_port** ou ainda defina através da criação da variável de ambiente **PORT** com o valor desejado para a porta.
