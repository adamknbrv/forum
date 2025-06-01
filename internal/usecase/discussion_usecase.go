package usecase

import (
	"github.com/adamknbrv/forum/internal/entity"
	"github.com/adamknbrv/forum/internal/repository"
)

type DisUseCase interface {
	GetAllDiscussions() ([]entity.Discussion, error)
	GetDiscussionsByID(id int) (entity.Discussion, error)
	GetDiscussionsByUserID(uid int) ([]entity.Discussion, error)
	CreateDiscussions(dis entity.Discussion) error
	UpdateDiscussions(discussion entity.Discussion) error
	DeleteDiscussions(id int) error
}

type DiscussionUseCase struct {
	repo repository.DisRepository
}

func NewDisUseCase(repo repository.DisRepository) DisUseCase {
	return &DiscussionUseCase{repo: repo}
}

func (d *DiscussionUseCase) GetAllDiscussions() ([]entity.Discussion, error) {
	return d.repo.GetAll()
}

func (d *DiscussionUseCase) GetDiscussionsByUserID(uid int) ([]entity.Discussion, error) {
	return d.repo.GetByUserID(uid)
}

func (d *DiscussionUseCase) GetDiscussionsByID(id int) (entity.Discussion, error) {
	return d.repo.GetByID(id)
}

func (d *DiscussionUseCase) CreateDiscussions(dis entity.Discussion) error {
	return d.repo.Create(dis)
}

func (d *DiscussionUseCase) UpdateDiscussions(discussion entity.Discussion) error {
	return d.repo.Update(discussion)
}

func (d *DiscussionUseCase) DeleteDiscussions(id int) error {
	return d.repo.Delete(id)
}
