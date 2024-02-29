package management

import "encoding/json"

func String(data any) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func MustString(data any) string {
	s, err := String(data)
	if err != nil {
		panic(err)
	}
	return s
}
