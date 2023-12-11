package utils

import (
	"encoding/json"
	"regexp"

	"github.com/spf13/cast"
)

func Marshal(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func UnMarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func VerifyPassword(password, hashedPassword string) bool {
	return password == hashedPassword
}

func VerifyEmail(email string) bool {
	reg := regexp.MustCompile(`^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+`)
	return reg.MatchString(email)
}

func TransformStringToInt64(list []string) []int64 {
	var result []int64
	for _, v := range list {
		result = append(result, cast.ToInt64(v))
	}
	return result
}
