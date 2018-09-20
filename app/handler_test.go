package app_test

import (
	"io/ioutil"

	"bitbucket.org/zkrhm-fdn/fire-starter/app"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zkrhm/testhelper"
)

var _ = Describe("Handlers", func() {
	app := app.NewApp()
	It("Says Hello", func() {
		test := testhelper.NewHttpTest("POST", "/hello", "", app.Hello)
		rr, err := test.DoRequestTest()

		Expect(err).ShouldNot(HaveOccurred())
		Expect(rr.Code).Should(Equal(200))
		Expect(rr.Header().Get("Content-Type")).Should(Equal("application/json"))
		Expect(ioutil.ReadAll(rr.Body)).Should(MatchJSON(`{"code":200,"message":"Hello"}`))
	})
})
