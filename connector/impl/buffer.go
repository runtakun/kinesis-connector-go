package impl

import (
	"sync/atomic"
	"time"
)

type BasicMemoryBuffer struct {
	buffer               []interface{}
	numRecordsToBuffer   int
	bytesToBuffer        uint64
	millisecondsToBuffer int64

	firstSeqNum       string
	lastSeqNum        string
	byteCount         uint64
	previousFlushTime time.Time
}

func (b *BasicMemoryBuffer) ConsumeRecord(record interface{}, recordSize uint64, seqNum string) {
	if len(b.buffer) == 0 {
		b.firstSeqNum = seqNum
	}
	b.lastSeqNum = seqNum
	b.buffer = append(b.buffer, record)
	atomic.AddUint64(&b.byteCount, recordSize)
}

func (b *BasicMemoryBuffer) Clear() {
	b.buffer = b.buffer[0:0]
	atomic.AddUint64(&b.byteCount, uint64(0))
	b.previousFlushTime = time.Now()
}

func (b *BasicMemoryBuffer) ShouldFlush() bool {
	timelapse := int64(time.Since(b.previousFlushTime) / time.Millisecond)
	return len(b.buffer) > 0 && (len(b.buffer) >= b.numRecordsToBuffer || atomic.LoadUint64(&b.byteCount) >= b.bytesToBuffer || timelapse >= b.millisecondsToBuffer)
}

func (b *BasicMemoryBuffer) GetRecords() []interface{} {
	return nil
}

func (b *BasicMemoryBuffer) GetLastSequenceNumber() string {
	return b.lastSeqNum
}

func NewBasicMemoryBuffer(buffer []interface{}, numRecordsToBuffer int, bytesToBuffer uint64, millisecondsToBuffer int64) *BasicMemoryBuffer {
	return &BasicMemoryBuffer{buffer: buffer, numRecordsToBuffer: numRecordsToBuffer, bytesToBuffer: bytesToBuffer, millisecondsToBuffer: millisecondsToBuffer}
}
