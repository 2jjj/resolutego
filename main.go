package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// Substitua pelo seu token do bot
	Token := "OTUyNzA5NDk5MzE4MzE3MTI2.G7eI5w.GMXSQuFILgFlodLqdc9TN2a6-k4fIUyXGjkEO0"

	// Cria uma nova sessão do Discord
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Erro ao criar sessão do Discord:", err)
		return
	}

	// Adiciona um handler para responder a mensagens
	dg.AddHandler(messageCreate)

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

// Handler para mensagens recebidas
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignora mensagens enviadas pelo próprio bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Responde a mensagens específicas
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong!")
	} else if m.Content == "checksaldos" {
		s.ChannelMessageSend(m.ChannelID, "verificando saldos")
        
	}
}
