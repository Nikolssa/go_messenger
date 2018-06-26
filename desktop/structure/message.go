package structure

import (
	"github.com/jinzhu/gorm"
)

//Message is a model to Database table
type Message struct {
	gorm.Model

	User  User
	Group Group

	Content            string `json:"message_content"`
	MessageSenderID    uint   `json:"message_sender_id"`
	MessageRecipientID uint   `json:"message_recepient_id"`
	MessageContentType string `json:"message_content_type_id"`
}