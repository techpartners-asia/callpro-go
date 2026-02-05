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

	OrderCampaignRequest struct {
		Name        string      `json:"name"`
		IsWithText  int         `json:"isWithText"` // 0 эсвэл 1 /0 үед олон дугаар луу нэг текст илгээнэ, 1 үед нэг текстийг нэг дугаар луу илгээнэ/
		Text        string      `json:"text"`
		From        string      `json:"from"`
		BeginDate   string      `json:"begin_date"`
		BeginHour   string      `json:"begin_hour"`
		BeginMinute string      `json:"begin_minute"`
		Numbers     interface{} `json:"numbers"`
	}
)
