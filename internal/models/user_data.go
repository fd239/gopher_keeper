package models

type DataType int32

const (
	TypeText DataType = 1 + iota
	TypeCard
	TypeFile
)

// UserData user data model
type UserData struct {
	Id       int
	Type     DataType
	EntityId int
	UserId   int
}

// DataText user data text type
type DataText struct {
	Id   int
	Text string
	Meta string
}

// DataCard user data card type
type DataCard struct {
	Id     int
	Number string
	Meta   string
}

// NewDataText returns a new user data text instance
func NewDataText(id int, text, meta string) *DataText {
	return &DataText{
		Id:   id,
		Text: text,
		Meta: meta,
	}
}

// NewDataCard returns a new user data card instance
func NewDataCard(id int, number, meta string) *DataCard {
	return &DataCard{
		Id:     id,
		Number: number,
		Meta:   meta,
	}
}
