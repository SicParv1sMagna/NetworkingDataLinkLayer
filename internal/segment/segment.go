package segment

import (
	"github.com/SicParv1sMagna/NetworkingDataLinkLayer/internal/utils"
	"github.com/sirupsen/logrus"
)

const (
	// mask для получения старших 4 битов
	upper4BitsMask = 0xF0
	// mask для получения младших 4 битов
	lower4BitsMask = 0x0F
)

// SplitSegmentToCycleCodes разбивает сегмент на цикл-коды
func (s *Segment) SplitSegmentToCycleCodes(log *logrus.Logger) []utils.CycleCode {
	const op = "segment.Segment.SplitSegmentToCycleCodes"
	log.WithField("operation", op)

	var cycleCodes []utils.CycleCode

	log.Info("разделение сегмента на цикл-коды")

	for _, byteValue := range s.Payload {
		// Получаем старшие 4 бита
		upperCode := uint(byteValue & upper4BitsMask >> 4)
		// Получаем младшие 4 бита
		lowerCode := uint(byteValue & lower4BitsMask)
		cycleCodes = append(cycleCodes, utils.CycleCode{Code: upperCode})
		cycleCodes = append(cycleCodes, utils.CycleCode{Code: lowerCode})
	}

	return cycleCodes
}

// Simulate моделирует кодирование, ошибки и декодирование для каждого цикл-кода в последовательности
func (s *Segment) Simulate(cycleCodes []utils.CycleCode, log *logrus.Logger) []utils.CycleCode {
	const op = "segment.Segment.Simulate"
	log.WithField("operation", op)

	log.Info("cимуляция (кодирование, ошибка и декодирование)")

	for _, code := range cycleCodes {
		code.Encode()
		if code.ErrorSimulate() {
			s.HadError = true
			log.WithField("code", code).Info("в части сегмента произошла ошибка с вероятностью 10%")
		}
		code.Decode()
	}

	return cycleCodes
}

// SplitCycleCodesToSegment собирает последовательность цикл-кодов обратно в сегмент
func (s *Segment) SplitCycleCodesToSegment(cycleCodes []utils.CycleCode, log *logrus.Logger) {
	const op = "segment.Segment.Simulate"
	log.WithField("operation", op)

	if len(cycleCodes)%2 != 0 {
		log.Warn("нечетное количество цикл-кодов. Последний цикл-код будет проигнорирован.")
	}

	var payload []byte

	for i := 0; i < len(cycleCodes)-1; i += 2 {
		byteValue := (byte(cycleCodes[i].Code) << 4) | (byte(cycleCodes[i+1].Code) & lower4BitsMask)
		payload = append(payload, byteValue)
	}
	s.Payload = payload

	log.WithField("segment", s).Info("сегмент собран")
}
