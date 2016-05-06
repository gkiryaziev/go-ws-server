package log_handler

import (
	"github.com/jmoiron/sqlx"

	"github.com/gkiryaziev/go-ws-server/models"
	"github.com/gkiryaziev/go-ws-server/utils"
)

type logService struct {
	db *sqlx.DB
}

// newLogService return logService object.
func newLogService(db *sqlx.DB) *logService {
	return &logService{db}
}

// getLogs return all logs.
func (ls *logService) getLogs() *utils.ResultTransformer {

	logs := []LogModel{}

	err := ls.db.Select(&logs, "select * from logs order by id asc limit 500")
	if err != nil {
		panic(err)
	}

	header := models.Header{"ok", len(logs), logs}
	result := utils.NewResultTransformer(header)

	return result
}

// truncateTable truncate table.
func (ls *logService) truncateTable(table string) error {
	_, err := ls.db.NamedExec("delete from :table", map[string]interface{}{"table": table})
	if err != nil {
		return err
	}

	_, err = ls.db.Exec("VACUUM")
	if err != nil {
		return err
	}

	return nil
}
