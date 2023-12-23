package test

import (
	"ConfBackend/network"
	"ConfBackend/util"
	"ConfBackend/util/coord"
	"encoding/binary"
	"fmt"
	"math/rand"
	"testing"
)

func TestLineaerCalc(t *testing.T) {
	x, y, z, err := coord.CalcLinearCoord([][]float64{{2.0, 0.0, 0.0}, {0.0, 0.0, 0.0}}, [][]float64{{2.0}, {2.82842}})
	fmt.Println(x, y, z, err)
	if err != nil {
		return
	}
}

func TestBigEndian(t *testing.T) {
	a := binary.BigEndian.Uint16([]byte{1, 224})
	// print the bits
	fmt.Printf("Binary representation of %d: %016b\n", a, a)
}

func TestParseMsgBody(t *testing.T) {
	b1 := []byte{0, 5}
	b2 := []byte{0, 1}
	b3 := []byte{0, 5}
	b4 := byte('a')
	b5 := byte(1)
	b6 := []byte{1, 2, 3, 4}
	b7 := []byte{'a', 'f', 'e'}
	b8 := []byte{0, 0, 0, 0, 0}
	// b9 contains 260 bytes
	b9 := []byte{1, 3, 5, 7, 9}
	//for i := 0; i < 260; i++ {
	//	b9 = append(b9, 1)
	//}

	// concat them all
	b := append(b1, b2...)
	b = append(b, b3...)
	b = append(b, b4)
	b = append(b, b5)
	b = append(b, b6...)
	b = append(b, b7...)
	b = append(b, b8...)
	b = append(b, b9...)

	aa, _ := network.ParseMsgBody(b)
	print(aa)

}

func TestSaveRedisMsgBody(t *testing.T) {
	b1 := []byte{0, 5}
	b2 := []byte{0, 1}
	b3 := []byte{0, 2}
	b4 := byte('a')
	b5 := byte(1)
	b6 := []byte{1, 2, 3, 4}
	b7 := []byte{'a', 'f', 'e'}
	b8 := []byte{0, 0, 0, 0, 0}
	// b9 contains 260 bytes
	b9 := []byte{1, 3, 5, 7, 9}
	//for i := 0; i < 260; i++ {
	//	b9 = append(b9, 1)
	//}

	// concat them all
	b := append(b1, b2...)
	b = append(b, b3...)
	b = append(b, b4)
	b = append(b, b5)
	b = append(b, b6...)
	b = append(b, b7...)
	b = append(b, b8...)
	b = append(b, b9...)

	network.HandleSinglePacket(b)

}

func TestHashGen(t *testing.T) {
	fmt.Println(string(util.FourDigitSHA256([]byte("test"))))
}

// test GenerateMessageChunks
func TestGenerateMessageChunks(t *testing.T) {
	// generate random bytes
	var randomBytes []byte

	randomBytes = make([]byte, 1000)
	rand.Read(randomBytes)

	res := network.GenerateMessageChunks('a', randomBytes, 1)
	fmt.Println(res)
}

// test util.Uint16ToTwoBytes
func TestUint16ToTwoBytes(t *testing.T) {
	fmt.Println(util.Uint16ToTwoBytes(480))
}

func TestUtf8(t *testing.T) {
	str := "你好"
	var b []byte
	b = append(b, str...)
	fmt.Println(b)
}
