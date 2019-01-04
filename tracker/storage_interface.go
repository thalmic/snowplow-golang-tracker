package tracker

const (
	DB_DRIVER       = "sqlite3"
	DB_TABLE_NAME   = "events"
	DB_COLUMN_ID    = "id"
	DB_COLUMN_EVENT = "event"
)

type Storage interface {
	AddEventRow(payload Payload) bool
	DeleteAllEventRows() int64
	DeleteEventRows(ids []int) int64
	GetAllEventRows() []EventRow
	GetEventRowsWithinRange(eventRange int) []EventRow
}

type RawEventRow struct {
	id    int
	event []byte
}

type RawEventRowUint struct {
	id    uint
	event []byte
}

type EventRow struct {
	id    int
	event Payload
}
