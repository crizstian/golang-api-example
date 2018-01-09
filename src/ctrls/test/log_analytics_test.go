package ctrls_test

import (
	. "easycast/src/ctrls/analytics"
	. "easycast/src/db"
	"encoding/json"
	"fmt"

	mgo "gopkg.in/mgo.v2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	db       *mgo.Database
	logAnlyt LogAnalytics
)

var _ = Describe("Ctrls Log Analytics", func() {

	Describe("Log Analytics", func() {
		It("Can init db and and bind it to log analytics interface", func() {
			session, err := MongoDB()
			Expect(err).ShouldNot(HaveOccurred())

			// defer session.Close()
			db = session.DB("easycast")
			Expect(db).ToNot(Equal(nil))

			logAnlyt = InitLogAnalytics(db)
			Expect(logAnlyt).ToNot(Equal(nil))
		})

		It("Can get bandwidth by account ", func() {
			accountEntry := "a1s2d3f4"
			startTime, endTime := "2017-07-03T00:00:00.000Z", "2017-07-30T00:00:00.000Z"

			result, err := logAnlyt.GetBandwidthByAccount(accountEntry, startTime, endTime)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(result)).Should(BeNumerically(">=", 0), "should return bandwidth by account")

			pretty, err := json.MarshalIndent(result, "", "  ")
			Expect(err).ShouldNot(HaveOccurred())
			fmt.Print(string(pretty))
		})

		It("Can get locations by account", func() {
			accountEntry := "a1s2d3f4"
			startTime, endTime := "2017-07-03T00:00:00.000Z", "2017-07-30T00:00:00.000Z"

			result, err := logAnlyt.GetLocationsByAccount(accountEntry, startTime, endTime)
			Expect(err).ShouldNot(HaveOccurred())
			// Expect(len(result)).Should(BeNumerically(">=", 0), "should return bandwidth by stream divided in times")

			pretty, err := json.MarshalIndent(result, "", "  ")
			Expect(err).ShouldNot(HaveOccurred())
			fmt.Print(string(pretty))
		})
	})
})
