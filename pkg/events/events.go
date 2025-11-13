package events

// Name enumerates event identifiers emitted across services.
type Name string

const (
	EventUserSignedUp        Name = "user.signed_up"
	EventUserLoggedIn        Name = "user.logged_in"
	EventBlogPublished       Name = "blog.published"
	EventBlogScheduled       Name = "blog.scheduled"
	EventCommentCreated      Name = "comment.created"
	EventNotificationCreated Name = "notification.created"
	EventAnalyticsRecorded   Name = "analytics.recorded"
	EventBookmarkSaved       Name = "bookmark.saved"
)

// Envelope represents a lightweight async event payload.
type Envelope struct {
	Name    Name            `json:"name"`
	Source  string          `json:"source"`
	Payload map[string]any `json:"payload"`
}
