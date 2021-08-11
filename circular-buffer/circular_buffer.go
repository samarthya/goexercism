package circular

import (
	"errors"
	"log"
)

// QueueElement Element in the queue
type QueueElement struct {
	// byte to store
	B byte
	O bool
}

// Buffer represent buffer
type Buffer struct {
	B                        []QueueElement
	Head, Tail, Size, Length int
}

type MyBuffer interface {
	NewBuffer(size int) *Buffer
	ReadByte() (byte, error)
	WriteByte(c byte) error
	Overwrite(c byte)
	Reset() // put buffer in an empty state
}

//Reset resets the pen drive
func (b *Buffer) Reset() {
	b.Head = 0
	b.Tail = 0
	b.Length = 0
	return
}

//Generates a new Buffer
func NewBuffer(size int) (b *Buffer) {
	log.Println("initializing the buffer: ", size)
	b = &Buffer{B: make([]QueueElement, size), Head: 0, Tail: 0, Size: size, Length: 0}
	return
}

//WriteByte writes to the queue
func (b *Buffer) WriteByte(c byte) error {
	if (b.Length >= 0) && (b.Length < b.Size) {
		b.Length++
		log.Println("writing at ", b.Tail)
		// Add the element
		b.B[b.Tail].B = c
		//Element is occupied now
		b.B[b.Tail].O = true
		// Move the tail to the next empty location.
		b.Tail = ((b.Tail + 1) % b.Size)
		return nil
	}

	return errors.New("no space to enter the new element")
}

// ReadByte reads the bytes
func (b *Buffer) ReadByte() (c byte, e error) {
	if b.Length != 0 {

		// Start element
		start := b.Head

		// Read the Head
		if b.B[start].O { //Should be occupied
			c = b.B[start].B

			// Set occupied to false
			b.B[start].O = false
			b.B[start].B = 0

			//reduce the length
			b.Length--

			// Move to the next element
			b.Head = ((b.Head + 1) % b.Size)

			return c, nil
		}
	}

	return 0, errors.New("no element in the queue")
}

//Overwrite overwrite the byte
func (b *Buffer) Overwrite(c byte) {
	log.Println("overwrite called")
	log.Printf("Value %v", b.B)
	if b.Length == 0 {
		log.Println("no element present.. calling WriteByte")
		b.WriteByte(c)
		return
	}

	if b.Size == b.Length {
		log.Println("queue is full so overwriting head")
		b.B[b.Head].B = c
		b.Head = (b.Head + 1) % b.Size
		return
	}

	// Find the empty slow between Head and
	log.Println("find the first empty location")
	for i := b.Head; i < b.Size; i++ {
		// Add to the first empty location found
		if !b.B[i].O {
			log.Println("writing at ", i)
			b.B[i].B = c
			b.B[i].O = true
			b.Tail = ((i + 1) % b.Size)
			b.Length++
			return
		}
	}

	//if it reaches here that means no empty slot
	//overwrite the oldest
	for i := b.Head; i < b.Size; i++ {
		log.Println("writing at (Head) ", i)
		b.B[i].B = c
		b.B[i].O = true
		//Move the head to the next oldest
		b.Head = (b.Head + 1) % b.Size
	}

	// If it has reached here it means
	return
}
