package usecase

import (
	"github.com/Defenestrationq/forum-api/internal/entity"
	"github.com/Defenestrationq/forum-api/internal/repository"
)

type MsgUseCase interface {
	GetAllMessage(disID int) ([]entity.Message, error)
	GetMessageByUserID(uid int) ([]entity.Message, error)
	CreateMessage(message entity.Message) error
	UpdateMessage(message entity.Message) error
	DeleteMessage(id int) error
}

type MessageUseCase struct {
	repo repository.MsgRepository
}

func NewMessageUseCase(repo repository.MsgRepository) MsgUseCase {
	return &MessageUseCase{repo: repo}
}

func (d *MessageUseCase) GetAllMessage(disID int) ([]entity.Message, error) {
	return d.repo.GetAll(disID)
}

func (d *MessageUseCase) GetMessageByUserID(uid int) ([]entity.Message, error) {
	return d.repo.GetByUserID(uid)
}

func (d *MessageUseCase) CreateMessage(message entity.Message) error {
	return d.repo.Create(message)
}

func (d *MessageUseCase) UpdateMessage(message entity.Message) error {
	return d.repo.Update(message)
}

func (d *MessageUseCase) DeleteMessage(id int) error {
	return d.repo.Delete(id)
}
