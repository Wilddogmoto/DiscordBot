package mid

import (
	"github.com/Wilddogmoto/DiscordBot/data"
	"github.com/Wilddogmoto/DiscordBot/loggds"
	"github.com/bwmarrin/discordgo"
)

// CheckAdmin Checks message was sent by admin or owner
func CheckAdmin(s *discordgo.Session, m *discordgo.MessageCreate) bool {

	var (
		log   = loggds.Logg
		guild *discordgo.Guild
		err   error
	)

	if guild, err = s.Guild(m.GuildID); err != nil {
		log.Errorf("error on getting guild: %s", err)

		return false
	}

	for _, role := range m.Member.Roles { // check member role is admin or owner
		if role == data.Admin || guild.OwnerID == m.Author.ID {

			return true
		}
	}

	log.Errorf("error on getting roles member: %s", m.Author.Username)
	return false
}
