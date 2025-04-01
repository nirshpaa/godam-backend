package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/nirshpaa/godam-backend/libraries/api"
	"github.com/nirshpaa/godam-backend/models"
)

// ClosingStocks : struct for set ClosingStocks Dependency Injection
type ClosingStocks struct {
	Db  *sql.DB
	Log *log.Logger
}

// Closing : http handler for closing stock
func (u *ClosingStocks) Closing(w http.ResponseWriter, r *http.Request) {
	var closingStock models.ClosingStock

	err := closingStock.Closing(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		api.ResponseError(w, fmt.Errorf("closing stock: %v", err))
		return
	}

	api.ResponseOK(w, nil, http.StatusOK)
}
