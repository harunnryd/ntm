package command_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/satori/go.uuid"
	. "ntm/command"
)

var _ = Describe("UnAssignTopicCommand", func() {
	var (
		unAssignTopicCommand *UnAssignTopicCommand
		newsID uuid.UUID
		topicId1 uuid.UUID
		topicId2 uuid.UUID
	)

	Context("when i call NewUnAssignTopicCommand", func() {
		BeforeEach(func() {
			newsID = uuid.Must(uuid.NewV4())
			topicId1 = uuid.Must(uuid.NewV4())
			topicId2 = uuid.Must(uuid.NewV4())
			unAssignTopicCommand = NewUnAssignTopicCommand(newsID, topicId1, topicId2)
		})

		It("should fulfill ID", func() {
			Expect(unAssignTopicCommand.ID).NotTo(BeNil())
		})

		It("should fulfill Type", func() {
			Expect(unAssignTopicCommand.Type).To(Equal("UnAssignTopicCommand"))
		})

		It("should fulfill NewsID", func() {
			Expect(unAssignTopicCommand.NewsID).To(Equal(newsID))
		})

		It("should contains TopicIds", func() {
			Expect(unAssignTopicCommand.TopicIds).To(ContainElement(topicId1))
		})
	})
})
