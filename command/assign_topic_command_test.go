package command_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/satori/go.uuid"
	. "ntm/command"
)

var _ = Describe("AssignTopicCommand", func() {
	var (
		assignTopicCommand *AssignTopicCommand
		newsID uuid.UUID
		topicId1 uuid.UUID
		topicId2 uuid.UUID
	)

	Context("when i call NewAssignTopicCommand", func() {
		BeforeEach(func() {
			newsID = uuid.Must(uuid.NewV4())
			topicId1 = uuid.Must(uuid.NewV4())
			topicId2 = uuid.Must(uuid.NewV4())
			assignTopicCommand = NewAssignTopicCommand(newsID, topicId1, topicId2)
		})

		It("should fulfill ID", func() {
			Expect(assignTopicCommand.ID).NotTo(BeNil())
		})

		It("should fulfill Type", func() {
			Expect(assignTopicCommand.Type).To(Equal("AssignTopicCommand"))
		})

		It("should fulfill NewsID", func() {
			Expect(assignTopicCommand.NewsID).To(Equal(newsID))
		})

		It("should contains TopicIds", func() {
			Expect(assignTopicCommand.TopicIds).To(ContainElement(topicId1))
		})
	})
})
