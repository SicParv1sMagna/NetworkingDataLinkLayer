package http

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/SicParv1sMagna/NetworkingDataLinkLayer/internal/segment"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	BaseURL string
}

func NewHandler(baseURL string) *Handler {
	return &Handler{BaseURL: baseURL}
}

// @Summary EncodeSegmentSimulate.
// @Description Кодирует и декодирует полученный в виде байт сегмент, вносит ошибку, исправляет ее, так же с вероятностью возвращает сегмент на траспортный уровень.
// @Tags Code
// @Accept json
// @Produce json
// @Param segment body segment.Segment true "Пользовательский объект в формате JSON"
// @Success 200 "Успешный ответ"
// @Failure 400 {object} swag.ErrorResponse "Ошибка в запросе"
// @Failure 500 {object} swag.ErrorResponse "Внутренняя ошибка сервера"
// @Router /channel/code [post]
func (h *Handler) EncodeSegmentSimulate(c *gin.Context) {
	var segment segment.Segment
	if err := c.BindJSON(&segment); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "не удалось прочитать JSON: "+ err.Error()})
		return
	}

	segment.SplitCycleCodesToSegment(segment.Simulate(segment.SplitSegmentToCycleCodes()))

	randomNumber := rand.Intn(100)

	if randomNumber < 2 {
		log.Info("потеря сегмента")
		c.JSON(http.StatusOK, gin.H{"message": "сегмент утерян"})
		return
	}

	segmentJSON, err := json.Marshal(segment)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка при кодировании сегмента в JSON: "+ err.Error()})
		return
	}

	resp, err := http.Post(h.BaseURL, "application/json", bytes.NewBuffer(segmentJSON))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ошибка при отправке сегмента на эндпоинт:: "+ err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("ошибка: неверный код состояния ответа")
		c.JSON(http.StatusBadRequest, gin.H{"error": "ошибка: неверный код состояния ответа"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "сегмент успешно отправлен на эндпоинт"})
}
