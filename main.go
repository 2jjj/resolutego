package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"resolutebot/handler"

	"github.com/bwmarrin/discordgo"
)

type Config struct {
	Token string `json:"token"`
}

func main() {
	// Cria uma nova sessão do Discord
	config, err := loadConfig("config.json")
	if err != nil {
		fmt.Println("Erro ao carregar configurações:", err)
		return
	}

	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("Erro ao criar sessão do Discord:", err)
		return
	}
	// Adiciona um handler para responder a mensagens
	dg.AddHandler(handler.PingHandler)

	// Abre uma conexão com o Discord
	err = dg.Open()
	if err != nil {
		fmt.Println("Erro ao abrir conexão com o Discord:", err)
		return
	}

	fmt.Println("Bot está rodando. Pressione CTRL+C para encerrar.")
	// Espera pelo encerramento do programa
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	// Fecha a sessão
	dg.Close()
}

func loadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}