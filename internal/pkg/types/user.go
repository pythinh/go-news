package types

// User models
type User struct {
	ID          string `json:"id,omitempty" bson:"_id"`
	Username    string `json:"username,omitempty" bson:"username"`
	Fullname    string `json:"fullname" bson:"fullname"`
	Email       string `json:"email,omitempty" bson:"email"`
	Password    string `json:"password,omitempty" bson:"password"`
	OldPassword string `json:"old_password" bson:"old_password"`
	Priority    int    `json:"priority,omitempty" bson:"priority"`
}
