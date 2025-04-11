package config

import (
	"fmt"
	"server/model/appTypes"
	"strings"
)

// System 系统配置
type System struct {
	Host           string `json:"-" yaml:"host"`
	Port           int    `json:"-" yaml:"port"`
	Env            string `json:"-" yaml:"env"`
	RouterPrefix   string `json:"-" yaml:"router_prefix"`
	UseMultipoint  bool   `json:"use_multipoint" yaml:"use_multipoint"`
	SessionsSecret string `json:"sessions_secret" yaml:"sessions_secret"`
	OssType        string `json:"oss_type" yaml:"oss_type"`
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

func (s System) Storage() appTypes.Storage {
	switch strings.ToLower(s.OssType) {
	case "local", "Local":
		return appTypes.Local
	case "minio", "Minio":
		return appTypes.Minio
	default:
		return appTypes.Local
	}
}
