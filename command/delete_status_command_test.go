package command_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/satori/go.uuid"
	. "ntm/command"
)

var _ = Describe("DeleteStatusCommand", func() {
	var (
		deleteStatusCommand *DeleteStatusCommand
		statusID uuid.UUID
	)

	Context("when i call NewDeleteStatusCommand", func() {
		BeforeEach(func() {
			statusID = uuid.Must(uuid.NewV4())
			deleteStatusCommand = NewDeleteStatusCommand(statusID)
		})

		It("should fulfill Type", func() {
			Expect(deleteStatusCommand.Type).To(Equal("DeleteStatusCommand"))
		})

		It("should fulfill NewsID", func() {
			Expect(deleteStatusCommand.StatusID).To(Equal(statusID))
		})
	})
})
