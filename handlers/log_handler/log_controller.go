package log_handler

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/gkiryaziev/go-ws-server/utils"
)

type logController struct {
	service *logService
}

// NewLogController return new logController object.
func NewLogController(db *sqlx.DB) *logController {
	return &logController{newLogService(db)}
}

// GetLogs return all logs.
func (lc *logController) GetLogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	logs, err := lc.service.getLogs().ToJson()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, utils.ErrorMessage(500, err.Error()))
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, logs)
}
