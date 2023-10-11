package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
)

func main() {
	f, err := os.Open("m.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rows := make([]string, 0)
	scanner := bufio.NewScanner(f)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	for scanner.Scan() {
		r := scanner.Text()
		if len(r) > 50 {
			r = r[50:]
		}
		rows = append(rows, r)

		if len(rows) >= 150000 {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	m := make(map[string]int)
	for _, row := range rows {
		found := false

		for k, v := range m {
			similarity := strutil.Similarity(row, k, metrics.NewHamming())

			if similarity >= 0.5 {
				m[k] = v + 1
				found = true
				break
			}
		}

		if !found {
			m[row] = 1
		}
	}

	for k, v := range m {
		fmt.Printf("%d;%s\n", v, k)
	}
}
