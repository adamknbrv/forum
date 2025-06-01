package handlers

import (
	"fmt"
	"github.com/adamknbrv/forum/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Получить все сообщения по ID обсуждения
// @Description Возвращает все сообщения для указанного обсуждения
// @Tags Сообщения
// @Accept json
// @Produce json
// @Param id path int true "ID обсуждения"
// @Success 200 {array} entity.Message "Успешный запрос"
// @Failure 400 {object} object "Невалидный параметр ID"
// @Failure 500 {object} object "Ошибка сервера"
// @Router /messages/{id} [get]
func (f *ForumHandler) GetAllMessage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Невалидный параметр в id"})
		return
	}

	messages, err := f.msg.GetAllMessage(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения всех сообщений по id обсуждения"})
		return
	}
	c.JSON(http.StatusOK, messages)
}

// @Summary Получить сообщения по ID пользователя
// @Description Возвращает все сообщения, созданные указанным пользователем
// @Tags Сообщения
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {array} entity.Message "Успешный запрос"
// @Failure 400 {object} object "Невалидный параметр ID"
// @Failure 500 {object} object "Ошибка сервера"
// @Router /messages/user/{id} [get]
func (f *ForumHandler) GetMessageByUserID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Невалидный параметр в id"})
		return
	}

	userMsg, err := f.msg.GetMessageByUserID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка получения всех сообщений пользователя: %v", err)})
		return
	}
	c.JSON(http.StatusOK, userMsg)
}

// @Summary Создать новое сообщение
// @Description Создает новое сообщение в обсуждении
// @Tags Сообщения
// @Accept json
// @Produce json
// @Param message body entity.Message true "Данные сообщения"
// @Success 200 {object} object "Сообщение создано"
// @Failure 400 {object} object "Невалидные параметры"
// @Failure 500 {object} object "Ошибка сервера"
// @Router /messages [post]
func (f *ForumHandler) CreateMessage(c *gin.Context) {
	var msg entity.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Невалидные параметры сообщения"})
		return
	}

	err := f.msg.CreateMessage(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка создания сообщения: %v", err)})
		return
	}
	c.JSON(http.StatusOK, "сообщение создано")
}

// @Summary Обновить сообщение
// @Description Обновляет существующее сообщение
// @Tags Сообщения
// @Accept json
// @Produce json
// @Param message body entity.Message true "Обновленные данные сообщения"
// @Success 200 {object} object "Сообщение изменено"
// @Failure 400 {object} object "Невалидные параметры"
// @Failure 500 {object} object "Ошибка сервера"
// @Router /messages [put]
func (f *ForumHandler) UpdateMessage(c *gin.Context) {
	var msg entity.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Невалидные параметры сообщения"})
		return
	}

	err := f.msg.UpdateMessage(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка редактирования сообщения: %v", err)})
		return
	}
	c.JSON(http.StatusOK, "сообщение изменено")
}

// @Summary Удалить сообщение
// @Description Удаляет сообщение по его ID
// @Tags Сообщения
// @Accept json
// @Produce json
// @Param id path int true "ID сообщения"
// @Success 200 {object} object "Сообщение удалено"
// @Failure 400 {object} object "Невалидный параметр ID"
// @Failure 500 {object} object "Ошибка сервера"
// @Router /messages/{id} [delete]
func (f *ForumHandler) DeleteMessage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Невалидный параметр в id"})
		return
	}

	err = f.msg.DeleteMessage(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка удаления сообщения: %v", err)})
		return
	}
	c.JSON(http.StatusOK, "сообщение удалено")
}
