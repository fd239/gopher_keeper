package models

import (
	"github.com/fd239/gopher_keeper/pkg/pb"
	uuid "github.com/satori/go.uuid"
)

type DataType int32

const (
	TypeText DataType = 1 + iota
	TypeCard
	TypeFile
)

// DataText user data text type
type DataText struct {
	Id     int
	Type   DataType
	UserId string
	Text   string
	Meta   string
}

func (c *DataText) ToProto() *pb.DataText {
	return &pb.DataText{
		Text: c.Text,
		Meta: c.Meta,
	}
}

// DataCard user data card type
type DataCard struct {
	Id     int
	Type   DataType
	UserId string
	Number string
	Meta   string
}

func (c *DataCard) ToProto() *pb.DataCard {
	return &pb.DataCard{
		Number: c.Number,
		Meta:   c.Meta,
	}
}

type DataFile struct {
	FileId   uuid.UUID
	Type     DataType
	UserId   string
	FileType string
	Path     string
}

// NewDataText returns a new user data text instance
func NewDataText(text, meta string) *DataText {
	return &DataText{
		Type: TypeText,
		Text: text,
		Meta: meta,
	}
}

// NewDataCard returns a new user data card instance
func NewDataCard(number, meta string) *DataCard {
	return &DataCard{
		Type:   TypeCard,
		Number: number,
		Meta:   meta,
	}
}
