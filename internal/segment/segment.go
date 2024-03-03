package segment

import (
	"github.com/SicParv1sMagna/NetworkingDataLinkLayer/internal/utils"
)

func (s *Segment) SplitSegmentToCycleCodes() []utils.CycleCode {
	var cycleCodes []utils.CycleCode

	for _, byteValue := range s.Payload {
		for i := uint(0); i < 8; i += 4 {
			code := uint(byteValue >> i & 0xF)
			cycleCodes = append(cycleCodes, utils.CycleCode{Code: code})
		}
	}

	return cycleCodes
}

func (s *Segment) Simulate(cycleCodes []utils.CycleCode) []utils.CycleCode {
	for _, code := range cycleCodes {
		code.EnCode()
		code.ErrorSimulate()
		code.DeCode()
	}

	return cycleCodes
}

func (s *Segment) SplitCycleCodesToSegment(cycleCodes []utils.CycleCode) {
	for _, code := range cycleCodes {
		s.Payload = append(s.Payload, byte(code.Code))
	}
}
