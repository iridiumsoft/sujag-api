package main

import (
	"github.com/sirupsen/logrus"
	"time"
	"github.com/user/sujag/app"
	"github.com/user/sujag/controllers"
	"github.com/user/sujag/conf"
	"gopkg.in/mgo.v2"
)

func main() {

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	config := conf.FromFile("conf/config.json")

	db, err := initDB(config)

	for err != nil {
		logrus.WithError(err).Error("Failed init connection to database. Trying again after 10 seconds")
		time.Sleep(time.Second * 10)
		db, err = initDB(config)
	}

	application := app.New(db)

	c := controllers.New(config, application)

	logrus.Fatal(c.Routing())
}

func initDB(conf conf.Main) (*mgo.Database, error) {
	session, err := mgo.Dial("mongodb://localhost/" + conf.DB.Name)
	db := session.DB(conf.DB.Name)
	return db, err
}
