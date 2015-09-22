package connector

type UnmodifiableBuffer struct {
	buffer  Buffer
	records []interface{}
}

func (b *UnmodifiableBuffer) ConsumeRecord(record interface{}, recordSize uint64, seqNum string) {
	panic("This is an unmodifiable buffer")
}

func (b *UnmodifiableBuffer) ShouldFlush() bool {
	return b.buffer.ShouldFlush()
}

func (b *UnmodifiableBuffer) GetRecords() []interface{} {
	return b.records
}

func (b *UnmodifiableBuffer) GetLastSequenceNumber() string {
	return b.buffer.GetLastSequenceNumber()
}

func (b *UnmodifiableBuffer) Clear() {
	panic("This is an unmodifiable buffer")
}

func NewUnmodifiableBuffer(buffer Buffer, records []interface{}) *UnmodifiableBuffer {
	return &UnmodifiableBuffer{buffer, records}
}
