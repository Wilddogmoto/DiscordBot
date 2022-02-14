package main

import (
	"flag"
	"github.com/Wilddogmoto/DiscordBot/cfg"
	"github.com/Wilddogmoto/DiscordBot/connect"
	"github.com/Wilddogmoto/DiscordBot/data"
	"github.com/Wilddogmoto/DiscordBot/loggds"
)

func main() {

	filename := flag.String("config", "", "path to config file")
	flag.Parse()

	//init logger
	loggds.InitLogger()

	//init bot configuration
	cfg.Configuration(*filename)

	//init database connect
	data.DBConnect()

	//init connect bot
	connect.ConnDiscordBot(cfg.Parameter.Discord.Token)

	//responseBot.AddBlackList("./blacklist.txt")

}
