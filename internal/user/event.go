package user

const (
	KafkaUserCreatedTopic = "user-created"
	KafkaUserUpdatedTopic = "user-updated"
	KafkaUserDeletedTopic = "user-deleted"
)

type Event struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
}
