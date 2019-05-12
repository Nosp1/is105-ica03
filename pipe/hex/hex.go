package hex

import (
	"encoding/hex"
	"fmt"
)

func HexReturn(s string) string {
	txt := []byte(s)
	dst := make([]byte, hex.EncodedLen(len(txt)))
	hex.Encode(dst, txt)
	return fmt.Sprintf("%s", dst)
}

func HexReturnBytes(b []byte) []byte {
	txt := b
	dst := make([]byte, hex.EncodedLen(len(txt)))
	hex.Encode(dst, txt)
	return dst
}