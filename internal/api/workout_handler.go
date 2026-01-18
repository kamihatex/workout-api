package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"workoutGo/internal/store"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct {
	workoutStore store.WorkoutStore
	logger       *log.Logger
}

func NewWorkoutHandler(workoutStore store.WorkoutStore, logger *log.Logger) *WorkoutHandler {
	return &WorkoutHandler{
		workoutStore: workoutStore,
		logger:       logger,
	}
}

func (wh *WorkoutHandler) HandleGetWorkoutById(w http.ResponseWriter, r *http.Request) {
	paramsWorkoutID := chi.URLParam(r, "id")
	if paramsWorkoutID == "" {
		http.NotFound(w, r)
		return
	}

	workoutID, err := strconv.ParseInt(paramsWorkoutID, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "this is the workout id %d\n", workoutID)
}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	var workout store.Workout
	err := json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		wh.logger.Printf("ERROR: decodingCreateWorkout: %v", err)
		http.Error(w, "failed to create workout", http.StatusInternalServerError)
		//utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request sent"})
		return
	}

	createdWorkout, err := wh.workoutStore.CreateWorkout(&workout)
	if err != nil {
		wh.logger.Printf("ERROR: createWorkout: %v", err)
		http.Error(w, "failed to create workout", http.StatusInternalServerError)
		//utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "failed to create workout"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdWorkout)
}
