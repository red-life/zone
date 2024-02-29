package resolver

import "encoding/json"

func FromString[T any](s string) (T, error) {
	var t T
	err := json.Unmarshal([]byte(s), &t)
	return t, err

}

func MustFromString[T any](s string) T {
	t, err := FromString[T](s)
	if err != nil {
		panic(err)
	}
	return t
}
