package config

type Email struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	From     string `json:"from" yaml:"from"`
	Nickname string `json:"nickname" yaml:"nickname"`
	Secret   string `json:"secret" yaml:"secret"`
	IsSSL    bool   `json:"is_ssl" yaml:"is_ssl"`
}
