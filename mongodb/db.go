package mongodb

import "github.com/globalsign/mgo"

func NewClient(host string) (*mgo.Session, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Strong, true)
	return session, nil
}
