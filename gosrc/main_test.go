package main

import (
	"fmt"
	"testing"
)

func Test_getCurrentETHPrice(t *testing.T) {
	t.Run("Valid test", func(t *testing.T) {
		price := getCurrentETHPrice()
		fmt.Println(price)
	})
}