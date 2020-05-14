package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

var currentPos = 0
var spaceProbability = 80

func main() {
	var lastLine = make([]string, 100)
	green := color.New(color.FgGreen).Add(color.BgBlack)

	for {
		var currentLine []string

		for currentPos < 100 {
			if lastLine[currentPos] == " " {
				spaceProbability = 90
			} else {
				spaceProbability = 10
			}
			if rand.Intn(100) < spaceProbability {
				currentLine = append(currentLine, " ")
				currentPos++
			} else {
				currentLine = append(currentLine, fmt.Sprintf("%c", rand.Intn(127-33)+33))
				currentPos++
			}
		}
		currentPos = 0

		_, _ = green.Println(currentLine)
		time.Sleep(50 * time.Millisecond)
		lastLine = currentLine
	}
}
