package C_test

import (
	. "github.com/onsi/ginkgo/integration/ginko-run-1/root/C"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("C", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
