package activities

import (
	"carbon/internal/database"
	"carbon/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-chi/chi/v5"
)

type CreateActivityRequest struct {
	UserID      uint   `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type VerifyRequest struct {
	Status string `json:"status"`
}

// POST /api/v1/activities
func CreateActivity(w http.ResponseWriter, r *http.Request) {

	var req CreateActivityRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	activity := models.Activity{
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		Type:        req.Type,
		Status:      "Pending",
		Credits:     0,
	}

	if err := database.DB.Create(&activity).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(activity)
}

// GET /api/v1/activities
func GetActivities(w http.ResponseWriter, r *http.Request) {

	var activities []models.Activity

	if err := database.DB.Find(&activities).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activities)
}

// GET /api/v1/activities/{id}
func GetActivityByID(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	var activity models.Activity

	if err := database.DB.First(&activity, id).Error; err != nil {
		http.Error(w, "Activity not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activity)
}

// PUT /api/v1/activities/{id}
func UpdateActivity(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	var activity models.Activity

	if err := database.DB.First(&activity, id).Error; err != nil {
		http.Error(w, "Activity not found", http.StatusNotFound)
		return
	}

	var req CreateActivityRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	activity.Title = req.Title
	activity.Description = req.Description
	activity.Type = req.Type

	if err := database.DB.Save(&activity).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(activity)
}

// DELETE /api/v1/activities/{id}
func DeleteActivity(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	if err := database.DB.Delete(&models.Activity{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func VerifyActivity(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	var activity models.Activity

	if err := database.DB.First(&activity, id).Error; err != nil {
		http.Error(w, "Activity not found", http.StatusNotFound)
		return
	}

	var req VerifyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	activity.Status = req.Status

	if req.Status == "Approved" {

		switch activity.Type {

		case "Tree Plantation":
			activity.Credits = 10

		case "Recycling":
			activity.Credits = 5

		case "Solar":
			activity.Credits = 50

		default:
			activity.Credits = 1
		}

	} else {

		activity.Credits = 0
	}

	if err := database.DB.Save(&activity).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(activity)
}

func UploadProof(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("proof")
	if err != nil {
		http.Error(w, "File not found", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(header.Filename))

	dst, err := os.Create("./uploads/" + filename)
	if err != nil {
		http.Error(w, "Could not save file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Upload successful",
		"file":    filename,
		"url":     "/uploads/" + filename,
	})
}
