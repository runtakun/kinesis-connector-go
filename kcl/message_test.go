package kcl_test

import (
	"fmt"

	"code.google.com/p/go-uuid/uuid"

	. "github.com/runtakun/kinesis-connector-go/kcl"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

type testProcessor struct {
}

func (p *testProcessor) Initialize(shardID string) error {
	return nil
}
func (p *testProcessor) ProcessRecords(records []*Record, cp *CheckPointer) error {
	return nil
}

func (p *testProcessor) Shutdown(reason string, cp *CheckPointer) error {
	return nil
}

type ginkgoLogger struct {
}

func (l *ginkgoLogger) Printf(format string, v ...interface{}) {
	GinkgoWriter.Write([]byte(fmt.Sprintf(format, v...)))
	GinkgoWriter.Write([]byte{0xa}) //new line
}

var _ = Describe("Message", func() {

	Describe("receive normal message", func() {
		var stdin *gbytes.Buffer
		var stdout *gbytes.Buffer
		var processor RecordProcessor

		var shardID string

		BeforeEach(func() {
			SetLogger(&ginkgoLogger{})

			stdin = gbytes.NewBuffer()
			stdout = gbytes.NewBuffer()
			processor = &testProcessor{}

			shardID = uuid.NewRandom().String()
			go RunIO(processor, stdin, stdout)
		})

		It("received intialize action", func() {
			stdin.Write([]byte(fmt.Sprintf(`{"action": "initialize", "shardId": "%s"}`, shardID)))
			Eventually(stdout).Should(gbytes.Say(`{"action":"status","responseFor":"initialize"}`))
		})

		It("received processRecords action", func() {
			stdin.Write([]byte(fmt.Sprintf(`{"action": "processRecords", "records": [{"data": "data", "partitionKey": "pk1", "sequenceNumber": "001"}, {"data": "data", "partitionKey": "pk1", "sequenceNumber": "002"}]}`)))
			Eventually(stdout).Should(gbytes.Say(`{"action":"status","responseFor":"processRecords"}`))
		})

		It("received shutdown action", func() {
			stdin.Write([]byte(fmt.Sprintf(`{"action": "shutdown", "reason": "TERMINATE"}`)))
			Eventually(stdout).Should(gbytes.Say(`{"action":"status","responseFor":"shutdown"}`))
		})
	})
})
