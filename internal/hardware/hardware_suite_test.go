package hardware_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHardware(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hardware Spec")
}
