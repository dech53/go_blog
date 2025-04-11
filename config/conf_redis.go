package config

type Redis struct {
	Address  string `json:"address" yaml:"address"`
	Password string `json:"password" yaml:"password"`
	DB       int    `json:"db" yaml:"db"`
}
