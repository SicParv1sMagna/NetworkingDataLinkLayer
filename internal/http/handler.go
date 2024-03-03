package http

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/SicParv1sMagna/NetworkingDataLinkLayer/internal/segment"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	BaseURL string
}

func NewHandler(baseURL string) *Handler {
	return &Handler{BaseURL: baseURL}
}

func (h *Handler) EncodeSegmentSimulate(c *gin.Context) {
	var segment segment.Segment
	if err := c.BindJSON(&segment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "не удалось прочитать JSON"})
		return
	}

	segment.SplitCycleCodesToSegment(segment.Simulate(segment.SplitSegmentToCycleCodes()))

	randomNumber := rand.Intn(100)

	if randomNumber < 2 {
		c.JSON(http.StatusOK, gin.H{"message": "сегмент утерян"})
		return
	}

	segmentJSON, err := json.Marshal(segment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка при кодировании сегмента в JSON"})
		return
	}

	resp, err := http.Post(h.BaseURL, "application/json", bytes.NewBuffer(segmentJSON))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ошибка при отправке сегмента на эндпоинт:"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ошибка: неверный код состояния ответа:"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "сегмент успешно отправлен на эндпоинт"})
}
