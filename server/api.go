package server

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

type Api struct {
	Db *gorm.DB
}

func (e *Api) Index(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(
		struct{ Message string }{
			Message: "test reply from api",
		},
	)
}

func (e *Api) Get(w http.ResponseWriter, req *http.Request) {
	todos := []Todo{{
		Title:   "Abc",
		Content: "Def",
		Done:    false,
	}, {
		Title:   "Def",
		Content: "Ghi",
		Done:    false,
	}}
	json.NewEncoder(w).Encode(todos)
}

func (e *Api) Post(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	e.Db.Save(&todo)
}
