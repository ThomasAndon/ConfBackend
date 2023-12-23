package util

func Uint16ToTwoBytes(i uint16) []byte {
	return []byte{byte(i >> 8), byte(i)}
}
