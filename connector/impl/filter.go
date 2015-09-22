package impl

import "github.com/runtakun/kinesis-connector-go/kcl"

type AllPassFilter struct {
}

func (f *AllPassFilter) KeepRecod(record *kcl.Record) bool {
	return true
}
