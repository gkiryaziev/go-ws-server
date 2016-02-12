package log_handler

import (
	"../../models"
	"../../utils"

	"github.com/jmoiron/sqlx"
)

type logService struct {
	db *sqlx.DB
}

func newLogService(db *sqlx.DB) *logService {
	return &logService{db}
}

// get logs.
func (this *logService) getLogs() *utils.ResultTransformer {

	logs := []LogModel{}

	err := this.db.Select(&logs, "select * from logs order by id asc limit 500")
	if err != nil {
		panic(err)
	}

	header := models.Header{"ok", len(logs), logs}
	result := utils.NewResultTransformer(header)

	return result
}

// truncate table.
func (this *logService) truncateTable(table string) error {
	_, err := this.db.NamedExec("delete from :table", map[string]interface{}{"table": table})
	if err != nil {
		return err
	}

	_, err = this.db.Exec("VACUUM")
	if err != nil {
		return err
	}

	return nil
}
