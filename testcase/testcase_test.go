package testcase_test

import (
	"fmt"

	"github.com/onsi/ginkgo"
)

var _ = ginkgo.BeforeFarmWork(func() {
	fmt.Println("BeforeFarmWork")
})
var _ = ginkgo.AfterFarmWork(func() {
	fmt.Println("AfterFarmWork")
})
var _ = ginkgo.BeforeSuite(func() {
	fmt.Println("BeforeSuite")
})
var _ = ginkgo.AfterSuite(func() {
	fmt.Println("AfterSuite")
})
var _ = ginkgo.BeforeEach(func() {
	fmt.Println("BeforeEach")
})
var _ = ginkgo.AfterEach(func() {
	fmt.Println("AfterEach")
})

var _ = ginkgo.Describe("Testcase", func() {
	ginkgo.It("", func() {
		fmt.Println("Testcase")
	})
})
