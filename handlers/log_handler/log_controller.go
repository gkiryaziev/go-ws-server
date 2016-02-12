package log_handler

import (
	"fmt"
	"net/http"

	"../../utils"

	"github.com/jmoiron/sqlx"
)

type logController struct {
	service *logService
}

func NewLogController(db *sqlx.DB) *logController {
	return &logController{newLogService(db)}
}

// get logs.
func (this *logController) GetLogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	logs, err := this.service.getLogs().ToJson()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, utils.ErrorMessage(500, err.Error()))
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, logs)
}
