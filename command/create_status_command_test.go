package command_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "ntm/command"
)

var _ = Describe("CreateStatusCommand", func() {
	var createStatusCommand *CreateStatusCommand

	Context("when i call NewCreateStatusCommand", func() {
		BeforeEach(func() {
			createStatusCommand = NewCreateStatusCommand("publish")
		})

		It("should fulfill ID", func() {
			Expect(createStatusCommand.ID).NotTo(BeNil())
		})

		It("should fulfill Type", func() {
			Expect(createStatusCommand.Type).To(Equal("CreateStatusCommand"))
		})

		It("should fulfill Title", func() {
			Expect(createStatusCommand.Name).To(Equal("publish"))
		})
	})
})
