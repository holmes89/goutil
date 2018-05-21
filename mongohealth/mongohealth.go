package mongohealth

import (
	"github.com/dimiro1/health"
	mgo "gopkg.in/mgo.v2"
)

//MongoChecker checker for mongo db
type MongoChecker struct {
	Session *mgo.Session
}

// NewMongoChecker returns a new db.Checker with the given URL
func NewMongoChecker(session *mgo.Session) MongoChecker {
	return MongoChecker{Session: session}
}

// Check execute session ping and add version
func (c MongoChecker) Check() health.Health {

	health := health.NewHealth()

	err := c.Session.Ping()
	if err != nil {
		health.Down().AddInfo("error", err.Error())
		return health
	}

	build, _ := c.Session.BuildInfo()
	health.Up().AddInfo("version", build.Version)

	return health
}
