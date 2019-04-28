package main

import (
	"encoding/csv"
	//"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

  file,_ := os.Open("file.csv")

	lines,_ := csv.NewReader(file).ReadAll()

	correct := 0
	for i, line := range lines {
		fmt.Printf("Problem #%d: %s = ", i+1, line[0])
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == strings.TrimSpace(line[1]) {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(lines))
}
