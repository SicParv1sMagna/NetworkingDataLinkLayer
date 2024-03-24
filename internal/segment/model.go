package segment

import "time"

type Segment struct {
	ID            time.Time `json:"id" example:"2024-03-09T12:04:08Z"`
	TotalSegments uint      `json:"total_segments" example:"10"`
	SenderName    string    `json:"sender_name" example:"Марк Гревцов"`
	SegmentNumber uint      `json:"segment_number" example:"1"`
	Payload []byte `json:"payload" example:"116,104,105,115,32,105,115,32,97,32,98,121,116,101,32,97,114,114,97,121,32,111,102,32,98,121,116,101,115"`
}
