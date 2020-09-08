package __kyu

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`(?mi)[^A-Za-z0-9\s\.]`)

func Balance(book string) string {
	var r []string
	book = re.ReplaceAllString(book, "")

	scanner := bufio.NewScanner(strings.NewReader(book))
	var (
		balance, total, average float64
		err                     error
	)
	i := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		if i == 0 {
			balance, err = strconv.ParseFloat(text, 32)
			if err != nil {
				log.Fatal(err)
			}
			r = append(r, fmt.Sprintf("Original Balance: %.2f", balance))
		} else {
			texts := strings.Split(text, " ")
			cost, err := strconv.ParseFloat(texts[2], 32)
			if err != nil {
				log.Fatal(err)
			}
			text = fmt.Sprintf("%s %s %.2f", texts[0], texts[1], cost)
			total += cost
			balance = balance - cost
			r = append(r, fmt.Sprintf("%s Balance %.2f", text, balance))
		}
		i++
	}
	average = total / float64(len(r)-1)
	r = append(r, fmt.Sprintf("Total expense  %.2f", total), fmt.Sprintf("Average expense  %.2f", average))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return strings.Join(r, "\n")
}
