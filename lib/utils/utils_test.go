package utils

import (
	"fmt"
	"strings"
	"testing"
)

func TestCalculateHash(t *testing.T) {
	r := strings.NewReader("this object will have only 1 instance")
	s := CalculateHash(r)
	fmt.Println(s)
}
