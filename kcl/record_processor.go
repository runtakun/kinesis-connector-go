package kcl

type RecordProcessor interface {
	Initialize(shardID string) error
	ProcessRecords(records []*Record, cp *CheckPointer) error
	Shutdown(reason string, cp *CheckPointer) error
}
