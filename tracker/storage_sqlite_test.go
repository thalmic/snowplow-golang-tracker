// +build sqlite

package tracker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestStorageSQLite3Init asserts behaviour of SQLite storage functions.
func TestStorageSQLite3Init(t *testing.T) {
	assert := assert.New(t)
	storage := *InitStorageSQLite3("/home/vagrant/test.db")
	assert.NotNil(storage)
	assert.Equal("/home/vagrant/test.db", storage.DbName)

	defer func() {
		if err := recover(); err != nil {
			assert.NotNil(err)
		}
	}()
	storage = *InitStorageSQLite3("~/")
}

// TestSQLite3AddGetDeletePayload asserts ability to add, delete and get payloads.
func TestSQLite3AddGetDeletePayload(t *testing.T) {
	assert := assert.New(t)
	storage := *InitStorageSQLite3("/home/vagrant/test.db")
	assertDatabaseAddGetDeletePayload(assert, storage)
}

func TestSQLite3PanicRecovery(t *testing.T) {
	assert := assert.New(t)

	result := execDeleteQuery(nil, "")
	assert.Equal(int64(0), result)

	eventRows := execGetQuery(nil, "")
	assert.Equal(0, len(eventRows))

	addResult := execAddStatement(nil, nil)
	assert.False(addResult)
}
