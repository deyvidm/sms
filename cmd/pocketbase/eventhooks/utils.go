package eventhooks

import (
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

func (ehs *EventhookStore) saveRecord(r *models.Record) error {
	return ehs.app.Dao().RunInTransaction(func(txDao *daos.Dao) error {
		return txDao.SaveRecord(r)
	})
}
