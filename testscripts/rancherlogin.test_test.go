package testscripts_test

import (
	"crypto/tls"
	"fmt"

	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("/login-200", func() {
	It("rancher login test", func() {
		client := resty.New()
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("Accept", "application/json").
			SetBody(`{"username":"admin","password":"Happytestingrancher"}`).
			Post("https://test.rancher/v3-public/localProviders/local?action=login")

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		Expect(resp.StatusCode()).To(Equal(201))
		Expect(err).To(BeNil(), "Error making the request")
		Expect(resp.StatusCode()).To(Equal(201), "Expected status code 200")
		Expect(resp.String()).To(ContainSubstring("token"), "Expected response to contain a token")

	})

})
