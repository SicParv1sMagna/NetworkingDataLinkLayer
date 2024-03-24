package segment

import (
	"github.com/SicParv1sMagna/NetworkingDataLinkLayer/internal/utils"
	log "github.com/sirupsen/logrus"
)

func (s *Segment) SplitSegmentToCycleCodes() []utils.CycleCode {
	var cycleCodes []utils.CycleCode

	for _, byteValue := range s.Payload {
		for i := uint(0); i < 8; i += 4 {
			code := uint(byteValue >> i & 0xF)
			cycleCodes = append(cycleCodes, utils.CycleCode{Code: code})
		}
	}
	
	log.Info("полученный сегмент разделен на биты: ", cycleCodes)
	return cycleCodes
}

func (s *Segment) Simulate(cycleCodes []utils.CycleCode) []utils.CycleCode {
	for _, code := range cycleCodes {
		code.Encode()
		code.ErrorSimulate()
		code.Decode()
	}

	log.Info("симуляция (кодирование, ошибка и декодирование) прошла успешно, полученная последовательность: ", cycleCodes)
	return cycleCodes
}

func (s *Segment) SplitCycleCodesToSegment(cycleCodes []utils.CycleCode) {
	for _, code := range cycleCodes {
		s.Payload = append(s.Payload, byte(code.Code))
	}
	log.Info("сегмент успешно собран: ", cycleCodes)
}
