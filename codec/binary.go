package codec

import (
	"fmt"
	"os"
	"time"
)

type Message struct {
	data    []byte
	counter uint64
}

func (m *Message) Data() []byte {
	return m.data
}

func MakeMsg(f *os.File, size uint64) (Message, error) {
	data := make([]byte, size)
	m := Message{data, 0}
	_, e := f.Read(data)

	return m, e

}

func MsToTime(ms uint64) time.Time {
	//return time.Unix(0, int64(ms)*int64(time.Millisecond))
	return time.Unix(int64(ms)/1000, 0)
}

func (m *Message) Get(inc uint64) []byte {
	t := m.counter + inc
	if t > uint64(len(m.data)) {
		//return m.data, false
		av := uint64(len(m.data)) - m.counter
		panic(fmt.Sprintf("Only %v available but %v requested.", av, inc))
	}
	a := m.data[m.counter:t]
	m.counter = t
	return a
}

func ToInt(data []byte) uint64 {
	//i, j := binary.Uvarint(data)
	i := Uvarint(Rev(data))
	//fmt.Printf("%s %v %v\n", data, data, i)
	return i
	//fmt.Printf("%s", MsToTime(i))
}

func Rev(buf []byte) []byte {
	l := len(buf)
	for i := 0; i < l/2; i++ {
		buf[i], buf[l-i-1] = buf[l-i-1], buf[i]
	}
	return buf
}

func Uvarint(buf []byte) uint64 {
	var x uint64
	for i, b := range buf {
		x |= (uint64(b) << (uint8(i) * 8))
	}
	return x
}
