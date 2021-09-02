package util

import (
	"log"
	"testing"
)

func Test_reverse(t *testing.T) {
	str := "abcdefg"
	str = reverseStr([]byte(str))
	log.Println(str)
}
