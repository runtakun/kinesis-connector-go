package kcl

import "fmt"

type message struct {
	Action     string    `json:"action"`
	ShardID    *string   `json:"shardId,omitempty"`
	Records    []*Record `json:"records,omitempty"`
	Checkpoint *string   `json:"checkpoint,omitempty"`
	Error      *string   `json:"error,omitempty"`
	Reason     *string   `json:"reason,omitempty"`
}

type messageHandler struct {
	ih *ioHandler
	rp RecordProcessor
	cp *CheckPointer
}

func (mh *messageHandler) doAction() error {

	msg, err := mh.ih.receiveMessage()
	if err != nil {
		return err
	}

	switch msg.Action {
	case "initialize":
		err = mh.rp.Initialize(*msg.ShardID)
	case "processRecords":
		mh.cp.checkPointAllowed = true
		err = mh.rp.ProcessRecords(msg.Records, mh.cp)
	case "shutdown":
		mh.cp.checkPointAllowed = true
		err = mh.rp.Shutdown(*msg.Reason, mh.cp)
	default:
		err = fmt.Errorf("invalid action: %s", msg.Action)
	}

	if err != nil {
		return err
	}

	return mh.ih.sendStatus(msg.Action)
}
