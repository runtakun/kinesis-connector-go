package kcl

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Data           string `json:"data"`
	PartitionKey   string `json:"partitionKey"`
	SequenceNumber string `json:"sequenceNumber"`
}

type message struct {
	Action     string    `json:"action"`
	ShardID    *string   `json:"shard_id,omitempty"`
	Records    []*Record `json:"records,omitempty"`
	Checkpoint *string   `json:"checkpoint,omitempty"`
	Error      *string   `json:"error,omitempty"`
	Reason     *string   `json:"reason,omitempty"`
}

type response struct {
	Action string `json:"action"`
	For    string `json:"responseFor,omitempty"`
}

func doAction(rp RecordProcessor, cp *CheckPointer, msg *message) error {

	var err error

	switch msg.Action {
	case "initialize":
		err = rp.Initialize(*msg.ShardID)
	case "processRecords":
		cp.checkPointAllowed = true
		err = rp.ProcessRecords(msg.Records, cp)
	case "shutdown":
		cp.checkPointAllowed = true
		err = rp.Shutdown(*msg.Reason, cp)
	default:
		err = fmt.Errorf("invalid action: %s", msg.Action)
	}

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&response{Action: "status", For: msg.Action})
}
