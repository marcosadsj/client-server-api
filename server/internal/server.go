package internal

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"server-api/internal/database"
	"server-api/internal/database/model"
	"server-api/internal/dto"
	"time"
)

type USDBRL struct {
	USDBRL struct {
		Code        string `json:"code"`
		Codein      string `json:"codein"`
		Name        string `json:"name"`
		High        string `json:"high"`
		Low         string `json:"low"`
		VarBid      string `json:"varBid"`
		PctChange   string `json:"pctChange"`
		Bid         string `json:"bid"`
		Ask         string `json:"ask"`
		Timestamp   string `json:"timestamp"`
		Create_Date string `json:"create_date"`
	}
}

func Start() {

	mux := http.NewServeMux()

	mux.HandleFunc("/cotacao", handleGetDollar)

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatalln(err)
	}
}

func handleGetDollar(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)

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

	var bodyJson USDBRL

	err = json.Unmarshal(body, &bodyJson)

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Server: %v", bodyJson)

	database.InsertDollar(&model.DollarPrice{Price: bodyJson.USDBRL.Bid, Created_at: time.Now()})

	dollarBytes, err := json.Marshal(&dto.DollarPrice{Bid: bodyJson.USDBRL.Bid})

	if err != nil {
		log.Fatalln(err)
	}

	_, err = w.Write(dollarBytes)

	if err != nil {
		log.Fatalln(err)
	}
}
