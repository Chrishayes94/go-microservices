package base

type Response struct {
	TrackingNumber			string	`json:"TrackingNumber"`
	PieceTrackingNumbers 	[]string `json:"pieceTrackingNumbers"`
	LabelBytes				[]byte	`json:"labelBytes"`
	LabelHtml				string	`json:"labelHtml"`
}
