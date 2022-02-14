package connect

import (
	"github.com/Wilddogmoto/DiscordBot/loggds"
	"github.com/Wilddogmoto/DiscordBot/responseBot"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func ConnDiscordBot(token string) {

	var (
		log = loggds.Logg
		dg  *discordgo.Session
		err error
	)

	// Create a new Discord session
	dg, err = discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("error on create bot connection: %s", err)
	}

	dg.StateEnabled = true

	//Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(responseBot.MessageBotResponse) //public command for bot
	dg.AddHandler(responseBot.AdminCommand)       //private command for bot

	// receiving message events
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord
	if err = dg.Open(); err != nil {
		log.Errorf("error open connection: %s", err)
		return

	}

	log.Info("Bot is running!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// after close  Discord session
	if err = dg.Close(); err != nil {
		log.Errorf("error on close connection: %s", err)
	}

}
