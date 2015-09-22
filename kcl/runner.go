package kcl

import "io"

func RunCLI(rp RecordProcessor) {

	ih := newCLIHandler()
	cp := &CheckPointer{ih, true}
	mh := &messageHandler{ih, rp, cp}

	run(mh)
}

func RunIO(rp RecordProcessor, in io.Reader, out io.Writer) {

	ih := newIOHandler(in, out)
	cp := &CheckPointer{ih, true}
	mh := &messageHandler{ih, rp, cp}

	run(mh)
}

func run(mh *messageHandler) {
	for {
		if err := mh.doAction(); err != nil {
			logger.Printf("action process err: %s", err.Error())
			break
		}
	}
}
