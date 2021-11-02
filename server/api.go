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
