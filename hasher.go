package hasher

import (
	"math"
	"unsafe"
)

type Hasher interface {
	HashCode() int32
}
type Hashable interface {
	Hasher
}

func HashCode(input any) int32 {
	switch value := input.(type) {

	case Hasher:
		return value.HashCode()

	case bool:
		return HashCodeBool(value)

	case int:
		if unsafe.Sizeof(value) <= 32 {
			return int32(value)
		}
		return HashCodeLong(uint64(value))

	case int8:
		return int32(value)

	case int16:
		return int32(value)

	case int32:
		return int32(value)

	case int64:
		return HashCodeLong(uint64(value))

	case uint:
		if unsafe.Sizeof(value) <= 16 {
			return int32(value)
		}
		return HashCodeLong(uint64(value))

	case uint8:
		return int32(value)

	case uint16:
		return int32(value)

	case uint32:
		return HashCodeLong(uint64(value))

	case uint64:
		return HashCodeLong(uint64(value))

	case float32:
		return HashCodeFloat(value)

	case float64:
		return HashCodeDouble(value)

	case string:
		return HashCodeString(value)

	default:
		panic("hash function of this type not implemented")
	}
}

func HashCodeBool(input bool) int32 {
	if input {
		return 1
	}
	return 0
}

func HashCodeString(input string) int32 {
	result := int32(0)
	if result == 0 && len(input) > 0 {
		for i := range len(input) {
			result = 31 * result + int32(input[i])
		}
	}
	return result
}

func HashCodeLong(input uint64) int32 {
	return int32(input ^ input >> 32)
}

func HashCodeFloat(input float32) int32 {
	return int32(math.Float32bits(input))
}

func HashCodeDouble(input float64) int32 {
	value := int64(math.Float64bits(input))
	return int32(value ^ value >> 32)
}
