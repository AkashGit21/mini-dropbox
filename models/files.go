package models

import "time"

const (
	STATUS_INACTIVE = iota
	STATUS_ACTIVE
)

var (
	FileStatus_name = map[int8]string{
		STATUS_INACTIVE: "inactive",
		STATUS_ACTIVE:   "active",
	}

	FileStatus_value = map[string]int8{
		"active":   STATUS_ACTIVE,
		"inactive": STATUS_INACTIVE,
	}
)

type FileStatus int8

func (x FileStatus) String() string {
	if val, ok := FileStatus_name[int8(x)]; ok {
		return val
	}
	return "inactive"
}

type Metadata struct {
	ID          int64      `json:"id,omitempty"`
	Filename    string     `json:"filename"`
	SizeInBytes int64      `json:"size_in_bytes"`
	S3ObjectKey string     `json:"s3_object_key"`
	Description string     `json:"description,omitempty"`
	MimeType    string     `json:"mime_type,omitempty"`
	Status      FileStatus `json:"status"`
	PrevKey     string     `json:"prev_key,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
