package sbarternetwork

// CallStack holds the callstack information together with the previous callstacks
// Ideal for debugging purposes to follow the requests as they flow through the distributed architecture.
type CallStack struct {
	SequenceID    int        `json:"sequenceId"`
	Direction     int        `json:"direction"`
	CorrelationID string     `json:"correlationId"`
	Timestamp     int64      `json:"timestamp"`
	Project       string     `json:"project"`
	Function      string     `json:"function"`
	Path          string     `json:"path"`
	Line          int        `json:"line"`
	Previous      *CallStack `json:"previous,omitempty"`
}
