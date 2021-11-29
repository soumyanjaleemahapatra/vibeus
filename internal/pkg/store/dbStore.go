package store

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"github.com/soumyanjaleemahapatra/vibeus/internal/pkg/auth"
	"github.com/soumyanjaleemahapatra/vibeus/internal/pkg/config"
	"github.com/soumyanjaleemahapatra/vibeus/internal/pkg/vibe"
)

type DBStore struct {
	Conn *gorm.DB
}

func Init(conf *config.Configuration) (*DBStore, error) {
	db := DBStore{}
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.DatabaseHost, conf.DatabasePort, conf.DatabaseUser, conf.DatabasePassword, conf.DatabaseName)
	conn, err := gorm.Open("postgres", dns)
	if err != nil {
		return nil, err
	}
	db.Conn = conn
	db.migrate()
	return &db, nil
}
func (d *DBStore) migrate() {
	log.Info("starting database migration")
	d.Conn.AutoMigrate(vibe.Vibe{}, vibe.Comment{}, auth.User{})
	log.Info("database migration complete")
}
func (d *DBStore) CreateVibe(ctx context.Context, scream *vibe.Vibe) error {
	return d.Conn.Create(scream).Error
}
func (d *DBStore) GetVibe(ctx context.Context, id string) (*vibe.Vibe, error) {
	var s vibe.Vibe
	if err := d.Conn.Where("id  = ?", id).First(&s).Error; err != nil {
		return nil, err
	}
	return &s, nil
}
func (d *DBStore) ListVibes(ctx context.Context) ([]*vibe.Vibe, error) {
	var vibes []*vibe.Vibe
	if err := d.Conn.Find(&vibes).Error; err != nil {
		return nil, err
	}
	return vibes, nil
}
func (d *DBStore) DeleteVibe(ctx context.Context, id string) error {
	return d.Conn.Delete(&vibe.Vibe{}, id).Error
}
