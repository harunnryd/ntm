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

var _ = Describe("CreateStatusHandler", func() {
	var (
		h error
		e *echo.Echo
		c echo.Context
		req *http.Request
		rec *httptest.ResponseRecorder
	)

	Context("when i call NewCreateStatusHandler", func() {
		BeforeEach(func() {
			e = echo.New()
			req = httptest.NewRequest(echo.POST, "/statuses", strings.NewReader(`
				{
					"name": "draft"
				}`))
			req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec = httptest.NewRecorder()
			c = e.NewContext(req, rec)
			h = NewCreateStatusHandler(c)
		})

		It("should no error", func() {
			Expect(h).To(BeNil())
		})

		It("should response http status created", func() {
			Expect(rec.Code).To(Equal(http.StatusCreated))
		})

		It("should response json", func() {
			Expect(rec.Body.String()).To(ContainSubstring(`"message":"success!"`))
		})

	})
})
