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
