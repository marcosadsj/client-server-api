package report

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var Buffer *bufio.Writer

func Open() {
	file, err := os.Create("cotacao.txt")

	if err != nil {
		log.Fatalln(err)
	}

	Buffer = bufio.NewWriter(file)
}

func InsertDollarPrice(dollarPrince string) {

	Buffer.WriteString(fmt.Sprintf("Dólar: %s", dollarPrince))

	Buffer.Flush()
}
