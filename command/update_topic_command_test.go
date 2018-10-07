package command_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/satori/go.uuid"
	. "ntm/command"
)

var _ = Describe("UpdateTopicCommand", func() {
	var (
		updateTopicCommand *UpdateTopicCommand
		topicID uuid.UUID
	)

	Context("when i call NewUpdateTopicCommand", func() {
		BeforeEach(func() {
			topicID = uuid.Must(uuid.NewV4())
			updateTopicCommand = NewUpdateTopicCommand(topicID, "poliklitik geli geli")
		})

		It("should fulfill Type", func() {
			Expect(updateTopicCommand.Type).To(Equal("UpdateTopicCommand"))
		})

		It("should fulfill NewsID", func() {
			Expect(updateTopicCommand.TopicID).To(Equal(topicID))
		})

		It("should fulfill Name", func() {
			Expect(updateTopicCommand.Name).To(Equal("poliklitik geli geli"))
		})
	})
})
