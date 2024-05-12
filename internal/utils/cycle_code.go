package utils

import (
	"math/bits"
	"math/rand"
)

type CycleCode struct {
	Code uint
}

func (c *CycleCode) Encode() {
	c.Code <<= 3
	c.Code = c.Code ^ remainderFinding(c.Code)
}

func (c *CycleCode) ErrorSimulate() bool {
	randomValue := rand.Intn(10)

	if randomValue <= 0 {
		bitsLength := bits.Len(c.Code)
		if bitsLength > 0 {
			errorPosition := uint(rand.Intn(bitsLength))
			c.Code ^= (1 << errorPosition)
			return true
		}
	}
	return false
}

func (c *CycleCode) Decode() {
	var count uint
	сodeArray := uintToBitsArray(c.Code)

	var end bool

	for remains := remainderFinding(bitsArrayToUint(сodeArray)); !end; remains = remainderFinding(bitsArrayToUint(сodeArray)) {
		if bits.Len(remains) <= 1 {
			end = true
			сodeArray[len(сodeArray)-1] = сodeArray[len(сodeArray)-1] ^ 1
			break
		}

		сodeArray = cyclicShiftLeft(сodeArray)
		count++
	}

	for ; count > 0; count-- {
		сodeArray = cyclicShiftRight(сodeArray)
	}

	c.Code = bitsArrayToUint(сodeArray[:len(сodeArray)-3])
}
