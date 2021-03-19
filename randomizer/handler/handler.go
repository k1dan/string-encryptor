package handler

import (
	"encoding/json"
	"github.com/k1dan/string-encryptor/randomizer/encryptor"
	string_creator "github.com/k1dan/string-encryptor/randomizer/string-creator"
	"net/http"
	"strconv"
)

type Number struct {
	N string `json:"n"`
}

type Strings struct {
	S []string	`json:"s"`
}

func CreateStrings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var number Number
	var strings Strings
	_ = json.NewDecoder(r.Body).Decode(&number)
	n, _ := strconv.Atoi(number.N)
	for i:=0; i< n; i++ {
		s := string_creator.RandStringRunes()
		strings.S = append(strings.S, s)
	}
	str := strings.S
	encr := encryptor.Encrypt(str)
	json.NewEncoder(w).Encode(encr)
}



