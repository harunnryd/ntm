package command_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/satori/go.uuid"
	. "ntm/command"
)

var _ = Describe("UpdateNewsCommand", func() {
	var (
		updateNewsCommand *UpdateNewsCommand
		newsID uuid.UUID
	)

	Context("when i call NewUpdateNewsCommand", func() {
		BeforeEach(func() {
			newsID = uuid.Must(uuid.NewV4())
			updateNewsCommand = NewUpdateNewsCommand(newsID, "car cir cur", "is body car cir cur")
		})

		It("should fulfill Type", func() {
			Expect(updateNewsCommand.Type).To(Equal("UpdateNewsCommand"))
		})

		It("should fulfill NewsID", func() {
			Expect(updateNewsCommand.NewsID).To(Equal(newsID))
		})

		It("should fulfill Title", func() {
			Expect(updateNewsCommand.Title).To(Equal("car cir cur"))
		})

		It("should fulfill Body", func() {
			Expect(updateNewsCommand.Body).To(Equal("is body car cir cur"))
		})
	})
})
