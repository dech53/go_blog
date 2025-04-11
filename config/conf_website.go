package config

// Website 网站信息
type Website struct {
	Logo                 string `json:"logo" yaml:"logo"`
	FullLogo             string `json:"full_logo" yaml:"full_logo"`
	Title                string `json:"title" yaml:"title"`
	Slogan               string `json:"slogan" yaml:"slogan"`
	SloganEn             string `json:"slogan_en" yaml:"slogan_en"`
	Description          string `json:"description" yaml:"description"`
	Version              string `json:"version" yaml:"version"`
	CreatedAt            string `json:"created_at" yaml:"created_at"`
	IcpFiling            string `json:"icp_filing" yaml:"icp_filing"`
	PublicSecurityFiling string `json:"public_security_filing" yaml:"public_security_filing"`
	BilibiliURL          string `json:"bilibili_url" yaml:"bilibili_url"`
	GiteeURL             string `json:"gitee_url" yaml:"gitee_url"`
	GithubURL            string `json:"github_url" yaml:"github_url"`
	Name                 string `json:"name" yaml:"name"`
	Job                  string `json:"job" yaml:"job"`
	Address              string `json:"address" yaml:"address"`
	Email                string `json:"email" yaml:"email"`
	QQImage              string `json:"qq_image" yaml:"qq_image"`
	WechatImage          string `json:"wechat_image" yaml:"wechat_image"`
}
