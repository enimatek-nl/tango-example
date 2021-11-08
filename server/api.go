package server

import (
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
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

func (e *Api) Process(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("Error parsing form: %s\n", err)
		return
	}

	id := 0
	if i, err := strconv.Atoi(r.Form.Get("id")); err == nil {
		id = i
	}

	switch r.Method {
	case http.MethodPost:
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		defer r.Body.Close()
		decoder := json.NewDecoder(r.Body)
		var todo Todo
		err := decoder.Decode(&todo)

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 400)
			return
		}
		e.Db.Save(&todo)
	case http.MethodDelete:
		if id != 0 {
			e.Db.Delete(&Todo{}, "id = ?", id)
		}
	case http.MethodGet:
		if id == 0 {
			var todos []Todo
			e.Db.Model(&Todo{}).Find(&todos)
			json.NewEncoder(w).Encode(todos)
		} else {
			var todo Todo
			e.Db.Model(&Todo{}).Where("id = ?", id).Find(&todo)
			json.NewEncoder(w).Encode(todo)
		}
	}
}
