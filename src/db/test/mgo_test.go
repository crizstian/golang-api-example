package test_test

import (
	. "easycast/src/db"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mongo DB Test", func() {

	Describe("Testing Mongo Queries", func() {
		It("Can connect to mongodb", func() {
			session, err := MongoDB()
			Expect(err).ShouldNot(HaveOccurred())
			defer session.Close()

			c := session.DB("easycast").C("analytics")
			count, err2 := c.Count()
			Expect(err2).ShouldNot(HaveOccurred())

			Expect(count).To(BeNumerically(">=", 0), "should return number of documents in collection")
			fmt.Println("count", count)
			fmt.Println("Test Ok")
		})
	})
})
