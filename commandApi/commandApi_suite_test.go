package commandApi_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCommandApi(t *testing.T) {
	BeforeSuite(func() {
		os.Setenv("ENV", "testing")
	})

	RegisterFailHandler(Fail)
	RunSpecs(t, "CommandApi Suite")
}
