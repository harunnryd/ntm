package commandApi_test

import (
	"github.com/labstack/echo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	. "ntm/commandApi"
	"strings"
)

var _ = Describe("UpdateStatusHandler", func() {
	var (
		h error
		e *echo.Echo
		c echo.Context
		req *http.Request
		rec *httptest.ResponseRecorder
	)

	Context("when i call NewUpdateStatusHandler", func() {
		BeforeEach(func() {
			e = echo.New()
			req = httptest.NewRequest(echo.PATCH, "/", strings.NewReader(`
				{
					"name": "draft"
				}`))
			req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec = httptest.NewRecorder()
			c = e.NewContext(req, rec)

			c.SetPath("/statuses/:status_id")
			c.SetParamNames("status_id")
			c.SetParamValues("a6ce6a11-7eea-4490-837b-5b6af31aebe1")

			h = NewUpdateStatusHandler(c)
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
