package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string `yaml:"address" env-required:"true"`
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var ConfigPath string

	ConfigPath = os.Getenv("")
	if ConfigPath == "" {
		flags := flag.String("config", "", "path related to config path")
		flag.Parse()

		ConfigPath = *flags

		if ConfigPath == "" {
			log.Fatal("config path is not set")
		}
	}

	// check on provided path we have given the config file or not

	if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
		log.Fatalf("config path file does not exist %s", ConfigPath)
	}

	// if file set, we have to read the file and serialze the file as per the enviroment (eg, local,prod)

	var cfg Config

	//cleanenv.Readconfig  - return the error

	err := cleanenv.ReadConfig(ConfigPath, &cfg)
	if err != nil {
		log.Fatalf("error while reading the configpath file %s", err.Error())

	}

	return &cfg

}
