package conf

import "github.com/jasontconnell/conf"

type Config struct {
	ConnectionString string `json:"connectionString"`
	SrcFieldId       string `json:"srcFieldId"`
	DestFieldId      string `json:"destFieldId"`
	DestTable        string `json:"destTable"`
}

func LoadConfig(filename string) Config {
	cfg := Config{}
	conf.LoadConfig(filename, &cfg)
	return cfg
}
