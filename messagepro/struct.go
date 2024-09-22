package messagepro

type (
	// MessageProSendRequest struct {
	// 	Key  string `json:"key"`
	// 	From string `json:"from"`
	// 	To   string `json:"to"`
	// 	Text string `json:"text"`
	// }
	MessageProSendResponse struct {
		Result    string `json:"Result"`
		MessageID int    `json:"Message ID"`
		Reason    string `json:"Reason,omitempty"`
	}
	// MessageProGetstatusRequest struct {
	// 	Key string `json:"key"`
	// 	ID  string `json:"id"`
	// }
	MessageProGetstatusResponse struct {
		Delivered         string `json:"delivered"`
		Result            string `json:"result"`
		SourceNumber      string `json:"source_number"`
		DestinationNumber string `json:"destination_number"`
		Text              string `json:"text"`
	}
	// MessageProFetchRequest struct {
	// 	From string `json:"from"`
	// 	To   string `json:"to"`
	// 	Text string `json:"text"`
	// }
)
