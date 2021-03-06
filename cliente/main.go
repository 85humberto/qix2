package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Qix struct {
	Transação    string `json:"transacao"`
	Complexidade int    `json:"complexidade"`
}

func main() {
	//Semente para o rand
	rand.Seed(int64(time.Now().Nanosecond()))
	t := []string{"Tranferência", "Receber", "Pagar"}

	for {
		// Sorteia um tipo de transferência e um valor de complexidade
		q := Qix{t[rand.Intn((len(t)))], 1 + rand.Intn(10)}
		log.Println("Transação gerada: ", q)
		// Codifica json
		jsonReq, err := json.Marshal(q)
		if err != nil {
			log.Fatalln(err)
		}
		_, err = http.Post("http://balanceador:8080/qix", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
		if err != nil {
			log.Println(err)
		}
		// Espera 2 segundo
		time.Sleep(2 * time.Second)
	}
}
