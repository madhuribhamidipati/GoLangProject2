
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	Stuff Data
}

type Data struct {
	Fruit   Fruits
	Veggies Vegetables
}

type Fruits map[string]int
type Vegetables map[string]int

func serveRest(w http.ResponseWriter, r *http.Request) {
	response, err := getJsonResponse()
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(response))
}

func main() {
	http.HandleFunc("/", serveRest)
	http.ListenAndServe("localhost:1337", nil)

}

func getJsonResponse() ([]byte, error) {
	Fruits := make(map[string]int)
	Fruits["Apples"] = 25
	Fruits["Oranges"] = 11
	Vegetables := make(map[string]int)
	Vegetables["Carrots"] = 21
	Vegetables["Peppers"] = 0

	d := Data{Fruits, Vegetables}
	p := Payload{d}

	return json.MarshalIndent(p, "", "  ")
}
