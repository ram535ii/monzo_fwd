package main

import (
  "encoding/json"
  "fmt"
  "os"
  "bytes"
  "net/http"

  // "io/ioutil"

  "github.com/apex/go-apex"
)


type Transaction struct {
  AccountId   string `json:"account_id"`
  Amount      int    `json:"amount"`
  Description string `json:"description"`
  Currency    string `json:"currency"`
}
type TransactionEvent struct {
  Type        string `json:"type"`
  Transaction        `json:"data"`
}

type Container struct {
  Body        string `json:"body"`
}

func main() {
  apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
    fmt.Fprintf(os.Stderr, "EVENT!\n\n%s\n\n", event)

    var container Container
    var transactionEvent TransactionEvent

    if err := json.Unmarshal(event, &container); err != nil {
      return nil, err
    }

    if err := json.Unmarshal([]byte(container.Body), &transactionEvent); err != nil {
      return nil, err
    }

    telegramBot := os.Getenv("telegram_bot")
    telegramRecipient := os.Getenv("telegram_recipient")

    var jsonStr = []byte(fmt.Sprintf(`{"chat_id": "%v", "text": "%v - %v"}`, telegramRecipient, transactionEvent.Transaction.Amount, transactionEvent.Transaction.Description))

    url := fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage", telegramBot)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
      panic(err)
    }
    defer resp.Body.Close()

    fmt.Fprintf(os.Stderr, "RESPONSE: %s", resp.Status)
    // body, _ := ioutil.ReadAll(resp.Body)
    // fmt.Println("response Body:", string(body))

    return container, nil
  })
}
