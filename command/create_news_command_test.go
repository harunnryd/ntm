package command_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/satori/go.uuid"
	. "ntm/command"
)

var _ = Describe("CreateNewsCommand", func() {
	var (
		createNewsCommand *CreateNewsCommand
		topicId1 uuid.UUID
		topicId2 uuid.UUID
	)
	Context("when i call NewCreateNewsCommand", func() {
		BeforeEach(func() {
			topicId1 = uuid.Must(uuid.NewV4())
			topicId2 = uuid.Must(uuid.NewV4())
			createNewsCommand = NewCreateNewsCommand("how to maintain car", "i believe, i have car", topicId1, topicId2)
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

		It("should contains topicID", func() {
			Expect(createNewsCommand.TopicIds).To(ContainElement(topicId1))
		})
	})
})
