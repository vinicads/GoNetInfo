# Aplicacao de Linha de Comando em Go

Esta é uma aplicação de linha de comando desenvolvida em Go para buscar IPs e nomes de servidores na Internet. O projeto foi criado seguindo um curso de Go.

## Tecnologias Utilizadas

- Go
- Biblioteca `urfave/cli` para criação da CLI

## Instalação

Clone este repositório e troque a URL conforme necessário:

```sh
 git clone https://github.com/vinicads/GoNetInfo.git
 cd GoNetInfo
```

## Como Usar

Para rodar a aplicação, utilize um dos comandos abaixo:

### Buscar Nome do Servidor
```sh
go run main.go servidor --host exemplo.com
```

### Buscar IPs
```sh
go run main.go ip --host exemplo.com
```

Caso não seja informado um `--host`, o padrão será `google.com.br`.

## Estrutura do Projeto

```
/<NOME_DO_PROJETO>
│── main.go
│── app/
│   ├── app.go
│── go.mod
│── go.sum
│── README.md
```

## Autor

Projeto desenvolvido seguindo um curso de Go pelo Otávio Augusto Gallego.
