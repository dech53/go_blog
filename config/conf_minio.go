package config

type Minio struct {
	Bucket    string `json:"bucket" yaml:"bucket"`
	ImgPath   string `json:"img_path" yaml:"img_path"`
	EndPoint  string `json:"end_point" yaml:"end_point"`
	AccessKey string `json:"access_key" yaml:"access_key"`
	SecretKey string `json:"secret_key" yaml:"secret_key"`
	UseSSL    bool   `json:"use_ssl" yaml:"use_ssl"`
}
