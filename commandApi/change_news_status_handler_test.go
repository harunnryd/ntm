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

var _ = Describe("ChangeNewsStatusHandler", func() {
	var (
		h error
		e *echo.Echo
		c echo.Context
		req *http.Request
		rec *httptest.ResponseRecorder
	)

	Context("when i call NewChangeNewsStatusHandler", func() {
		BeforeEach(func() {
			e = echo.New()
			req = httptest.NewRequest(echo.PATCH, "/", strings.NewReader(`
				{
					"status_id": "bda8923e-9a29-4756-858e-978dfe01920b"
				}`))
			req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec = httptest.NewRecorder()
			c = e.NewContext(req, rec)
			c.SetPath("news/:news_id/statuses")
			c.SetParamNames("news_id")
			c.SetParamNames("28f4c7d2-c5d5-485f-ad32-88a4cba489ba")
			h = NewChangeNewsStatusHandler(c)
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
