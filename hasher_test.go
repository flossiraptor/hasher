package hasher

import (
	"testing"
)

func TestHashCodeHasherInterface(t *testing.T) {
	subject := new(hasherInterfaceTest)
	if HashCode(subject) != 5 {
		t.Error("The test interface should provide the hashcode 5")
	}
}

func TestHashCodeBool(t *testing.T) {
	if HashCode(false) != 0 {
		t.Error("FALSE must hash to 0")
	}
	if HashCode(true) != 1 {
		t.Error("TRUE must hash to 1")
	}
}

func TestHashCodeInt(t *testing.T) {
	for key, value := range testCasesInt32() {
		if HashCode(int(key)) != value {
			t.Errorf("The number '%d' should hash to the code '%d'", key, value)
		}
	}
}
func TestHashCodeInt8(t *testing.T) {
	for key, value := range testCasesInt8() {
		if HashCode(key) != value {
			t.Errorf("The number '%d' should hash to the code '%d'", key, value)
		}
	}
}

func TestHashCodeInt16(t *testing.T) {
	for key, value := range testCasesInt16() {
		if HashCode(key) != value {
			t.Errorf("The number '%d' should hash to the code '%d'", key, value)
		}
	}
}

func TestHashCodeInt32(t *testing.T) {
	for key, value := range testCasesInt32() {
		if HashCode(key) != value {
			t.Errorf("The number '%d' should hash to the code '%d'", key, value)
		}
	}
}

func TestHashCodeInt64(t *testing.T) {
	for key, value := range testCasesInt64() {
		if HashCode(key) != value {
			t.Errorf("The number '%d' should hash to the code '%d'", key, value)
		}
	}
}

func TestHashCodeUint(t *testing.T) {
	for key, value := range testCasesUint32() {
		if HashCode(int(key)) != value {
			t.Errorf("The number '%d' should hash to the code '%d'", key, value)
		}
	}
}
func TestHashCodeUint8(t *testing.T) {
	for key, value := range testCasesUint8() {
		if HashCode(key) != value {
			t.Errorf("The number '%d' should hash to the code '%d'", key, value)
		}
	}
}

func TestHashCodeUint16(t *testing.T) {
	for key, value := range testCasesUint16() {
		if HashCode(key) != value {
			t.Errorf("The number '%d' should hash to the code '%d'", key, value)
		}
	}
}

func TestHashCodeUint32(t *testing.T) {
	for key, value := range testCasesUint32() {
		if HashCode(key) != value {
			t.Errorf("The number '%d' should hash to the code '%d'", key, value)
		}
	}
}

func TestHashCodeUint64(t *testing.T) {
	for key, value := range testCasesUint64() {
		if HashCode(key) != value {
			t.Errorf("The number '%d' should hash to the code '%d'", key, value)
		}
	}
}
func TestHashCodeFloat32(t *testing.T) {
	for key, value := range testCasesFloat32() {
		if HashCode(key) != value {
			t.Errorf("The number '%f' should hash to the code '%d'", key, value)
		}
	}
}
func TestHashCodeFloat64(t *testing.T) {
	for key, value := range testCasesFloat64() {
		if HashCode(key) != value {
			t.Errorf("The number '%f' should hash to the code '%d'", key, value)
		}
	}
}
func TestHashCodeString(t *testing.T) {
	for key, value := range testCasesString() {
		if HashCode(key) != value {
			t.Errorf("The string '%s' should hash to the code '%d'", key, value)
		}
	}
}

func testCasesInt8() map[int8]int32 {
	return map[int8]int32{
		-127: -127,
		-64: -64,
		0: 0,
		127: 127,
	}
}
func testCasesInt16() map[int16]int32 {
	return map[int16]int32{
		-32768: -32768,
		-127: -127,
		-64: -64,
		0: 0,
		127: 127,
		32767: 32767,
	}
}
func testCasesInt32() map[int32]int32 {
	return map[int32]int32{
		-2147483648: -2147483648,
		-127: -127,
		-64: -64,
		0: 0,
		127: 127,
		2147483647: 2147483647,
	}
}
func testCasesInt64() map[int64]int32 {
	return map[int64]int32{
		-9223372036854775808: -2147483648,
		-555: 554,
		0: 0,
		555: 555,
		9223372036854775807: -2147483648,
	}
}

func testCasesUint8() map[uint8]int32 {
	return map[uint8]int32{
		0: 0,
		64: 64,
		255: 255,
	}
}
func testCasesUint16() map[uint16]int32 {
	return map[uint16]int32{
		0: 0,
		64: 64,
		255: 255,
		65535: 65535,
	}
}
func testCasesUint32() map[uint32]int32 {
	return map[uint32]int32{
		0: 0,
		64: 64,
		255: 255,
		65535: 65535,
		2147483647: 2147483647,
		4000000000: -294967296,
		4294967295: -1,
	}
}
func testCasesUint64() map[uint64]int32 {
	return map[uint64]int32{
		0: 0,
		64: 64,
		255: 255,
		65535: 65535,
		2147483647: 2147483647,
		4000000000: -294967296,
		4294967295: -1,
		9223372036854775807: -2147483648,
	}
}

func testCasesFloat32() map[float32]int32 {
	return map[float32]int32{
		-0.5: -1090519040,
		0: 0,
		0.5: 1056964608,
		1: 1065353216,
	}
}
func testCasesFloat64() map[float64]int32 {
	return map[float64]int32{
		-5: -1072431104,
		-1: -1074790400,
		-0.5: -1075838976,
		0: 0,
		0.5: 1071644672,
		1: 1072693248,
		5: 1075052544,
	}
}

func testCasesString() map[string]int32 {
	return map[string]int32{
		"Golang": 2138494710,
		"Hi": 2337,
		"Hello": 69609650,
		"Supercalifragilisticexpialidocious": -1673062653,
		"The quick brown fox jumps over the lazy dog": -609428141,
	}
}

type hasherInterfaceTest struct {
}
func (ht *hasherInterfaceTest) HashCode() int32 {
	return 5
}
