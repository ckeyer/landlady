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
			mgo.Index{
				Key: []string{"url.originurl"},
			},
			mgo.Index{
				Key: []string{"url.realurl"},
			},
			mgo.Index{
				Key: []string{"url.clearurl"},
			},
			mgo.Index{
				Key: []string{"handleat"},
			},
		},
	}

	return indexes.Setup(db)
}
