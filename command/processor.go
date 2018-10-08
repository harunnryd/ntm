package command

import (
	"ntm/db"
	"ntm/model"
	"ntm/repo"
)

func (cmd *CreateStatusCommand) Process() error {
	newStatus := model.NewStatus()
	newStatus.ID = cmd.ID
	newStatus.Name = cmd.Name

	return repo.NewStatusRepo().Create(newStatus)
}

func (cmd *UpdateStatusCommand) Process() error {
	updateStatus := model.NewTopic()
	updateStatus.Name = cmd.Name

	return repo.NewStatusRepo().Update(cmd.StatusID, updateStatus)
}

func (cmd *DeleteStatusCommand) Process() error {
	return repo.NewStatusRepo().Delete(cmd.StatusID)

}

func (cmd *CreateTopicCommand) Process() error {
	newTopic := model.NewTopic()
	newTopic.ID = cmd.ID
	newTopic.Name = cmd.Name

	return repo.NewTopicRepo().Create(newTopic)
}

func (cmd *UpdateTopicCommand) Process() error {
	updateTopic := model.NewTopic()
	updateTopic.Name = cmd.Name

	return repo.NewTopicRepo().Update(cmd.TopicID, updateTopic)
}

func (cmd *DeleteTopicCommand) Process() error {
	return repo.NewTopicRepo().Delete(cmd.TopicID)

}

func (cmd *CreateNewsCommand) Process() error {
	newNews := model.NewNews()
	newNews.ID = cmd.ID
	newNews.Title = cmd.Title
	newNews.Body = cmd.Body

	return repo.NewNewsRepo().Create(newNews)
}

func (cmd *UpdateNewsCommand) Process() error {
	updateNews := model.NewNews()
	updateNews.Title = cmd.Title
	updateNews.Body = cmd.Body

	return repo.NewNewsRepo().Update(cmd.NewsID, updateNews)
}

func (cmd *DeleteNewsCommand) Process() error {
	return repo.NewNewsRepo().Delete(cmd.NewsID)
}

func (cmd *AssignTopicCommand) Process() error {
	return db.NewORM().Connect().Exec(`insert into news_topics (news_id, topic_id) values (?, ?)`, cmd.NewsID, cmd.TopicID).Error
}

func (cmd *UnAssignTopicCommand) Process() error {
	return db.NewORM().Connect().Exec(`delete from news_topics where news_id = ? and topic_id = ?`, cmd.NewsID, cmd.TopicID).Error
}

func (cmd *ChangeNewsStatusCommand) Process() error {
	updateNews := new(model.News)
	updateNews.StatusID = cmd.StatusID
	return repo.NewNewsRepo().Update(cmd.NewsID, updateNews)
}