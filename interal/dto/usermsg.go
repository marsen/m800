package dto

type Message struct {
	UserID string `json:"userid" bson:"userid"`
	Text   string `json:"text" bson:"text"`
}
