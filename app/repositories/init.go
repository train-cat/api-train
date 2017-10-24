package repositories

import (
	"fmt"
	"os"

	"aahframework.org/aah.v0"
	"aahframework.org/log.v0"
	"github.com/jinzhu/gorm"
	// Load MySQL driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	code "github.com/train-sh/api-train/app/errors"
)

var db *gorm.DB

func init() {
	aah.OnStart(InitDatabase)
}

// InitDatabase connection
func InitDatabase(_ *aah.Event) {
	config, ok := aah.AppConfig().GetSubConfig("database")

	if !ok {
		log.Errorf("Missing database configuration")
		os.Exit(code.MissingConfigDatabase)
	}

	username, ok := config.String("username")
	exitOnError("username", ok)
	password, ok := config.String("password")
	exitOnError("password", ok)
	hostname, ok := config.String("hostname")
	exitOnError("hostname", ok)
	port, ok := config.Int("port")
	exitOnError("port", ok)
	name, ok := config.String("name")
	exitOnError("name", ok)
	debug, ok := config.Bool("debug")
	exitOnError("debug", ok)

	d, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, hostname, port, name))

	if err != nil {
		log.Errorf("Can't open database connection: %s", err.Error())
		os.Exit(code.InitDatabase)
	}

	if debug {
		db = d.Debug()
	} else {
		db = d
	}

	db.SingularTable(true)
}

func exitOnError(key string, ok bool) {
	if !ok {
		log.Errorf("Key %s missing from database configuration", key)
		os.Exit(code.MissingConfigDatabase)
	}
}
