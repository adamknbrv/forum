package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/adamknbrv/forum/internal/entity"
	"github.com/adamknbrv/forum/pkg/log"
	"go.uber.org/zap"
	"time"
)

type MsgRepository interface {
	GetAll(disID int) ([]entity.Message, error)
	GetByUserID(uid int) ([]entity.Message, error)
	Create(message entity.Message) error
	Update(message entity.Message) error
	Delete(id int) error
}

type MessageRep struct {
	db *sql.DB
}

func NewMesRep(db *sql.DB) MsgRepository {
	return &MessageRep{db: db}
}

func (m *MessageRep) GetAll(disID int) ([]entity.Message, error) {
	rows, err := m.db.Query(`SELECT id, user_id, content, discussion_id, create_at FROM messages WHERE discussion_id = $1`, disID)
	if err != nil {
		log.Send.Error("Ошибка получения всех сообщений обсуждения",
			zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var messages []entity.Message
	for rows.Next() {
		var message entity.Message
		err = rows.Scan(
			&message.ID,
			&message.UserID,
			&message.Content,
			&message.DiscussionID,
			&message.CreateAt,
		)
		if err != nil {
			log.Send.Error("Ошибка получения данных из таблицы",
				zap.Error(err))
			return nil, err
		}
		messages = append(messages, message)
	}

	if err = rows.Err(); err != nil {
		log.Send.Error("Ошибка итерации по результатам", zap.Error(err))
		return nil, err
	}

	return messages, nil
}

func (m *MessageRep) GetByUserID(uid int) ([]entity.Message, error) {
	rows, err := m.db.Query(`SELECT * FROM messages WHERE user_id = $1`, uid)
	if err != nil {
		log.Send.Error("Ошибка получения всех сообщений пользователя",
			zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var messages []entity.Message
	for rows.Next() {
		var message entity.Message
		if err = rows.Scan(
			&message.ID,
			&message.UserID,
			&message.DiscussionID,
			&message.Content,
			&message.CreateAt,
		); err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				log.Send.Error("Ошибка получения данных из таблицы",
					zap.Error(err))
				return nil, nil
			}
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func (m *MessageRep) Create(message entity.Message) error {
	query := `INSERT INTO messages (user_id, discussion_id, content, create_at) 
			  VALUES ($1, $2, $3, $4)
			  RETURNING id`

	var id int
	err := m.db.QueryRow(
		query,
		message.UserID,
		message.DiscussionID,
		message.Content,
		time.Now(),
	).Scan(&id)

	if err != nil {
		log.Send.Error("Ошибка создания сообщения",
			zap.Error(err))
		return err
	}
	return nil
}

func (m *MessageRep) Update(message entity.Message) error {
	query := `UPDATE messages
			  SET content = $2, discussion_id = $3
			  WHERE id = $1`
	exec, err := m.db.Exec(query, message.ID, message.Content, message.DiscussionID)
	if err != nil {
		log.Send.Error("Ошибка запроса на редактирования сообщения",
			zap.Error(err))
		return err
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		log.Send.Error("Ошибка получения измененных строк сообщения",
			zap.Error(err))
		return err
	}

	if affected == 0 {
		log.Send.Error("Сообщение не изменено",
			zap.Error(err))
		return fmt.Errorf("Сообщение не изменено")
	}
	return nil
}

func (m *MessageRep) Delete(id int) error {
	exec, err := m.db.Exec(`DELETE FROM messages WHERE id = $1`, id)
	if err != nil {
		log.Send.Error("Ошибка запроса на удаление сообщения",
			zap.Error(err))
		return err
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		log.Send.Error("Ошибка получения измененных строк сообщения",
			zap.Error(err))
		return err
	}

	if affected == 0 {
		log.Send.Error("Сообщение не удалено",
			zap.Error(err))
		return fmt.Errorf("Сообщение не удалено")
	}
	return nil
}
