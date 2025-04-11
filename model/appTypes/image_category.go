package appTypes

import (
	"encoding/json"
)

type Category int

const (
	Null Category = iota
	System
	Carousel
	Cover
	Illustration
	AdImage
	Logo
)

func (c Category) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *Category) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*c = ToCategory(str)
	return nil
}

func (c Category) String() string {
	switch c {
	case Null:
		return "未使用"
	case System:
		return "系统"
	case Carousel:
		return "背景"
	case Cover:
		return "封面"
	case Illustration:
		return "插图"
	case AdImage:
		return "广告"
	case Logo:
		return "友链"
	default:
		return "未知类别"
	}
}

func ToCategory(str string) Category {
	switch str {
	case "未使用":
		return Null
	case "系统":
		return System
	case "背景":
		return Carousel
	case "封面":
		return Cover
	case "插图":
		return Illustration
	case "广告":
		return AdImage
	case "友链":
		return Logo
	default:
		return -1
	}
}
