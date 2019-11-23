package config

import (
	"github.com/globalsign/mgo"
	"github.com/mmuflih/envgo/conf"
	"log"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-01-31 09:25
**/

func NewMongoDB(cfg conf.Config) (error, *mgo.Database) {
	address := cfg.GetString(`mongodb.address`)
	database := cfg.GetString(`mongodb.database`)
	log.Println("mongo config", address, database)
	session, err := mgo.Dial(address)
	session.SetMode(mgo.Monotonic, true)
	return err, session.DB(database)
}
