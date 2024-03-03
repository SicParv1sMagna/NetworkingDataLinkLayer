package utils

import (
	"math/bits"
	"math/rand"
	"time"
)

type CycleCode struct {
	Code uint
}

func (c *CycleCode) EnCode() {
	c.Code <<= 3
	c.Code = c.Code ^ remainderFinding(c.Code)
}

func (c *CycleCode) ErrorSimulate() {
	var globalRand = rand.New(rand.NewSource(time.Now().UnixNano()))

	randomValue := globalRand.Intn(10)

	if randomValue <= 0 {
		errorPosition := uint(rand.Intn(bits.Len(c.Code)))
		c.Code ^= (1 << errorPosition)
	}
}

func (c *CycleCode) DeCode() {
	var count uint
	CodeArray := uintToBitsArray(c.Code)

	var end bool

	for remains := remainderFinding(bitsArrayToUint(CodeArray)); !end; remains = remainderFinding(bitsArrayToUint(CodeArray)) {
		if bits.Len(remains) <= 1 {
			end = true
			CodeArray[len(CodeArray)-1] = CodeArray[len(CodeArray)-1] ^ 1
			break
		}

		CodeArray = cyclicShiftLeft(CodeArray)
		count++
	}

	for ; count > 0; count-- {
		CodeArray = cyclicShiftRight(CodeArray)
	}

	c.Code = bitsArrayToUint(CodeArray[:len(CodeArray)-3])
}
