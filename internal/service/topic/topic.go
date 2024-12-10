package topic

type TopicService interface {
	Subscribe(topic string) error
}
