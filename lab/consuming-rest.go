package lab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	url = "https://jsonplaceholder.typicode.com/users"
)

type (
	Users []struct {
		ID       int     `json:"id"`
		Name     string  `json:"name"`
		Username string  `json:"username"`
		Email    string  `json:"email"`
		Address  Address `json:"address"`
		Phone    string  `json:"phone"`
		Website  string  `json:"website"`
		Company  Company `json:"company"`
	}

	Address struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     string `json:"geo"`
	}

	Geo struct {
		Lat string `json:"lat"`
		Lng string `json:"lng"`
	}

	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	}
)

func ConsumingRest() {
	println("Consuming Rest")

	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	reponseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var reponseObject Users
	json.Unmarshal(reponseData, &reponseObject)

	for index, value := range reponseObject {
		fmt.Println(index, value)
	}
}
