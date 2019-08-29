package main

import (
	"os"
	"time"

	"github.com/maxjw/golang-tutorial/maths/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
