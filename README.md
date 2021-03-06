## API Goponto

API de gerenciamento, inclusão e consulta das marcações de ponto, da aplicação Goponto

[![Hex.pm](https://img.shields.io/hexpm/l/plug.svg)](https://github.com/goponto/goponto-api/blob/master/LICENSE)

Goponto é um aplicativo online para controle `pessoal` de registro de ponto e totalização de horas.

[English version here](https://github.com/goponto/goponto-api/master/README-en.md)

## Documentação

- [Introdução](#introdução)
- [Frameworks utilizados](#frameworks)
- [Funcionalidades](#funcionalidades)
- [Metodos](#métodos)
- [Banco de dados](#banco-de-dados)
- [Modo de uso](#modo-de-uso)
- [Filtros de pesquisa](#filtros-de-busca)
- [licença](#licença)

### Introdução

Este projeto via realizar o controle `pessoal` dos registro de ponto, ou seja, substituir a velha planilha e fornecer um meio mais prático e organizado para conferência dos registro de ponto. 

Apesar de não fazer quaisquer integração com sistemas de ponto eletrônico, vale citar que as informaçõe contidas no recido de ponto, são normatizada e obrigatórias.

Na [PORTARIA Nº 1.510, DE 21 DE AGOSTO DE 2009](http://www.trtsp.jus.br/geral/tribunal2/ORGAOS/MTE/Portaria/P1510_09.html) no Art. 11 – § 2º, informa o campos mínimos que deve ser considerados no recibo impresso. Desta forma estes mesmos campos vão ser tratados nesta API. 

- Cabeçalho contendo o título “Comprovante de Registro de Ponto do Trabalhador”;
- Identificação do empregador com nome, CNPJ/CPF e CEI, caso exista;
- O local da prestação de serviço;
- O número de fabricação do REP;
- A identificação do trabalhador contendo nome e número do PIS;
- A data e horário do registro; e
- O NSR (Número Sequencial de Registro).

### Frameworks

- [framework Echo](https://echo.labstack.com/)

### Funcionalidades

- [ ] Inclusão, alteração e exclusão de registros de ponto;
- [ ] Inclusão, alteração de empresa;
- [ ] Autenticação *
- [ ] Upload de imagens *

### Fazer

- [ ] Tudo!

### Métodos

Simples API REST:


| URI      | HTTP Metodo | Contexto   | Descrição
| --------- | ----------- | ---------- | -------------
| `/`       | GET         | Register   | Retorna os dados de um registro de ponto pelo ID.
| `/`       | GET         | Register   | Lista/pesquisa um resgistro usando filtros e ordenação.
| `/register`  | POST        | Register   | Cria um novo registro, gerando um novo ID.
| `/register/:id` | PUT         | Register   | Atualiza prarcialmente um registro pelo ID.
| `/register/:id`  | DELETE      | Register   | Apara um registro pelo ID
| `/company`  | POST        | Company   | Cria uma nova empresa.
| `/company/:id` | PUT         | Company   | Atualiza prarcialmente uma empresa pelo ID.
| `/company/:id`  | DELETE      | Company   | Apara uma empresa pelo ID, se não possuir marcações de ponto.

### Banco de dados

- [ ] Postgres e/ou Standalone;

### Modo de uso

### Filtros de busca

## Licença

Todo código está sobre a licença [Apache License 2.0](https://github.com/goponto/goponto-api/blob/master/LICENSE).
