//Author Mohammad Naser Abbasanadi
//Creating Date 2018-10-20
// config.go is to gathering all configuration at one point

package config

import (
	"log"
	"os"
	"sync"

	"github.com/jinzhu/configor"
)

//Config structure
type Config struct {
	APPName string `default:"GolangOrdering"`

	SERVER struct {
		Address string `default:":8080"`
	}

	DB struct {
		Address  string `default:"185.159.153.22"`
		User     string `default:"karnik"`
		Dbname   string `default:"orders"`
		Password string `required:"true" default:"123456"`
		Port     string `default:"3306"`
	}

	LOGPATH string `default:"/tmp/log"`

	APIKEY string `default:"AIzaSyBzuM7atg360ClN4hmao7J3Y0UbvxSrkx8"`
}

var instance Config

var once sync.Once

//GetConfigInstance singletone
func GetConfigInstance() Config {
	once.Do(func() {
		pwd, _ := os.Getwd()

		err := configor.Load(&instance, pwd+"/config/config.yaml")
		log.Println(instance)
		if err != nil {
			log.Fatal("Error reading config structure: " + err.Error())
		}
	})
	return instance
}
