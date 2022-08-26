package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   uint   `json:"age"`
}

func Marshal() {
	Tobias := dog{"Tobias", "Dashchund", 7}
	Pacoca := dog{"Paçoca", "SDR", 4}

	tobiJSON, err := json.Marshal(Tobias)
	pacocaJSON, err := json.Marshal(Pacoca)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(tobiJSON))
	fmt.Println(bytes.NewBuffer(pacocaJSON))
}
