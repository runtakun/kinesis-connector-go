package connector

import (
	"errors"
	"fmt"
	"time"

	"github.com/runtakun/kinesis-connector-go/kcl"
)

type Filter interface {
	KeepRecord(v interface{}) bool
}

type Transformer interface {
	ToItem(record *kcl.Record) (interface{}, uint64)
	FromItem(item interface{}) (interface{}, error)
}

type Buffer interface {
	ConsumeRecord(record interface{}, recordSize uint64, sequenceNumber string)
	ShouldFlush() bool
	GetRecords() []interface{}
	GetLastSequenceNumber() string
	Clear()
}

type Emitter interface {
	Emit(buffer Buffer) []interface{}
	Fail(items []interface{})
	Shutdown()
}

type ConnectorProcessor struct {
	transformer Transformer
	filter      Filter
	buffer      Buffer
	emitter     Emitter

	shardID         string
	shutdown        bool
	initialized     bool
	retryLimit      int
	backoffInterval int
}

func (p *ConnectorProcessor) Initialize(shardID string) error {
	p.shardID = shardID
	p.initialized = true
	return nil
}

func (p *ConnectorProcessor) ProcessRecords(records []*kcl.Record, cp *kcl.CheckPointer) error {

	if p.shutdown {
		// TODO LOG
	}

	if !p.initialized {
		return errors.New("Record processor not initialized")
	}

	for _, record := range records {
		transFormedRecord, recordSize := p.transformer.ToItem(record)
		p.filterAndBufferRecord(transFormedRecord, record, recordSize)
	}

	if p.buffer.ShouldFlush() {
		emitItems := p.transformToOutput(p.buffer.GetRecords())
		p.emit(cp, emitItems)
	}

	return nil
}

func (p *ConnectorProcessor) filterAndBufferRecord(transformedRecord interface{}, record *kcl.Record, recordSize uint64) {
	if p.filter.KeepRecord(transformedRecord) {
		p.buffer.ConsumeRecord(transformedRecord, recordSize, record.SequenceNumber)
	}
}

func (p *ConnectorProcessor) transformToOutput(items []interface{}) []interface{} {
	var outputs []interface{}

	for _, item := range items {
		output, err := p.transformer.FromItem(item)
		if err != nil {
			// TODO LOG
		}
		outputs = append(outputs, output)
	}

	return outputs
}

func (p *ConnectorProcessor) emit(cp *kcl.CheckPointer, emitItems []interface{}) {
	unprocessed := make([]interface{}, len(emitItems))
	copy(emitItems, unprocessed)

	numTries := p.retryLimit + 1
	for i := 0; i < numTries; i++ {
		unprocessed = p.emitter.Emit(NewUnmodifiableBuffer(p.buffer, unprocessed))
		if len(unprocessed) == 0 {
			break
		}
		time.Sleep(time.Duration(p.backoffInterval) * time.Second)
	}
	if len(unprocessed) > 0 {
		p.emitter.Fail(unprocessed)
	}

	lastSeqNum := p.buffer.GetLastSequenceNumber()
	p.buffer.Clear()

	if lastSeqNum != "" {
		if err := cp.CheckPointSeq(lastSeqNum); err != nil {
			p.emitter.Fail(unprocessed)
		}
	}
}

func (p *ConnectorProcessor) Shutdown(reason string, cp *kcl.CheckPointer) error {
	if p.shutdown {
		// TODO LOG
	}
	switch reason {
	case "TERMINATE":
		p.emit(cp, p.transformToOutput(p.buffer.GetRecords()))
		if err := cp.CheckPointAll(); err != nil {
			// TODO LOG
		}
	case "ZOMBIE":
	// do nothing
	default:
		return fmt.Errorf("invalid shutdown reason: %s", reason)
	}
	p.emitter.Shutdown()
	p.shutdown = true

	return nil
}
