package impl

import (
	b64 "encoding/base64"
	"encoding/json"

	"github.com/runtakun/kinesis-connector-go/kcl"
)

type BasicTransformer struct {
}

func (t *BasicTransformer) ToItem(record *kcl.Record) (interface{}, uint64) {
	dec, _ := b64.StdEncoding.DecodeString(record.Data)
	return dec, uint64(len(dec))
}

func (t *BasicTransformer) FromItem(item interface{}) (interface{}, error) {
	return item, nil
}

type JsonTransformer struct {
	*BasicTransformer
}

func (t *JsonTransformer) ToItem(record *kcl.Record) (interface{}, uint64) {
	decorded, recordSize := t.BasicTransformer.ToItem(record)

	var m interface{}
	if err := json.Unmarshal(decorded.([]byte), m); err != nil {
		panic(err)
	}
	return m, recordSize
}

func (t *JsonTransformer) FromItem(item interface{}) (interface{}, error) {
	return item, nil
}
