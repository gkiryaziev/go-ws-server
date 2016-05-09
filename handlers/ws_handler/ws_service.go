package wshandler

import (
	"sync"

	"github.com/jmoiron/sqlx"
)

var mutex sync.RWMutex

type wsService struct {
	db *sqlx.DB
}

func newWSService(db *sqlx.DB) *wsService {
	return &wsService{db}
}

// get all subscribers by topic name.
func (wss *wsService) getSubscribers(topicName string) []string {
	var uids []string
	err := wss.db.Select(&uids, "select uid from subscribers where topic_id = (select id from topics where name = ?)",
		topicName)
	if err != nil {
		panic(err)
	}
	return uids
}

// subscribe to new topic.
func (wss *wsService) subscribe(topicName, uid string) error {

	// check if topic exist
	var topicID int
	err := wss.db.Get(&topicID, "select id from topics where name = ?", topicName)
	if err != nil {
		topicID = 0
	}

	if topicID == 0 {
		result, err := wss.db.NamedExec("insert into topics(name) values(:name)",
			map[string]interface{}{"name": topicName})
		if err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		_, err = wss.db.NamedExec("insert into subscribers(uid, topic_id) values(:uid, :id)",
			map[string]interface{}{"uid": uid, "id": id})
		if err != nil {
			return err
		}
	}

	if topicID > 0 {

		// check if subscriber exist
		var subscriberID int
		err := wss.db.Get(&subscriberID, "select id from subscribers where uid = ? and topic_id = ?", uid, topicID)
		if err != nil {
			subscriberID = 0
		}

		if subscriberID == 0 {
			_, err = wss.db.NamedExec("insert into subscribers(uid, topic_id) values(:uid, :id)",
				map[string]interface{}{"uid": uid, "id": topicID})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// unSubscribe from topic.
func (wss *wsService) unSubscribe(topicName, uid string) error {
	_, err := wss.db.NamedExec("delete from subscribers where uid = :uid and topic_id = (select id from topics where name = :name)",
		map[string]interface{}{"uid": uid, "name": topicName})
	if err != nil {
		return err
	}
	return nil
}

// unSubscribeAll unsubscribe from all topic.
func (wss *wsService) unSubscribeAll(uid string) error {
	_, err := wss.db.NamedExec("delete from subscribers where uid = :uid",
		map[string]interface{}{"uid": uid})
	if err != nil {
		return err
	}
	return nil
}

// addLog is add log.
func (wss *wsService) addLog(uid, remoteAddress, message string) error {
	mutex.Lock()
	defer mutex.Unlock()

	_, err := wss.db.NamedExec("insert into logs(uid, remote_address, message) values(:uid, :remote_address, :message)",
		map[string]interface{}{"uid": uid, "remote_address": remoteAddress, "message": message})
	if err != nil {
		return err
	}
	return nil
}
