CREATE TABLE clientes (
  cli_id INT NOT NULL AUTO_INCREMENT,
  cli_nome VARCHAR(255) NULL,
  cli_cidade VARCHAR(255) NULL,
  cli_uf VARCHAR(2) NULL,
  cli_cep VARCHAR(8) NULL,
  PRIMARY KEY (cli_id));

insert into clientes (cli_nome, cli_cidade, cli_uf, cli_cep) values ('José Bonifácio', 'Rio de Janeiro', 'RJ', '21560030');
insert into clientes (cli_nome, cli_cidade, cli_uf, cli_cep) values ('José de Alencar', 'Fortaleza', 'CE', '85741010');