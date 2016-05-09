package loghandler

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/gkiryaziev/go-ws-server/utils"
)

// LogController struct
type LogController struct {
	service *LogService
}

// NewLogController return new LogController object.
func NewLogController(db *sqlx.DB) *LogController {
	return &LogController{newLogService(db)}
}

// GetLogs return all logs.
func (lc *LogController) GetLogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	logs, err := lc.service.getLogs().ToJSON()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, utils.ErrorMessage(500, err.Error()))
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, logs)
}
