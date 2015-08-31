package models

type Category struct {
	Id    bson.ObjectId `json:"id" bson:"_id"`
	Title string        `json:"title" bson:"title"`
	Slug  string        `json:"slug" bson:"slug"`
}
