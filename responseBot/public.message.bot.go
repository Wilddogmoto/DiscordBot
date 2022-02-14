package responseBot

import (
	"bufio"
	"github.com/Wilddogmoto/DiscordBot/loggds"
	"github.com/Wilddogmoto/DiscordBot/mid"
	"github.com/bwmarrin/discordgo"
	"os"
	"regexp"
	"strings"
)

// MessageBotResponse public command for bot
func MessageBotResponse(s *discordgo.Session, m *discordgo.MessageCreate) {

	var (
		log = loggds.Logg
		err error
	)

	//ignore bot messages
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	//commands for bot
	switch {

	case strings.Contains(m.Content, s.State.User.ID):

		if _, err = s.ChannelMessageSend(m.ChannelID, "Бот на связи"); err != nil {
			// internal error
			log.Errorf("error send message for channel (%v); %s", m.ChannelID, err)
			return
		}
		return

	case m.Content == "!up":
		if _, err = s.ChannelMessageSend("", "up"); err != nil {
			// internal error
			log.Errorf("error send message for channel (%v); %s", m.ChannelID, err)
			return
		}
		return

	case m.Content == "!test":

		if _, err = s.ChannelMessageSend(m.ChannelID, "test"); err != nil {
			// internal error
			log.Errorf("error send message for channel (%v); %s", m.ChannelID, err)
			return
		}
		return

	default: //check message on black list

		//ignored blacklist for admin or owner
		if mid.CheckAdmin(s, m) {
			return
		}

		if !checkBlackList(m.Content) {
			//if false, word not found in blacklist
			return
		}

		//delete message if found in black list
		if err = s.ChannelMessageDelete(m.ChannelID, m.Message.ID); err != nil {
			// internal error
			log.Errorf("error delete message for channel (%v); %s", m.ChannelID, err)
			return
		}
	}
}

//checkBlackList Checks the message for forbidden words in the blacklist
func checkBlackList(message string) bool {

	var (
		log     = loggds.Logg
		err     error
		matched bool
	)

	file, err := os.Open("./blacklist.txt")
	if err != nil {
		// internal error
		log.Errorf("error open file: %s", err)
		return false
	}

	defer func() {
		if err = file.Close(); err != nil {
			// internal error
			log.Errorf("error close file: %s", err)
		}
	}()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		matched, err = regexp.MatchString(scanner.Text(), message)
		if err != nil {
			// internal error
			log.Errorf("word search error: %s", err)
			return false
		}
		if matched {
			return true
		}
	}

	return false
}
