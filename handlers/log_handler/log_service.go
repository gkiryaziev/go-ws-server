package loghandler

import (
	"github.com/jmoiron/sqlx"

	"github.com/gkiryaziev/go-ws-server/models"
	"github.com/gkiryaziev/go-ws-server/utils"
)

// LogService struct
type LogService struct {
	db *sqlx.DB
}

// newLogService return LogService object.
func newLogService(db *sqlx.DB) *LogService {
	return &LogService{db}
}

// getLogs return all logs.
func (ls *LogService) getLogs() *utils.ResultTransformer {

	logs := []LogModel{}

	err := ls.db.Select(&logs, "select * from logs order by id asc limit 500")
	if err != nil {
		panic(err)
	}

	header := models.Header{Status: "ok", Count: len(logs), Data: logs}
	result := utils.NewResultTransformer(header)

	return result
}

// truncateTable truncate table.
func (ls *LogService) truncateTable(table string) error {
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
