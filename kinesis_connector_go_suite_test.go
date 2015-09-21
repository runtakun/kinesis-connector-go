package kinesis_connector_go_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestKinesisConnectorGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "KinesisConnectorGo Suite")
}
