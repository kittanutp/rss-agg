package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", errors.New("unauthorized")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 || vals[0] != "Bearer" {
		return "", errors.New("invalid authentication")
	}
	return vals[1], nil
}
