package main

import "strings"

func concat(xs []string) string {
	s := xs[0]
	for _, v := range xs[1:] {
		s += " "
		s += v
	}
	return s
}

func join(xs []string) string {
	return strings.Join(xs, " ")
}
