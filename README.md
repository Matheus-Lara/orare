<h1 align="center">Orare.</h1>

<p align="center"><img src="https://github.com/user-attachments/assets/75aef21c-0e96-420f-8f87-ab2f52d07cff"></img></p>

---

De horários de missas a outros eventos, conectamos você com sua espiritualidade.

Este repositório contém a API backend do Orare.

## Tecnologias Utilizadas

- Back-end:
  - Go lang
    - GIN Framework
    - AWS lambda (infraestrutura)
  - CockroachDB (postgres)

## Setup

### Dependências
- [Go](https://go.dev/) (v1.23.0 ou superior) - Linguagem de programação de código aberto que facilita a construção de software simples, confiável e eficiente.
- [Docker](https://www.docker.com/) - Plataforma para desenvolvimento, envio e execução de aplicativos em contêineres.
- [Docker Compose](https://docs.docker.com/compose/) - Ferramenta para definir e executar aplicativos Docker com múltiplos contêineres usando um arquivo YAML.


### Executando o projeto

1. Para criar o .env e arquivos usados nos volumes dos containers docker:
```
make setup
```

> Este comando irá configurar a variável GOPATH, criar o diretório de volumes do docker, copiar o arquivo .env.example para .env (se não existir) e instalar o Wire.

> Realize eventuais ajustes nas portas do `.env` caso alguma porta já esteja ocupada em seu ambiente local.

2. Para subir o container docker:
```
docker compose up
```

3. Para resolver as dependências, na raiz do projeto, execute:

```
go mod tidy
```

4. Para iniciar o projeto, execute:

```
go run cmd/orare/main.go
```

A mensagem de log `[GIN-debug] Listening and serving HTTP on :<porta>` deve aparecer e será possível acessar a rota `localhost:<porta>/api/health` para o health check.

