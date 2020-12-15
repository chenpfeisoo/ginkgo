package B_test

import (
	. "github.com/onsi/ginkgo/integration/ginko-run-1/root/B"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("B", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
