package main

import "unicode"

func FindCommon(s string) rune {
	first := s[:len(s)/2]
	second := s[len(s)/2:]
	for _, r := range first {
		for _, r2 := range second {
			if r == r2 {
				return r
			}
		}
	}
	return ' '
}

func FindCommons(s string, s2 string) string {
	commons := ""
	for _, r := range s {
		for _, r2 := range s2 {
			if r == r2 {
				commons += string(r)
			}
		}
	}
	return commons
}

func ValueRune(r rune) int {
	directory := "abcdefghijklmnopqrstuvwxyz"
	for i, val := range directory {
		if val == unicode.ToLower(r) {
			if r != unicode.ToLower(r) {
				return i + 27
			}
			return i + 1
		}
	}
	return 0
}
