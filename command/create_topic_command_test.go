package command_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "ntm/command"
)

var _ = Describe("CreateTopicCommand", func() {
	var createTopicCommand *CreateTopicCommand

	Context("when i call NewCreateTopicCommand", func() {
		BeforeEach(func() {
			createTopicCommand = NewCreateTopicCommand("car")
		})

		It("should fulfill ID", func() {
			Expect(createTopicCommand.ID).NotTo(BeNil())
		})

		It("should fulfill Type", func() {
			Expect(createTopicCommand.Type).To(Equal("CreateTopicCommand"))
		})

		It("should fulfill Title", func() {
			Expect(createTopicCommand.Name).To(Equal("car"))
		})
	})
})
