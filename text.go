package main

import "strings"

type YesNoIdk int

const (
	Idk YesNoIdk = iota
	Yes
	No
)

func IsYesNoIdk(s string) YesNoIdk {
	s = strings.ToLower(s)
	strs := strings.Fields(s)
	for _, f := range strs {
		if f == "y" || f == "yes" || f == "yup" || f == "yay" {
			return Yes
		}
		if f == "n" || f == "no" || f == "nope" || f == "nay" {
			return No
		}
	}
	return Idk
}
