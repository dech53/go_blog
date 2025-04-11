package config

type ES struct {
	URL            string `json:"url" yaml:"url"`
	Username       string `json:"username" yaml:"username"`
	Password       string `json:"password" yaml:"password"`
	IsConsolePrint bool   `json:"is_console_print" yaml:"is_console_print"`
}
