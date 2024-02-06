package calculator

import (
	"encoding/json"
	"net/http"
	"strings"

	"specialist/internal/models"
	"specialist/internal/utils"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	utils.SetOKResult(w)
	result := Result(r.RequestURI)
	err := json.NewEncoder(w).Encode(
		models.APIResponse{Result: result},
	)
	if err != nil {
		return
	}
}

func Result(uri string) int {
	var number int
	splitPath := strings.Split(uri, "/")
	method := splitPath[len(splitPath)-1]

	switch {
	case method == "first" || method == "second":
		number = utils.GetRandomNumber()
	case method == "add":
		number = utils.GetRandomNumber() + utils.GetRandomNumber()
	case method == "sub":
		number = utils.GetRandomNumber() - utils.GetRandomNumber()
	case method == "mul":
		number = utils.GetRandomNumber() * utils.GetRandomNumber()
	case method == "div":
		number = utils.GetRandomNumber() / utils.GetRandomNumber()
	}
	return number
}
