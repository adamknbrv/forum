package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Defenestrationq/forum-api/internal/entity"
	"github.com/Defenestrationq/forum-api/pkg/log"
	"go.uber.org/zap"
	"time"
)

type DisRepository interface {
	GetAll() ([]entity.Discussion, error)
	GetByID(id int) (entity.Discussion, error)
	GetByUserID(uid int) ([]entity.Discussion, error)
	Create(dis entity.Discussion) error
	Update(discussion entity.Discussion) error
	Delete(id int) error
}

type DiscussionRep struct {
	db *sql.DB
}

func NewDiscRep(db *sql.DB) DisRepository {
	return &DiscussionRep{db: db}
}

func (m *DiscussionRep) GetAll() ([]entity.Discussion, error) {
	rows, err := m.db.Query(`select * from discussions`)
	if err != nil {
		log.Send.Error("Ошибка получения всех обсуждений",
			zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var discs []entity.Discussion
	for rows.Next() {
		var disc entity.Discussion
		if err := rows.Scan(&disc.ID, &disc.Title, &disc.UserID, &disc.CreateAt); err != nil {
			log.Send.Error("Ошибка сканирования данных обсуждения",
				zap.Error(err))
			return nil, err
		}
		discs = append(discs, disc)
	}

	if err := rows.Err(); err != nil {
		log.Send.Error("Ошибка после итерации по результатам",
			zap.Error(err))
		return nil, err
	}

	return discs, nil
}

func (m *DiscussionRep) GetByID(id int) (entity.Discussion, error) {
	rows := m.db.QueryRow(`select * from discussions where id = $1`, id)
	var disc entity.Discussion

	if err := rows.Scan(&disc.ID, &disc.Title, &disc.UserID, &disc.CreateAt); err != nil {
		return entity.Discussion{}, err
	}
	return disc, nil
}

func (m *DiscussionRep) GetByUserID(uid int) ([]entity.Discussion, error) {
	rows, err := m.db.Query(`select * from discussions where user_id = $1`, uid)
	if err != nil {
		log.Send.Error("Ошибка получения всех обсуждений пользователя",
			zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var discs []entity.Discussion
	for rows.Next() {
		var disc entity.Discussion
		if err = rows.Scan(&disc.ID, &disc.UserID, &disc.Title, &disc.CreateAt); err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				log.Send.Error("Ошибка получения данных из таблицы",
					zap.Error(err))
				return nil, nil
			}
		}
		discs = append(discs, disc)
	}

	return discs, nil
}

func (m *DiscussionRep) Create(dis entity.Discussion) error {
	query := `insert into discussions (title, user_id, create_at) 
			  values ($1, $2, $3)
			  returning id`

	var id int
	err := m.db.QueryRow(
		query,
		dis.Title,
		dis.UserID,
		time.Now(),
	).Scan(&id)

	if err != nil {
		log.Send.Error("Ошибка создания обсуждения",
			zap.Error(err))
		return err
	}
	return nil

}

func (m *DiscussionRep) Update(dis entity.Discussion) error {
	query := `update discussions
			  set title = $2
			  where id = $1`
	exec, err := m.db.Exec(query, dis.ID, dis.Title)
	if err != nil {
		log.Send.Error("Ошибка запроса на редактирования обсуждения",
			zap.Error(err))
		return err
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		log.Send.Error("Ошибка получения измененных строк обсуждения",
			zap.Error(err))
		return err
	}

	if affected == 0 {
		log.Send.Error("Обсуждение не изменено",
			zap.Error(err))
		return fmt.Errorf("Обсуждение не изменено")
	}
	return nil
}

func (m *DiscussionRep) Delete(id int) error {
	exec, err := m.db.Exec(`delete from discussions where id = $1`, id)
	if err != nil {
		log.Send.Error("Ошибка запроса на удаление обсуждения",
			zap.Error(err))
		return err
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		log.Send.Error("Ошибка получения измененных строк обсуждения",
			zap.Error(err))
		return err
	}

	if affected == 0 {
		log.Send.Error("Обсуждение не удалено",
			zap.Error(err))
		return fmt.Errorf("Обсуждение не удалено")
	}
	return nil
}
