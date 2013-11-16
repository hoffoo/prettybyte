package prettybyte

// Tiny tool to format arrays as pretty strings

import (
	"fmt"
	"io"
)

// Print a byte array in hex format - 16 bytes per line
func Pretty(b []byte, out io.Writer) {

	var offset int = 16

	if offset > len(b) {
		offset = len(b)
	}

	for idx := 0; idx < len(b); idx += 16 {

		out.Write([]byte(fmt.Sprintf("% 0x\n", b[idx:offset])))
		if offset += 16; offset > len(b) {
			offset = offset - (offset - len(b))
		}
	}
}
