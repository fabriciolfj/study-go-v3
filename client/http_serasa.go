package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"fabriciolfj.github/study/models"
)

func Request() {
	r, err := http.Get("http://localhost:9000/api/serasa/v1/2")
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var customer models.Customer
	err = json.Unmarshal(body, &customer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", customer)
}
