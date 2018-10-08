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
		topicID uuid.UUID
	)

	Context("when i call NewUnAssignTopicCommand", func() {
		BeforeEach(func() {
			newsID = uuid.Must(uuid.NewV4())
			topicID = uuid.Must(uuid.NewV4())
			unAssignTopicCommand = NewUnAssignTopicCommand(newsID, topicID)
		})

		It("should fulfill Type", func() {
			Expect(unAssignTopicCommand.Type).To(Equal("UnAssignTopicCommand"))
		})

		It("should fulfill NewsID", func() {
			Expect(unAssignTopicCommand.NewsID).To(Equal(newsID))
		})

		It("should contains TopicIds", func() {
			Expect(unAssignTopicCommand.TopicID).To(Equal(topicID))
		})
	})
})
