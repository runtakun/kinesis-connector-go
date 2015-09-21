package kcl

func Run(rp RecordProcessor) error {

	ih := NewCLIHandler()
	cp := &CheckPointer{ih, false}
	mh := &messageHandler{ih, rp, cp}

	return mh.doAction()
}
