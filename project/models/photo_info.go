package models

import "github.com/vmihailenco/msgpack"

type Photo struct {
	PhotoID       int64  `json:"photo_id" db:"photo_id" msgpack:"a"`
	PhotoTitle    string `json:"photo_title" db:"photo_title" msgpack:"b"`
	PhotoUrl      string `json:"photo_url" db:"photo_url" msgpack:"c"`
	PhotoShotDate string `json:"photo_shot_date" db:"photo_shot_date" msgpack:"d"`
	Describe      string `json:"describe" db:"describe" msgpack:"e"`
	CreateTime    string `json:"create_time" db:"create_time" msgpack:"f"`
	UpdateTime    string `json:"update_time" db:"update_time" msgpack:"g"`
}

func (p *Photo) MarshalBinary() ([]byte, error) {
	return msgpack.Marshal(p)
}

// UnmarshalBinary use msgpack
func (p *Photo) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, p)
}
