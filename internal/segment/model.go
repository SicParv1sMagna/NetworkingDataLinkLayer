package segment

import "time"

type Segment struct {
	ID            time.Time `json:"id" example:"2024-03-09T12:04:08Z"`
	TotalSegments uint      `json:"total_segments" example:"10"`
	SenderName    string    `json:"sender_name" example:"Марк Гревцов"`
	SegmentNumber uint      `json:"segment_number" example:"1"`
	HadError      bool      `json:"had_error" example:"false"`
	Payload       []byte    `json:"payload"`
}

type SegmentRequest struct {
	ID            time.Time `json:"id" example:"2024-03-09T12:04:08Z"`
	TotalSegments uint      `json:"total_segments" example:"10"`
	SenderName    string    `json:"sender_name" example:"Марк Гревцов"`
	SegmentNumber uint      `json:"segment_number" example:"1"`
	Payload       []byte    `json:"payload"`
}
