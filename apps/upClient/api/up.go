package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const UrlBase = "https://api.up.com.au/api/v1/"

const TransactionsEndpoint = "transactions"
const PingEndpoint = "util/ping"

const AccountsEndpoint = "accounts"

type MoneyObject struct {
	CurrencyCode string `json:"currencyCode"`
	Value        string `json:"value"`
	ValueAsBase  int64  `json:"valueInBaseUnits"`
}

type TransactionResponse struct {
	Type string `json:"type"`
	Id   string `json:"id"`

	Attributes struct {
		// TODO: A lot missing in here

		Description string      `json:"description"`
		Amount      MoneyObject `json:"amount"`
	} `json:"attributes"`
}

type Transaction struct {
	Data  []TransactionResponse `json:"data"`
	Links struct {
		Prev string `json:"prev"`
		Next string `json:"next"`
	} `json:"links"`
}

func MakeTransactionListRequest(token string, accountId string) (Transaction, error) {

	client := &http.Client{}

	url := UrlBase + AccountsEndpoint + "/" + accountId + "/" + TransactionsEndpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return Transaction{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return Transaction{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
		return Transaction{}, err
	}

	result := Transaction{}
	json.Unmarshal(body, &result)

	return result, nil
}
