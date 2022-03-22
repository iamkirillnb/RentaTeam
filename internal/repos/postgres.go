package repos

import (
	"github.com/iamkirillnb/Rentateam/internal/entities"
	"github.com/iamkirillnb/Rentateam/pkg/logging"
	"github.com/jmoiron/sqlx"
	"time"
)

type DbLaw struct {
	DB     *sqlx.DB
	Logger logging.Logger
}

func NewDbLaw(db *sqlx.DB, logger logging.Logger) DbLaw {
	return DbLaw{
		DB:     db,
		Logger: logger,
	}
}

func (d *DbLaw) GetAll() ([]*entities.FormData, error) {
	const qry = `select * from data`

	data := []*entities.FormData{}
	err := d.DB.Select(&data, qry)
	if err != nil {
		d.Logger.Info("select all data from DB fail")
		return nil, err
	}
	return data, nil
}

func (d *DbLaw) WriteData(data *entities.FormData) error {
	if data.Tag == "" {
		data.Tag = entities.DefaultTeg
	}
	data.CreatedAt = time.Now()

	const qry = `insert into data (title, text, tag, created_at) values (:title, :text, :tag, :created_at)`

	_, err := d.DB.NamedExec(qry, data)
	if err != nil {
		d.Logger.Info("write data to DB fail")
		return err
	}
	return nil
}
