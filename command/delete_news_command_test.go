package command_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/satori/go.uuid"
	. "ntm/command"
)

var _ = Describe("DeleteNewsCommand", func() {
	var (
		deleteNewsCommand *DeleteNewsCommand
		newsID uuid.UUID
	)

	Context("when i call NewDeleteNewsCommand", func() {
		BeforeEach(func() {
			newsID = uuid.Must(uuid.NewV4())
			deleteNewsCommand = NewDeleteNewsCommand(newsID)
		})

		It("should fulfill Type", func() {
			Expect(deleteNewsCommand.Type).To(Equal("DeleteNewsCommand"))
		})

		It("should fulfill NewsID", func() {
			Expect(deleteNewsCommand.NewsID).To(Equal(newsID))
		})
	})
})
