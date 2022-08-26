package json

import (
	"encoding/json"
	"fmt"
	"log"
)

func Unmarshal() {
	tobiJSON := `{"name":"Tobias", "breed":"Dashchund", "age":7}`
	pacocaJSON := `{"name":"Paçoca", "breed":"SDR", "age":4}`

	var Tobias dog
	var Pacoca dog

	if err := json.Unmarshal([]byte(pacocaJSON), &Pacoca); err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(tobiJSON), &Tobias); err != nil {
		log.Fatal(err)
	}

	fmt.Println(Tobias)
	fmt.Println(Pacoca)
}
