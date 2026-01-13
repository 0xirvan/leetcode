package main

import (
	"fmt"
	"strings"
)

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var b strings.Builder
	for _, str := range strs {
		b.WriteString(fmt.Sprintf("%d%s", len(str), "#"))
		b.WriteString(str)
	}
	return b.String()
}

func (s *Solution) Decode(encoded string) []string {
	var res []string
	i := 0

	for i < len(encoded) {
		l := 0
		for encoded[i] >= '0' && encoded[i] <= '9' {
			l = l*10 + int(encoded[i]-'0')
			i++
		}
		i++
		res = append(res, encoded[i:i+l])
		i += l
	}
	return res
}
