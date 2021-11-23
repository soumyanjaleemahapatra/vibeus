package store

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"github.com/soumyanjaleemahapatra/vibeus/internal/pkg/config"
	"github.com/soumyanjaleemahapatra/vibeus/internal/pkg/scream"
)

type DBStore struct {
	Conn *gorm.DB
}

func Init(conf *config.Configuration) (*DBStore, error) {
	db := DBStore{}
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.DatabaseHost, conf.DatabasePort, conf.DatabaseUser, conf.DatabasePassword, conf.DatabaseName)
	conn, err := gorm.Open("postgres", dns)
	log.Infof("database DNS: %s", dns)
	if err != nil {
		return nil, err
	}
	db.Conn = conn
	return &db, nil
}
func (d *DBStore) CreateScream(ctx context.Context, scream *scream.Scream) error {
	return nil
}
func (d *DBStore) GetScream(ctx context.Context, id string) (*scream.Scream, error) {
	return nil, nil
}
func (d *DBStore) ListScreams(ctx context.Context) ([]*scream.Scream, error) {
	return nil, nil
}
func (d *DBStore) DeleteScream(ctx context.Context, id string) error {
	return nil
}
