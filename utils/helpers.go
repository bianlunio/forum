package utils

import (
	"bytes"
	"strconv"
)

func String2Int(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

func JoinStrings(args ...string) string {
	var buffer bytes.Buffer
	for _, s := range args {
		buffer.WriteString(s)
	}
	return buffer.String()
}

func JoinUrl(args ...string) string {
	var s []string
	for i, arg := range args {
		if i != 0 {
			s = append(s, "/")
		}
		s = append(s, arg)
	}
	return JoinStrings(s...)
}

type Query map[string]string

func ParseQueryObject(query Query) string {
	s := []string{"?"}
	for k, v := range query {
		s = append(s, k, "=", v, "&")
	}
	s = s[:len(s)-1]
	return JoinStrings(s...)
}

func ParseQueryUrl(baseUrl string, query Query) string {
	qs := ParseQueryObject(query)
	return JoinStrings(baseUrl, qs)
}
