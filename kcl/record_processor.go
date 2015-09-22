package kcl

type Record struct {
	Data           string `json:"data"`
	PartitionKey   string `json:"partitionKey"`
	SequenceNumber string `json:"sequenceNumber"`
}

type RecordProcessor interface {
	Initialize(shardID string) error
	ProcessRecords(records []*Record, cp *CheckPointer) error
	Shutdown(reason string, cp *CheckPointer) error
}
