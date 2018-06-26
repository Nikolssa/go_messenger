package dbservice

import "go_messenger/server/models"

//Message struct
type MessageDBService struct {
	models.Message
}

//AddMessage func.
func (msg MessageDBService) AddMessage(message *models.Message) bool {
	dbConn.Where("username = ?", message.User.Username).First(&message.User)
	dbConn.Where("group_name = ?", message.Group.GroupName).First(&message.Group)
	if dbConn.NewRecord(message) {
		dbConn.Create(&message)
		return true
	}
	return false
}

//GetGroupMessages gets messages of special group with count limit.
func (msg MessageDBService) GetGroupMessages(group *models.Group, count uint) []models.Message {
	var messageList = []models.Message{}
	dbConn.Where("group_name = ?", group.GroupName).First(&group)
	dbConn.Where("message_recipient_id = ?", group.ID).Limit(count).Find(&messageList)
	return messageList
}