package responseBot

import (
	"github.com/Wilddogmoto/DiscordBot/loggds"
	"github.com/Wilddogmoto/DiscordBot/mid"
	"github.com/bwmarrin/discordgo"
)

// AdminCommand private command for bot
func AdminCommand(s *discordgo.Session, m *discordgo.MessageCreate) {

	var (
		log = loggds.Logg
		err error
	)

	//ignore bot messages
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	//check admin rules
	if !mid.CheckAdmin(s, m) {
		return
	}

	switch {
	case m.Content == "!help-adm":
		if _, err = s.ChannelMessageSend("", "true"); err != nil {
			// internal error
			log.Errorf("error send message for channel (%v); %s", m.ChannelID, err)
			return
		}

	case m.Content == "!img":

		if _, err = s.ChannelMessageSend("562362153026453507", "up"); err != nil {
			// internal error
			log.Errorf("error send message for channel (%v); %s", m.ChannelID, err)
			return
		}
		return
	}

}
