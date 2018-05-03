package main

import (
	"github.com/dracher/autorhvhprovison/provision"
	"github.com/dracher/autorhvhprovison/route"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

const (
	mongoURL = ""
	mongoDB  = ""
	ip       = ""
	port     = ""
)

func main() {
	// pre.Check()

	session, err := mgo.Dial(mongoURL)
	if err != nil {
		log.Panic(err)
	}
	defer session.Close()
	db := session.DB(mongoDB)
	autoCfg := provision.NewAutoConfig(db, ip, port)
	autoCfg.UpdateProfiles()
	autoCfg.UpdateBuilds()

	r := gin.Default()
	r.Use(route.AutoConfigMiddle(autoCfg))
	r.Static("/ks", "./static")

	api := r.Group("/api/v1")
	route.RegisteRoute(api)

	r.Run(":" + port)
}
