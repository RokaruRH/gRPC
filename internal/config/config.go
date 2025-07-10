package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env          string        `yaml:"env" env-default:"local"`
	Storage_path string        `yaml:"storage_path" env-required:"true"`
	Token_ttl    time.Duration `yaml: "token_ttl" env-required:"true"`
	GRPC         GRPCConfig    `yaml: "grpc"`
}

type GRPCConfig struct {
	Port   int           `yaml:"port"`
	Timeot time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	path := fetchConfigPath()

	if path == "" {
		panic("SHIT DEAL GUY. THAT METH IS BAD FOR US")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("OH SHIT THIS CONFIG ISN'T EXIST RETREAT BEACHES")
	}

	var config Config
	cleanenv.ReadConfig(path, &config)

	if err := cleanenv.ReadConfig(path, &config); err != nil {
		panic("WHY YOU CHEATING ME??? WHY AT MY BANK CASH IS 0???")
	}
	return &config
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG")
	}
	return res
}
