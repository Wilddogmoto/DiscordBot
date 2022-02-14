package responseBot

import (
	"bufio"
	"github.com/Wilddogmoto/DiscordBot/data"
	"github.com/Wilddogmoto/DiscordBot/loggds"
	"os"
)

func AddBlackList(name string) {

	var (
		log   = loggds.Logg
		err   error
		lines []string
	)

	file, err := os.Open(name)
	if err != nil {
		log.Errorf("error open file: %s", err)
		return
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Errorf("error close file: %s", err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		lines = append(lines, scanner.Text())
	}

	for _, r := range lines {

		word := data.BlackList{
			Word: r,
		}
		if err = data.DataBase.Create(&word).Error; err != nil {
			log.Errorf("error on creating black list: %s", err)
			return
		}
	}
}
