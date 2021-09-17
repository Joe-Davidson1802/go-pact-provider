package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/joe-davidson1802/go-pact-provider/models"
)

func GetTimeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	est, _ := time.LoadLocation("America/Cancun")
	times := []models.TimeResponse{
		{
			Zone:   "EST",
			Time:   time.Now().In(est),
			Offset: "-05",
		},
		{
			Zone:   "UTC",
			Time:   time.Now().UTC(),
			Offset: "+00",
		},
	}
	json.NewEncoder(w).Encode(times)
}
