package types

import "time"

// Article data
type Article struct {
	ID         string    `json:"id,omitempty" bson:"_id"`
	Author     string    `json:"author,omitempty" bson:"author"`
	Title      string    `json:"title,omitempty" bson:"title"`
	Slug       string    `json:"slug,omitempty" bson:"slug"`
	Content    string    `json:"content,omitempty" bson:"content"`
	Categories []string  `json:"categories" bson:"categories"`
	Tags       []string  `json:"tags" bson:"tags"`
	Update     time.Time `json:"update,omitempty" bson:"update"`
}
