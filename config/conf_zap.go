package config

type Zap struct {
	Level          string `json:"level" yaml:"level"`
	Filename       string `json:"filename" yaml:"filename"`
	MaxSize        int    `json:"max_size" yaml:"max_size"`
	MaxBackups     int    `json:"max_backups" yaml:"max_backups"`
	MaxAge         int    `json:"max_age" yaml:"max_age"`
	IsConsolePrint bool   `json:"is_console_print" yaml:"is_console_print"`
}
