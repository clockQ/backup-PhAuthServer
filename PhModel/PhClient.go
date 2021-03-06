package PhModel

import (
	"gopkg.in/mgo.v2/bson"
)

type Client struct {
	ID           string        `json:"-"`
	Id_          bson.ObjectId `json:"-" bson:"_id"`
	ClientID     string        `json:"client-id" bson:"client-id"`
	Secret       string        `json:"secret" bson:"secret"`
	Domain       string        `json:"domain" bson:"domain"`
	Describe     string        `json:"describe" bson:"describe"`
	RegisterDate float64        `json:"register-date" bson:"register-date"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (a Client) GetID() string {
	return a.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (a *Client) SetID(id string) error {
	a.ID = id
	return nil
}

func (a *Client) GetConditionsBsonM(parameters map[string][]string) bson.M {
	return bson.M{}
}
