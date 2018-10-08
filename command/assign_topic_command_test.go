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
		topicID uuid.UUID
	)

	Context("when i call NewAssignTopicCommand", func() {
		BeforeEach(func() {
			newsID = uuid.Must(uuid.NewV4())
			topicID = uuid.Must(uuid.NewV4())
			assignTopicCommand = NewAssignTopicCommand(newsID, topicID)
		})

		It("should fulfill Type", func() {
			Expect(assignTopicCommand.Type).To(Equal("AssignTopicCommand"))
		})

		It("should fulfill NewsID", func() {
			Expect(assignTopicCommand.NewsID).To(Equal(newsID))
		})

		It("should contains TopicIds", func() {
			Expect(assignTopicCommand.TopicID).To(Equal(topicID))
		})
	})
})
