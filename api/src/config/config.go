package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//STRING QUE CONECTA COM BANCO MYSQL
	StringConexaoBanco = ""

	//PORTA ONDE A API VAI ESTA RODANDO
	Porta = 0
)

//VAI CARREGAR E INICALIZAR AS VARIAVEIS DE AMBIENTE
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 8000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8", 
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

}	