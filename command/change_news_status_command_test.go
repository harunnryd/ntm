package command_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/satori/go.uuid"
	. "ntm/command"
)

var _ = Describe("ChangeNewsStatusCommand", func() {
	var (
		changeNewsStatusCommand *ChangeNewsStatusCommand
		newsID uuid.UUID
		statusID uuid.UUID
	)

	Context("when i call NewChangeNewsStatusCommand", func() {
		BeforeEach(func() {
			newsID = uuid.Must(uuid.NewV4())
			statusID = uuid.Must(uuid.NewV4())
			changeNewsStatusCommand = NewChangeNewsStatusCommand(newsID, statusID)
		})

		It("should fulfill ID", func() {
			Expect(changeNewsStatusCommand.ID).NotTo(BeNil())
		})

		It("should fulfill Type", func() {
			Expect(changeNewsStatusCommand.Type).To(Equal("ChangeNewsStatusCommand"))
		})

		It("should fulfill NewsID", func() {
			Expect(changeNewsStatusCommand.NewsID).To(Equal(newsID))
		})

		It("should fulfill StatusID", func() {
			Expect(changeNewsStatusCommand.StatusID).To(Equal(statusID))
		})
	})
})
