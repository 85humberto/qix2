package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	rand.Seed(int64(time.Now().Nanosecond()))
	t := []string{"Tranferência", "Receber", "Pagar"}

	for {
		q := Qix{t[rand.Intn((len(t)))], 1 + rand.Intn(10)}
		// fmt.Println(q)
		jsonReq, err := json.Marshal(q)
		if err != nil {
			log.Fatalln(err)
		}
		_, err = http.Post("http://localhost:8080/qix", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
		if err != nil {
			log.Fatalln(err)
		} else {
			fmt.Println("Transação enviada: ", q)
		}

		time.Sleep(1 * time.Second)
	}
}
