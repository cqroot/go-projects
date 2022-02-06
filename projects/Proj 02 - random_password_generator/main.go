package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	lowerCharSet   = "abcdefghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRST"
	specialCharSet = "!@#$%"
	numberSet      = "0123456789"
)

func main() {
	fmt.Println(Generate(
		16, true, true, true, false,
	))
}

func Generate(length int, useLowerCharSet bool, useUpperCharSet bool, useNumberSet bool, useSpecialCharSet bool) string {
	rand.Seed(time.Now().Unix())

	charSet := []rune{}
	if useLowerCharSet {
		charSet = append(charSet, []rune(lowerCharSet)...)
	}
	if useUpperCharSet {
		charSet = append(charSet, []rune(upperCharSet)...)
	}
	if useNumberSet {
		charSet = append(charSet, []rune(numberSet)...)
	}
	if useSpecialCharSet {
		charSet = append(charSet, []rune(specialCharSet)...)
	}

	var passwordBuilder strings.Builder
	for i := 0; i < length; i++ {
		passwordBuilder.WriteRune(charSet[rand.Intn(len(charSet))])
	}
	return passwordBuilder.String()
}
