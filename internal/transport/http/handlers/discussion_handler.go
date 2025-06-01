package handlers

import (
	"fmt"
	"github.com/Defenestrationq/forum-api/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Data struct {
	Title string `json:"title"`
}

// @Summary Получить все обсуждения
// @Description Возвращает список всех обсуждений на форуме
// @Tags Обсуждения
// @Accept json
// @Produce json
// @Success 200 {array} entity.Discussion "Список обсуждений"
// @Failure 500 {object} object "Ошибка сервера"
// @Router /discussions [get]
func (f *ForumHandler) GetAllDiscussion(c *gin.Context) {
	discussions, err := f.dis.GetAllDiscussions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка получения всех обсуждений: %v", err)})
		return
	}
	c.JSON(http.StatusOK, discussions)
}

// @Summary Получить обсуждение по ID
// @Description Возвращает обсуждение
// @Tags Обсуждения
// @Accept json
// @Produce json
// @Param id path int true "ID обсуждения"
// @Success 200 {array} entity.Discussion "Список обсуждений пользователя"
// @Failure 400 {object} object "Невалидный параметр ID"
// @Failure 500 {object} object "Ошибка сервера"
// @Router /discussions/{id} [get]
func (f *ForumHandler) GetDiscussionByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	discussion, err := f.dis.GetDiscussionsByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка получения обсуждения по id: %v", err)})
		return
	}
	c.JSON(http.StatusOK, discussion)
}

// @Summary Получить обсуждения по ID пользователя
// @Description Возвращает все обсуждения, созданные указанным пользователем
// @Tags Обсуждения
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {array} entity.Discussion "Список обсуждений пользователя"
// @Failure 400 {object} object "Невалидный параметр ID"
// @Failure 500 {object} object "Ошибка сервера"
// @Router /discussions/user/{id} [get]
func (f *ForumHandler) GetDiscussionsByUserID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Невалидный параметр id"})
		return
	}

	userDis, err := f.dis.GetDiscussionsByUserID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userDis)
}

// @Summary Создать новое обсуждение
// @Description Создает новое обсуждение на форуме
// @Tags Обсуждения
// @Accept json
// @Produce json
// @Param discussion body entity.Discussion true "Данные обсуждения"
// @Success 200 {object} object "Обсуждение создано"
// @Failure 400 {object} object "Невалидные данные"
// @Failure 500 {object} object "Ошибка сервера"
// @Router /discussions [post]
func (f *ForumHandler) CreateDiscussions(c *gin.Context) {
	var dis entity.Discussion
	if err := c.ShouldBindJSON(&dis); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Невалидные данные обсуждения"})
		return
	}

	err := f.dis.CreateDiscussions(dis)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "обсуждение создано")
}

// @Summary Обновить обсуждение
// @Description Обновляет существующее обсуждение
// @Tags Обсуждения
// @Accept json
// @Produce json
// @Param discussion body entity.Discussion true "Обновленные данные обсуждения"
// @Success 200 {object} object "Обсуждение изменено"
// @Failure 400 {object} object "Невалидные данные"
// @Failure 500 {object} object "Ошибка сервера"
// @Router /discussions/update/{id} [put]
func (f *ForumHandler) UpdateDiscussions(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var data Data
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Невалидные данные обсуждения"})
		return
	}

	fmt.Println(data.Title)

	dis, err := f.dis.GetDiscussionsByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	dis.Title = data.Title

	err = f.dis.UpdateDiscussions(dis)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "обсуждение изменено")
}

// @Summary Удалить обсуждение
// @Description Удаляет обсуждение по его ID
// @Tags Обсуждения
// @Accept json
// @Produce json
// @Param id path int true "ID обсуждения"
// @Success 200 {object} object "Обсуждение удалено"
// @Failure 400 {object} object "Невалидный параметр ID"
// @Failure 500 {object} object "Ошибка сервера"
// @Router /discussions/{id} [delete]
func (f *ForumHandler) DeleteDiscussions(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Невалидные параметр id"})
		return
	}

	err = f.dis.DeleteDiscussions(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "обсуждение удалено")
}
