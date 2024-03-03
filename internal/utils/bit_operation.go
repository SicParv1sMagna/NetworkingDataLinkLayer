package utils

import "math/bits"

func remainderFinding(code uint) uint {
	polinom := uint(0b1011)

	num := code >> 3

	var infBits [3]uint
	for i := 0; i < 3; i++ {
		infBits[2-i] = (code >> uint(i)) & 1
	}

	var remains uint
	var end bool

	for countZeros := 0; !end; {
		result := num ^ polinom
		result &= (1 << bits.Len(result)) - 1
		remains = result

		for ; 4-bits.Len(result) > 0; countZeros++ {
			if countZeros == 3 {
				end = true
				break
			}
			result = (result << 1) | infBits[countZeros] // Добавляем к результату
		}

		if bits.Len(result) != 4 {
			remains = result
		}

		num = result
	}

	return remains
}

func cyclicShiftLeft(codeArray []uint) []uint {
	return append(codeArray[1:], codeArray[0])
}

func cyclicShiftRight(codeArray []uint) []uint {

	if len(codeArray) == 0 {
		return codeArray
	}

	lastBit := codeArray[len(codeArray)-1]

	return append([]uint{lastBit}, codeArray[:len(codeArray)-1]...)
}

func uintToBitsArray(num uint) []uint {
	var bitsArray []uint

	for i := uint(0); i < 7; i++ {
		bit := (num >> (7 - 1 - i)) & 1
		bitsArray = append(bitsArray, bit)
	}

	return bitsArray
}

func bitsArrayToUint(bitsArray []uint) uint {
	var result uint

	for _, bit := range bitsArray {
		result = (result << 1) | bit
	}

	return result
}
