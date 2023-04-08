package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Data struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	for {
		post()
		time.Sleep(15 * time.Second)
	}
}

func post() {
	data := Data{
		Water: rand.Intn(100) + 1,
		Wind:  rand.Intn(100) + 1,
	}

	jsonBody, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatalf("Failed to POST data: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	fmt.Printf("%s\n", responseBody)

	statusWater(data.Water)
	statusWind(data.Wind)
	fmt.Println(strings.Repeat("#", 25))
}

func statusWater(water int) {
	if water < 5 {
		fmt.Println("status water : aman")
	} else if water >= 6 && water <= 8 {
		fmt.Println("status water : siaga")
	} else {
		fmt.Println("status water : bahaya")
	}
}

func statusWind(wind int) {
	if wind < 6 {
		fmt.Println("status wind : aman")
	} else if wind >= 7 && wind <= 15 {
		fmt.Println("status wind : siaga")
	} else {
		fmt.Println("status wind : bahaya")
	}
}
