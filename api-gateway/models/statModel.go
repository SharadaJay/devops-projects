package models

type OverallStat struct {
	ClusterName  string       `json:"cluster_name"`
	MessageStats MessageStats `json:"message_stats"`
	QueueTotals  QueueTotals  `json:"queue_totals"`
	ObjectTotals ObjectTotals `json:"object_totals"`
}

type QueueStat struct {
	MessageStats MessageStats `json:"message_stats"`
	Name         string       `json:"name"`
}

type MessageStats struct {
	DeliverGet        int         `json:"deliver_get"`
	DeliverGetDetails RateDetails `json:"deliver_get_details"`
	Publish           int         `json:"publish"`
	PublishDetails    RateDetails `json:"publish_details"`
}

type RateDetails struct {
	Rate float64 `json:"rate"`
}

type QueueTotals struct {
	Messages             int         `json:"messages"`
	MessagesDetails      RateDetails `json:"messages_details"`
	MessagesReady        int         `json:"messages_ready"`
	MessagesReadyDetails RateDetails `json:"messages_ready_details"`
}

type ObjectTotals struct {
	Channels    int `json:"channels"`
	Connections int `json:"connections"`
	Consumers   int `json:"consumers"`
	Exchanges   int `json:"exchanges"`
	Queues      int `json:"queues"`
}
