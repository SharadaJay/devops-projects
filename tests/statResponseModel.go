package tests

type StatResponse struct {
	OverallStats OverallStatResponse `json:"overall_statistics"`
	QueueStats   []QueueStatResponse `json:"queue_statistics"`
}

type OverallStatResponse struct {
	ClusterName               string  `json:"cluster_name"`
	TotalNumberOfQueues       int     `json:"total_number_of_queues"`
	MessagesDeliveredRecently int     `json:"messages_delivered_recently"`
	MessageDeliveryRate       float64 `json:"message_delivery_rate"`
	MessagesPublishedRecently int     `json:"messages_published_recently"`
	MessagePublishingRate     float64 `json:"message_publishing_rate"`
}

type QueueStatResponse struct {
	Name                      string  `json:"name"`
	MessagesDeliveredRecently int     `json:"messages_delivered_recently"`
	MessageDeliveryRate       float64 `json:"message_delivery_rate"`
	MessagesPublishedRecently int     `json:"messages_published_recently"`
	MessagePublishingRate     float64 `json:"message_publishing_rate"`
}