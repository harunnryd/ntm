package command_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "ntm/command"
)

var _ = Describe("CreateNewsCommand", func() {
	var createNewsCommand *CreateNewsCommand

	Context("when i call NewCreateNewsCommand", func() {
		BeforeEach(func() {
			createNewsCommand = NewCreateNewsCommand("how to maintain car", "i believe, i have car")
		})

		It("should fulfill ID", func() {
			Expect(createNewsCommand.ID).NotTo(BeNil())
		})

		It("should fulfill Type", func() {
			Expect(createNewsCommand.Type).To(Equal("CreateNewsCommand"))
		})

		It("should fulfill Title", func() {
			Expect(createNewsCommand.Title).To(Equal("how to maintain car"))
		})

		It("should fulfill Body", func() {
			Expect(createNewsCommand.Body).To(Equal("i believe, i have car"))
		})
	})
})
