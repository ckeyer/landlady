package pages

import (
	pb "github.com/funxdata/landlady/proto"
	"github.com/gogo/protobuf/types"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	mgo "gopkg.in/mgo.v2"
)

var _ pb.PagesServer = (*PagesServer)(nil)

type PagesServer struct {
	basedb *mgo.Database
}

func NewPagesServer(db *mgo.Database) *PagesServer {
	err := setupIndex(db)
	if err != nil {
		logrus.Errorf("setup index failed, %s", err)
	}

	return &PagesServer{
		basedb: db,
	}
}

// db
func (p *PagesServer) DB(ctx context.Context) *mgo.Database {
	return p.basedb
	// return p.basedb.Session.Clone().DB(p.basedb.Name)
}

// Save
func (p *PagesServer) Save(ctx context.Context, in *pb.Page) (*types.Empty, error) {
	db := p.DB(ctx)

	err := db.C(ColPages).Insert(in)
	if err != nil {
		return nil, err
	}

	return &types.Empty{}, nil
}
