package appTypes

import "encoding/json"

type Storage int

const (
	Local Storage = iota
	Minio
)

func (s Storage) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *Storage) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*s = ToStorage(str)
	return nil
}

func (s Storage) String() string {
	var str string
	switch s {
	case Local:
		str = "本地"
	case Minio:
		str = "minio"
	default:
		str = "未知存储"
	}
	return str
}

func ToStorage(str string) Storage {
	switch str {
	case "本地":
		return Local
	case "minio":
		return Minio
	default:
		return -1
	}
}
