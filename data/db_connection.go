package data

import (
	"fmt"
	"github.com/Wilddogmoto/DiscordBot/cfg"
	"github.com/Wilddogmoto/DiscordBot/loggds"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
)

type (
	Account struct {
		Server   string
		User     string
		Password string
		DB       string
	}
)

var (
	DataBase *gorm.DB
)

func DBConnect() {

	var (
		migrations []interface{}
		log        = loggds.Logg
		err        error
		connection = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&amp,utf8&parseTime=True",
			cfg.Parameter.DataBase.User,
			cfg.Parameter.DataBase.Password,
			cfg.Parameter.DataBase.Server,
			cfg.Parameter.DataBase.DB,
		)
	)

	log.Info("starting database connection...")

	DataBase, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatalf("bad connection to database: %s", err)
	}

	log.Info("start migration")

	//add models
	migrations = append(migrations, &BlackList{})

	for _, model := range migrations {
		name := reflect.TypeOf(model).Elem().Name()

		if err = DataBase.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(model); err != nil {
			log.Fatalf("error on creating table %s: %v", name, err)
		}
	}

	log.Info("database migration success")
}
