package wallet

import (
	"carbon/internal/database"
	"carbon/internal/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetWallet(w http.ResponseWriter, r *http.Request) {

	userID := chi.URLParam(r, "userID")

	id, _ := strconv.Atoi(userID)

	var activities []models.Activity

	database.DB.
		Where("user_id = ? AND status = ?", id, "Approved").
		Find(&activities)

	total := 0

	for _, activity := range activities {
		total += activity.Credits
	}

	json.NewEncoder(w).Encode(map[string]any{

		"user_id": userID,

		"total_credits": total,

		"activities": activities,
	})

}
