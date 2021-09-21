package circular

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// Buffer represent buffer
type Buffer struct {
	B    []byte
	Size int
}

type MyBuffer interface {
	NewBuffer(size int) *Buffer
	ReadByte() (byte, error)
	WriteByte(c byte) error
	Overwrite(c byte)
	Reset() // put buffer in an empty state
}

func (b *Buffer) String() string {
	var m strings.Builder
	fmt.Fprintf(&m, "size=%d ", b.Size)
	for i, v := range b.B {
		fmt.Fprintf(&m, "[%d]=%c ", i, v)
	}
	return m.String()
}

//Reset resets the pen drive
func (b *Buffer) Reset() {
	// Reset all values
	b.B = make([]byte, 0)
	return
}

//Generates a new Buffer
func NewBuffer(size int) (b *Buffer) {
	log.Println("initializing the buffer: ", size)
	b = &Buffer{B: make([]byte, 0), Size: size}
	return
}

//WriteByte writes to the queue
func (b *Buffer) WriteByte(c byte) error {
	log.Println("current size ", len(b.B))
	if len(b.B) < b.Size {
		b.B = append(b.B, c)
		return nil
	}

	return errors.New("no space to enter the new element")
}

// ReadByte reads the bytes
func (b *Buffer) ReadByte() (c byte, e error) {
	defer log.Println("read invoked ")
	if len(b.B) > 0 {
		c = b.B[0]
		// Remove the first element
		b.B = b.B[1:]
		return
	}

	return 0, errors.New("no element in the queue")
}

//Overwrite overwrite the byte
func (b *Buffer) Overwrite(c byte) {
	log.Println("overwrite called")

	if e := b.WriteByte(c); e != nil {
		log.Println("Q>", b)
		log.Printf("adding %c", c)
		b.B = append(b.B[1:], c)
		log.Println("Q>", b)
	}

	return
}
