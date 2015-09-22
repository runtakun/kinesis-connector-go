package kinesis_connector_go_test

import (
	"bytes"
	"fmt"

	"code.google.com/p/go-uuid/uuid"

	. "github.com/runtakun/kinesis-connector-go/kcl"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
	BeforeEach(func() {
		SetLogger(&ginkgoLogger{})
	})

	Context("receive normal message", func() {

		var stdout *bytes.Buffer
		var processor RecordProcessor

		var shardID string

		BeforeEach(func() {
			stdout = new(bytes.Buffer)
			processor = &testProcessor{}

			shardID = uuid.NewRandom().String()
		})

		It("received intialize action", func() {
			stdin := bytes.NewBufferString(fmt.Sprintf(`{"action": "initialize", "shardId": "%s"}`, shardID))
			err := RunFile(processor, stdin, stdout)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(stdout.Bytes()).Should(MatchJSON(`{"action": "status", "responseFor": "initialize"}`))
		})

		It("received processRecord action", func() {
			stdin := bytes.NewBufferString(fmt.Sprintf(`{"action": "processRecords", "data": "data", "partitionKey": "pk1", "sequenceNumber": "001"}`))
			err := RunFile(processor, stdin, stdout)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(stdout.Bytes()).Should(MatchJSON(`{"action": "status", "responseFor": "processRecords"}`))
		})

		It("received shutdown action", func() {
			stdin := bytes.NewBufferString(fmt.Sprintf(`{"action": "shutdown", "reason": "TERMINATE"}`))
			err := RunFile(processor, stdin, stdout)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(stdout.Bytes()).Should(MatchJSON(`{"action": "status", "responseFor": "shutdown"}`))
		})
	})
})
