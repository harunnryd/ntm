package command_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/satori/go.uuid"
	. "ntm/command"
)

var _ = Describe("DeleteTopicCommand", func() {
	var (
		deleteTopicCommand *DeleteTopicCommand
		topicID uuid.UUID
	)

	Context("when i call NewDeleteTopicCommand", func() {
		BeforeEach(func() {
			topicID = uuid.Must(uuid.NewV4())
			deleteTopicCommand = NewDeleteTopicCommand(topicID)
		})

		It("should fulfill Type", func() {
			Expect(deleteTopicCommand.Type).To(Equal("DeleteTopicCommand"))
		})

		It("should fulfill NewsID", func() {
			Expect(deleteTopicCommand.TopicID).To(Equal(topicID))
		})
	})
})
