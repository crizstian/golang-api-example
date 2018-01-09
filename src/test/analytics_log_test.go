package apitest

import (
	"easycast/src/api"
	"easycast/src/libs"
	"encoding/json"
	"fmt"
	"net/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Log Analytics Test", func() {

	type Body struct {
		Result  interface{} `json:"result"`
		Success bool        `json:"success"`
	}
	// user id no se necesita para el test
	// customer_id no se necesita para el test
	// createdInt no se necesita para el test
	// active_plan_id no se necesita para el test
	// TODO: token test necesita actualizarce
	var (
		token = libs.TokenTestModify("a1s2d3f4")
	)

	Describe("Testing Log Analytics API", func() {

		It("Can Get Account bandwidth from logs", func() {
			var b Body

			q := make(url.Values)
			q.Set("startTime", "2017-07-03T00:00:00.000Z")
			q.Set("endTime", "2017-07-31T00:00:00.000Z")

			// ApiTest (func, method, url, body, token, urlParams)
			body := libs.ApiTest(api.GetBandwidthByAccount, "GET", "/analytics/logs/bandwidth/?"+q.Encode(), nil, token, nil)
			err := json.Unmarshal([]byte(body), &b)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(b.Success).To(Equal(true))

			pretty, err := json.MarshalIndent(b, "", "  ")
			Expect(err).ShouldNot(HaveOccurred())
			fmt.Print(string(pretty))
		})

		It("Can Get Account bandwidth per time from logs", func() {
			var b Body

			q := make(url.Values)
			q.Set("startTime", "2017-07-03T00:00:00.000Z")
			q.Set("endTime", "2017-07-31T00:00:00.000Z")

			// ApiTest (func, method, url, body, token, urlParams)
			body := libs.ApiTest(api.GetBandwidthByAccountPerTime, "GET", "/analytics/logs/bandwidth/time/?"+q.Encode(), nil, token, nil)
			err := json.Unmarshal([]byte(body), &b)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(b.Success).To(Equal(true))

			pretty, err := json.MarshalIndent(b, "", "  ")
			Expect(err).ShouldNot(HaveOccurred())
			fmt.Print(string(pretty))
		})

		It("Can Get Account bandwidth divided by streams from logs", func() {
			var b Body

			q := make(url.Values)
			q.Set("startTime", "2017-07-03T00:00:00.000Z")
			q.Set("endTime", "2017-07-31T00:00:00.000Z")

			// ApiTest (func, method, url, body, token, urlParams)
			body := libs.ApiTest(api.GetBandwidthByStreams, "GET", "/analytics/logs/bandwidth/streams/?"+q.Encode(), nil, token, nil)
			err := json.Unmarshal([]byte(body), &b)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(b.Success).To(Equal(true))

			pretty, err := json.MarshalIndent(b, "", "  ")
			Expect(err).ShouldNot(HaveOccurred())
			fmt.Print(string(pretty))
		})

		It("Can Get Account bandwidth divided by streams from logs per time", func() {
			var b Body

			q := make(url.Values)
			q.Set("startTime", "2017-07-03T00:00:00.000Z")
			q.Set("endTime", "2017-07-31T00:00:00.000Z")

			// ApiTest (func, method, url, body, token, urlParams)
			body := libs.ApiTest(api.GetBandwidthByStreamsPerTime, "GET", "/analytics/logs/bandwidth/streams/time/?"+q.Encode(), nil, token, nil)
			err := json.Unmarshal([]byte(body), &b)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(b.Success).To(Equal(true))

			pretty, err := json.MarshalIndent(b, "", "  ")
			Expect(err).ShouldNot(HaveOccurred())
			fmt.Print(string(pretty))
		})

		It("Can Get Stream bandwidth from logs", func() {
			var b Body

			q := make(url.Values)
			q.Set("startTime", "2017-07-03T00:00:00.000Z")
			q.Set("endTime", "2017-07-31T00:00:00.000Z")

			// define url
			url := "/analytics/logs/bandwidth/H8SUumhjOR6PzKzm/?"
			// definimos la url con los parametros
			urlParams := "/analytics/logs/bandwidth/:stream/?"
			// creamos un dictionary key - value, que contenga los valores
			values := map[string]string{
				"stream": "H8SUumhjOR6PzKzm",
			}
			// creamos un interface para pasar un solo parametro a la funcion ApiTest
			params := map[string]interface{}{
				"url":    urlParams,
				"values": values,
			}

			// ApiTest (func, method, url, body, token, urlParams)
			body := libs.ApiTest(api.GetBandwidthByStream, "GET", url+q.Encode(), nil, token, params)
			err := json.Unmarshal([]byte(body), &b)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(b.Success).To(Equal(true))

			pretty, err := json.MarshalIndent(b, "", "  ")
			Expect(err).ShouldNot(HaveOccurred())
			fmt.Print(string(pretty))
		})

		It("Can Get Stream bandwidth per time from logs", func() {
			var b Body

			q := make(url.Values)
			q.Set("startTime", "2017-07-03T00:00:00.000Z")
			q.Set("endTime", "2017-07-31T00:00:00.000Z")

			// define url
			url := "/analytics/logs/bandwidth/H8SUumhjOR6PzKzm/?"
			// definimos la url con los parametros
			urlParams := "/analytics/logs/bandwidth/:stream/?"
			// creamos un dictionary key - value, que contenga los valores
			values := map[string]string{
				"stream": "H8SUumhjOR6PzKzm",
			}
			// creamos un interface para pasar un solo parametro a la funcion ApiTest
			params := map[string]interface{}{
				"url":    urlParams,
				"values": values,
			}

			// ApiTest (func, method, url, body, token, urlParams)
			body := libs.ApiTest(api.GetBandwidthByStreamPerTime, "GET", url+q.Encode(), nil, token, params)
			err := json.Unmarshal([]byte(body), &b)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(b.Success).To(Equal(true))

			pretty, err := json.MarshalIndent(b, "", "  ")
			Expect(err).ShouldNot(HaveOccurred())
			fmt.Print(string(pretty))
		})

		It("Can Get Locations by account", func() {
			var b Body

			q := make(url.Values)
			q.Set("startTime", "2017-07-03T00:00:00.000Z")
			q.Set("endTime", "2017-07-31T00:00:00.000Z")

			// ApiTest (func, method, url, body, token, urlParams)
			body := libs.ApiTest(api.GetLocationsByAccount, "GET", "/analytics/logs/locations/?"+q.Encode(), nil, token, nil)
			err := json.Unmarshal([]byte(body), &b)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(b.Success).To(Equal(true))

			pretty, err := json.MarshalIndent(b, "", "  ")
			Expect(err).ShouldNot(HaveOccurred())
			fmt.Print(string(pretty))
		})

		It("Can Get Locations by account", func() {
			var b Body

			q := make(url.Values)
			q.Set("startTime", "2017-07-03T00:00:00.000Z")
			q.Set("endTime", "2017-07-31T00:00:00.000Z")

			// ApiTest (func, method, url, body, token, urlParams)
			body := libs.ApiTest(api.GetLocationsByAccountPerTime, "GET", "/analytics/logs/locations/time/?"+q.Encode(), nil, token, nil)
			err := json.Unmarshal([]byte(body), &b)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(b.Success).To(Equal(true))

			pretty, err := json.MarshalIndent(b, "", "  ")
			Expect(err).ShouldNot(HaveOccurred())
			fmt.Print(string(pretty))
		})

		It("Can Get Locations by stream", func() {
			var b Body

			q := make(url.Values)
			q.Set("startTime", "2017-07-03T00:00:00.000Z")
			q.Set("endTime", "2017-07-31T00:00:00.000Z")

			// define url
			url := "/analytics/logs/locations/H8SUumhjOR6PzKzm/?"
			// definimos la url con los parametros
			urlParams := "/analytics/logs/locations/:stream/?"
			// creamos un dictionary key - value, que contenga los valores
			values := map[string]string{
				"stream": "H8SUumhjOR6PzKzm",
			}
			// creamos un interface para pasar un solo parametro a la funcion ApiTest
			params := map[string]interface{}{
				"url":    urlParams,
				"values": values,
			}

			// ApiTest (func, method, url, body, token, urlParams)
			body := libs.ApiTest(api.GetLocationsByStream, "GET", url+q.Encode(), nil, token, params)
			err := json.Unmarshal([]byte(body), &b)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(b.Success).To(Equal(true))

			pretty, err := json.MarshalIndent(b, "", "  ")
			Expect(err).ShouldNot(HaveOccurred())
			fmt.Print(string(pretty))
		})

		It("Can Get Locations by stream per time", func() {
			var b Body

			q := make(url.Values)
			q.Set("startTime", "2017-07-03T00:00:00.000Z")
			q.Set("endTime", "2017-07-31T00:00:00.000Z")

			// define url
			url := "/analytics/logs/locations/H8SUumhjOR6PzKzm/time/?"
			// definimos la url con los parametros
			urlParams := "/analytics/logs/locations/:stream/time/?"
			// creamos un dictionary key - value, que contenga los valores
			values := map[string]string{
				"stream": "H8SUumhjOR6PzKzm",
			}
			// creamos un interface para pasar un solo parametro a la funcion ApiTest
			params := map[string]interface{}{
				"url":    urlParams,
				"values": values,
			}

			// ApiTest (func, method, url, body, token, urlParams)
			body := libs.ApiTest(api.GetLocationsByStreamPerTime, "GET", url+q.Encode(), nil, token, params)
			err := json.Unmarshal([]byte(body), &b)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(b.Success).To(Equal(true))

			pretty, err := json.MarshalIndent(b, "", "  ")
			Expect(err).ShouldNot(HaveOccurred())
			fmt.Print(string(pretty))
		})

		It("Can Get storage by account", func() {
			var b Body

			// ApiTest (func, method, url, body, token, urlParams)
			body := libs.ApiTest(api.GetStorage, "GET", "/analytics/logs/storage", nil, token, nil)
			err := json.Unmarshal([]byte(body), &b)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(b.Success).To(Equal(true))

			pretty, err := json.MarshalIndent(b, "", "  ")
			Expect(err).ShouldNot(HaveOccurred())
			fmt.Print(string(pretty))
		})

		It("Can Get storage by account per time", func() {
			var b Body

			q := make(url.Values)
			q.Set("startTime", "2017-07-03T00:00:00.000Z")
			q.Set("endTime", "2017-07-31T00:00:00.000Z")

			// ApiTest (func, method, url, body, token, urlParams)
			body := libs.ApiTest(api.GetStoragePerTime, "GET", "/analytics/logs/storage/time/?"+q.Encode(), nil, token, nil)
			err := json.Unmarshal([]byte(body), &b)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(b.Success).To(Equal(true))

			pretty, err := json.MarshalIndent(b, "", "  ")
			Expect(err).ShouldNot(HaveOccurred())
			fmt.Print(string(pretty))
		})

	})
})
