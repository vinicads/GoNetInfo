package app

import (
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar retorna uma nova aplicação pronta para ser executada
func Gerar() *cli.App{
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"
	
	flags := []cli.Command{
		{
			Name: "ip",
			Usage: "Busca IPs de endereços na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "host",
					Value: "google.com.br",
				},
			},
			Action: buscarIps,
		},	
		{
			Name: "servidor",
			Usage: "Busca o nome do servidor na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "host",
					Value: "google.com.br",
				},
			},
			Action: buscarServidores,
		},
	}
	
	app.Commands = flags;

	return app
}

func buscarIps(c *cli.Context)  {
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		log.Println(ip)
	}
}

func buscarServidores(c *cli.Context)  {
	host := c.String("host")
	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		log.Println(servidor.Host)
	}
}