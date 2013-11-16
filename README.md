prettybyte
==========

Print bytes in a readable format

```go

func main() {
	b := make([]byte, 45)
	b[0] = 0x01
	b[44] = 0xAF
	

	prettybyte.Pretty(b, os.Stdout)
}

```
01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00 00 00 00 00 af
