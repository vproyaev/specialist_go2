package api

import (
	"encoding/json"
	"net/http"

	"day_1/models"
	"day_1/utils"
)

func CalculatorHandler(w http.ResponseWriter, r *http.Request) {
	utils.SetOKResult(w)
	result := CalculateResult(r.RequestURI)
	err := json.NewEncoder(w).Encode(
		models.APIResponse{Result: result},
	)
	if err != nil {
		return
	}
}

func CalculateResult(uri string) int {
	var number int
	switch {
	case uri == "/first" || uri == "/second":
		number = utils.GetRandomNumber()
	case uri == "/add":
		number = utils.GetRandomNumber() + utils.GetRandomNumber()
	case uri == "/sub":
		number = utils.GetRandomNumber() - utils.GetRandomNumber()
	case uri == "/mul":
		number = utils.GetRandomNumber() * utils.GetRandomNumber()
	case uri == "/div":
		number = utils.GetRandomNumber() / utils.GetRandomNumber()
	}
	return number
}
