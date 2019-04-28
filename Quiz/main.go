package main

import (
	"encoding/csv"
	//"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	file,_ := os.Open("file.csv")
	lines,_ := csv.NewReader(file).ReadAll()

	correct := 0
	for i, line := range lines {
		fmt.Printf("Problem #%d: %s = ", i+1, line[0])
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
    timer := time.NewTimer(5 * time.Second)
		select {
		case <-timer.C:
			fmt.Println()
			continue
		case answer := <-answerCh:
			if answer == strings.TrimSpace(line[1]) {
				correct++
			}
      continue
		}
	} 
	fmt.Printf("You scored %d out of %d.\n", correct, len(lines))
}
