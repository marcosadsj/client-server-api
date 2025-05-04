package client

import (
	report "client-server-api/client/report"
	"client-server-api/dto"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func Start() {

	report.Open()

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		log.Fatalln(err)
	}

	var dollar dto.DollarPrice

	err = json.Unmarshal(body, &dollar)

	if err != nil {
		log.Fatalln(err)
	}

	report.InsertDollarPrice(dollar.Bid)

	log.Printf("Client: Cotação Dollar: %s", dollar.Bid)
}
