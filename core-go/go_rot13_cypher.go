package main


import (
	"io";
	"os";
	"strings";
	"fmt"
)

type Rot13Reader struct {
	reader io.Reader
}

func (r *Rot13Reader) Read (p []byte) (bytesRead int, err error){
	// chain the readers. Read from attached Reader from upstream.
	bytesRead, err = r.reader.Read(p)

	if(err != nil) {
		return bytesRead, err
	}

	for i:=0; i < bytesRead; i++ {
		p[i] = rotate13(p[i])
	}

	return bytesRead, err
}

func rotate13 (b byte) byte {
	var val byte;
	switch {
	case (b >= 'A' && b <= 'M') || (b >= 'a' && b <= 'm'):
		val = b + 13
	case (b >= 'N' && b <= 'Z') || (b >= 'n' && b <= 'z':
		val = b - 13
	}
	fmt.Print ("Rotating ", b, " to ", val, "\n")
	return val
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := Rot13Reader{s}
    fmt.Println ()
    io.Copy(os.Stdout, &r)
    fmt.Println()
}