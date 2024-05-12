package http

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"

	"github.com/SicParv1sMagna/NetworkingDataLinkLayer/internal/segment"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	BaseURL string
	log     *logrus.Logger
}

func NewHandler(baseURL string, log *logrus.Logger) *Handler {
	return &Handler{BaseURL: baseURL,
		log: log}
}

// @Summary EncodeSegmentSimulate.
// @Description Кодирует и декодирует полученный в виде байт сегмент, вносит ошибку, исправляет ее, так же с вероятностью возвращает сегмент на траспортный уровень.
// @Tags Code
// @Accept json
// @Produce json
// @Param segment body segment.SegmentRequest true "Пользовательский объект в формате JSON"
// @Success 200 "Успешный ответ"
// @Failure 400 {object} swag.ErrorResponse "Ошибка в запросе"
// @Failure 500 {object} swag.ErrorResponse "Внутренняя ошибка сервера"
// @Router /channel/code [post]
func (h *Handler) EncodeSegmentSimulate(c *gin.Context) {
	const op = "handlers.EncodeSegmentSimulate"
	log := h.log.WithField("operation", op)

	var segReq segment.SegmentRequest
	if err := c.BindJSON(&segReq); err != nil {
		log.WithError(err).Error("ошибка чтения входящего JSON сегмента")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось прочитать JSON: " + err.Error()})
		return
	}

	log.WithField("segmentRequest", segReq).Info("на вход поступил сегмент")

	seg := segment.Segment{ID: segReq.ID,
		TotalSegments: segReq.TotalSegments,
		SenderName:    segReq.SenderName,
		SegmentNumber: segReq.SegmentNumber,
		Payload:       segReq.Payload,
	}

	cycleCode := seg.Simulate(seg.SplitSegmentToCycleCodes(h.log), h.log)
	seg.Payload = nil
	seg.SplitCycleCodesToSegment(cycleCode, h.log)

	randomNumber := rand.Intn(100)

	if randomNumber < 2 {
		log.WithField("segment", seg).Warn("потеря сегмента с вероятностью 2%")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "сегмент утерян"})
		return
	}

	segmentJSON, err := json.Marshal(seg)
	if err != nil {
		log.WithError(err).Error("ошибка при кодировании сегмента в JSON")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка при кодировании сегмента в JSON: " + err.Error()})
		return
	}

	response, err := http.Post(h.BaseURL, "application/json", bytes.NewBuffer(segmentJSON))
	if err != nil {
		log.WithError(err).Error("ошибка при отправке сегмента на эндпоинт")
		c.JSON(http.StatusBadRequest, gin.H{"error": "ошибка при отправке сегмента на эндпоинт: " + err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Error("неверный код состояния ответа")
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный код состояния ответа"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "сегмент успешно отправлен на эндпоинт"})
}
