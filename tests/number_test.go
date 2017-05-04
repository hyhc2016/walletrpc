package tests

import (
	"testing"
	"github.com/hyhc2016/walletrpc/core"
	"log"
	"fmt"
	"strconv"
)

func TestNumber(t *testing.T) {
	if a4, err := core.Add(999999.99999999, 0.99999997); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("a4 = ", strconv.FormatFloat(a4, 'f', -1, 64))
	}
}
