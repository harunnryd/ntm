package commandApi_test

import (
	"github.com/labstack/echo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	. "ntm/commandApi"
)

var _ = Describe("DeleteTopicHandler", func() {
	var (
		h error
		e *echo.Echo
		c echo.Context
		req *http.Request
		rec *httptest.ResponseRecorder
	)

	Context("when i call NewDeleteTopicHandler", func() {
		BeforeEach(func() {
			e = echo.New()
			req = httptest.NewRequest(echo.DELETE, "/", nil)
			req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec = httptest.NewRecorder()
			c = e.NewContext(req, rec)

			c.SetPath("/topics/:topic_id")
			c.SetParamNames("topic_id")
			c.SetParamValues("a6ce6a11-7eea-4490-837b-5b6af31aebe1")

			h = NewDeleteTopicHandler(c)
		})

		It("should no error", func() {
			Expect(h).To(BeNil())
		})

		It("should response http status accepted", func() {
			Expect(rec.Code).To(Equal(http.StatusAccepted))
		})

		It("should response json", func() {
			Expect(rec.Body.String()).To(ContainSubstring(`"message":"success!"`))
		})

	})
})
