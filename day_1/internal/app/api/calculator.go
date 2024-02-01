package api

import (
	"encoding/json"
	"net/http"

	"day_1/internal/models"
	utils2 "day_1/internal/utils"
)

func CalculatorHandler(w http.ResponseWriter, r *http.Request) {
	utils2.SetOKResult(w)
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
		number = utils2.GetRandomNumber()
	case uri == "/add":
		number = utils2.GetRandomNumber() + utils2.GetRandomNumber()
	case uri == "/sub":
		number = utils2.GetRandomNumber() - utils2.GetRandomNumber()
	case uri == "/mul":
		number = utils2.GetRandomNumber() * utils2.GetRandomNumber()
	case uri == "/div":
		number = utils2.GetRandomNumber() / utils2.GetRandomNumber()
	}
	return number
}
