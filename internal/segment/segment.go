package segment

import (
	"github.com/SicParv1sMagna/NetworkingDataLinkLayer/internal/utils"
	log "github.com/sirupsen/logrus"
)

const (
	// mask для получения старших 4 битов
	upper4BitsMask = 0xF0
	// mask для получения младших 4 битов
	lower4BitsMask = 0x0F
)

// SplitSegmentToCycleCodes разбивает сегмент на цикл-коды
func (s *Segment) SplitSegmentToCycleCodes() []utils.CycleCode {
	var cycleCodes []utils.CycleCode

	for _, byteValue := range s.Payload {
		// Получаем старшие 4 бита
		upperCode := uint(byteValue & upper4BitsMask >> 4)
		// Получаем младшие 4 бита
		lowerCode := uint(byteValue & lower4BitsMask)
		cycleCodes = append(cycleCodes, utils.CycleCode{Code: upperCode})
		cycleCodes = append(cycleCodes, utils.CycleCode{Code: lowerCode})
	}

	log.Info("Сегмент успешно разделен на цикл-коды: ", cycleCodes)
	return cycleCodes
}

// Simulate моделирует кодирование, ошибки и декодирование для каждого цикл-кода в последовательности
func (s *Segment) Simulate(cycleCodes []utils.CycleCode) []utils.CycleCode {
	for _, code := range cycleCodes {
		code.Encode()
		code.ErrorSimulate()
		code.Decode()
	}

	log.Info("Симуляция (кодирование, ошибка и декодирование) прошла успешно, полученная последовательность: ", cycleCodes)
	return cycleCodes
}

// SplitCycleCodesToSegment собирает последовательность цикл-кодов обратно в сегмент
func (s *Segment) SplitCycleCodesToSegment(cycleCodes []utils.CycleCode) {
	if len(cycleCodes)%2 != 0 {
		log.Warn("Нечетное количество цикл-кодов. Последний цикл-код будет проигнорирован.")
	}

	var payload []byte

	for i := 0; i < len(cycleCodes)-1; i += 2 {
		byteValue := (byte(cycleCodes[i].Code) << 4) | (byte(cycleCodes[i+1].Code) & lower4BitsMask)
		payload = append(payload, byteValue)
	}
	s.Payload = payload

	log.Info("Сегмент успешно собран: ", s.Payload)
}
