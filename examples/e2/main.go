package main

import (
	"fmt"

	"github.com/dragmz/tqr"
)

func main() {
	qr := tqr.New("https://github.com/dragmz/tqr", tqr.Invert())
	fmt.Println(qr)
}
