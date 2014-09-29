package model

type ServiceAnnouncement struct {
	Topic            string    `json:"topic" redis:"topic"`
	Schema           string    `json:"schema" redis:"schema"`
	Deprecated       *bool     `json:"deprecated,omitempty"`
	SupportedMethods *[]string `json:"methods" redis:"methods,json"`
	SupportedEvents  *[]string `json:"events" redis:"events,json"`
}
