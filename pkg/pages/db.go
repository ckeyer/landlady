package pages

import (
	"github.com/funxdata/commons/mongo"
	mgo "gopkg.in/mgo.v2"
)

const (
	ColPages = "pages"
)

func setupIndex(db *mgo.Database) error {
	indexes := &mongo.MgoIndexs{
		ColPages: []mgo.Index{
		// mgo.Index{
		// 	Key:    []string{"name"},
		// 	Unique: true,
		// },
		},
	}

	return indexes.Setup(db)
}
