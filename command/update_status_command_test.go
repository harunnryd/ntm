package command_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/satori/go.uuid"
	. "ntm/command"
)

var _ = Describe("UpdateStatusCommand", func() {
	var (
		updateStatusCommand *UpdateStatusCommand
		statusID uuid.UUID
	)

	Context("when i call NewUpdateStatusCommand", func() {
		BeforeEach(func() {
			statusID = uuid.Must(uuid.NewV4())
			updateStatusCommand = NewUpdateStatusCommand(statusID, "draft")
		})

		It("should fulfill Type", func() {
			Expect(updateStatusCommand.Type).To(Equal("UpdateStatusCommand"))
		})

		It("should fulfill NewsID", func() {
			Expect(updateStatusCommand.StatusID).To(Equal(statusID))
		})

		It("should fulfill Name", func() {
			Expect(updateStatusCommand.Name).To(Equal("draft"))
		})
	})
})
