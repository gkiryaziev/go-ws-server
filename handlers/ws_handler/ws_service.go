package ws_handler

import (
	"github.com/jmoiron/sqlx"
)

type wsService struct {
	db *sqlx.DB
}

func newWSService(db *sqlx.DB) *wsService {
	return &wsService{db}
}

// ==========================
// get all subscribers by topic name
// ==========================
func (this *wsService) getSubscribers(topicName string) []string {
	var uids []string
	err := this.db.Select(&uids, "select uid from subscribers where topic_id = (select id from topics where name = ?)",
		topicName)
	if err != nil {
		panic(err)
	}
	return uids
}

// ==========================
// subscribe to new topic
// ==========================
func (this *wsService) subscribe(topicName, uid string) error {

	// check if topic exist
	var topic_id int
	err := this.db.Get(&topic_id, "select id from topics where name = ?", topicName)
	if err != nil {
		topic_id = 0
	}

	if topic_id == 0 {
		result, err := this.db.NamedExec("insert into topics(name) values(:name)",
			map[string]interface{}{"name": topicName})
		if err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		_, err = this.db.NamedExec("insert into subscribers(uid, topic_id) values(:uid, :id)",
			map[string]interface{}{"uid": uid, "id": id})
		if err != nil {
			return err
		}
	}

	if topic_id > 0 {

		// check if subscriber exist
		var subscriber_id int
		err := this.db.Get(&subscriber_id, "select id from subscribers where uid = ? and topic_id = ?", uid, topic_id)
		if err != nil {
			subscriber_id = 0
		}

		if subscriber_id == 0 {
			_, err = this.db.NamedExec("insert into subscribers(uid, topic_id) values(:uid, :id)",
				map[string]interface{}{"uid": uid, "id": topic_id})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// ==========================
// unsubscribe from topic
// ==========================
func (this *wsService) unSubscribe(topicName, uid string) error {
	_, err := this.db.NamedExec("delete from subscribers where uid = :uid and topic_id = (select id from topics where name = :name)",
		map[string]interface{}{"uid": uid, "name": topicName})
	if err != nil {
		return err
	}
	return nil
}

// ==========================
// unsubscribe from all topic
// ==========================
func (this *wsService) unSubscribeAll(uid string) error {
	_, err := this.db.NamedExec("delete from subscribers where uid = :uid",
		map[string]interface{}{"uid": uid})
	if err != nil {
		return err
	}
	return nil
}

// ==========================
// add log
// ==========================
func (this *wsService) addLog(uid, remote_address, message string) error {
	_, err := this.db.NamedExec("insert into logs(uid, remote_address, message) values(:uid, :remote_address, :message)",
		map[string]interface{}{"uid": uid, "remote_address": remote_address, "message": message})
	if err != nil {
		return err
	}
	return nil
}
