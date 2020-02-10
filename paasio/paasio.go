package paasio

import (
	"io"
	"sync"
)

// counter It keeps the counter of ops and ops
type counter struct {
	// ops Number of operations
	ops int

	// bytes Number of ops read/write.
	bytes int64

	//implicit definition
	sync.RWMutex
}

// addCount adds the number of ops read
func (counter *counter) addCount(n int64) {
	counter.Lock()
	defer counter.Unlock()

	if n <= 0 {
		return
	}

	// Add the ops read and increase the count of operation.
	// bytes Number of ops.
	counter.bytes += n
	counter.ops++
}

// wrtieCounter Keeps the count of write operation in counter
type writeCounter struct {
	counter
	writer io.Writer
}

// readCounter Keeps the count of ops read
type readCounter struct {
	counter
	// Reader is the interface that wraps the basic Read method.
	reader io.Reader
}

// ReadWriteCount read write count
type ReadWriteCount struct {
	WriteCounter
	ReadCounter
}

func (rc *readCounter) Read(p []byte) (n int, err error) {
	// Read the ops
	n, err = rc.reader.Read(p)

	// Add the count.
	rc.addCount(int64(n))
	return n, err
}

// ReadCount Implementation of intf method
func (rc *readCounter) ReadCount() (int64, int) {
	rc.Lock()
	defer rc.Unlock()

	return rc.counter.bytes, rc.counter.ops
}

func (wc *writeCounter) Write(p []byte) (n int, err error) {
	n, err = wc.writer.Write(p)

	wc.addCount(int64(n))

	return n, err
}

func (wc *writeCounter) WriteCount() (n int64, ops int) {
	wc.Lock()
	defer wc.Unlock()

	return wc.counter.bytes, wc.counter.ops
}

// NewWriteCounter New write counter
func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{
		writer:  w,
		counter: counter{ops: 0, bytes: 0},
	}
}

// NewReadWriteCounter New read-write counter
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {

	return &ReadWriteCount{
		NewWriteCounter(rw),
		NewReadCounter(rw),
	}
}

// NewReadCounter New read counter
func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{
		reader:  r,
		counter: counter{ops: 0, bytes: 0},
	}
}
