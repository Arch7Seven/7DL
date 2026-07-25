package sqlite

import "7DL/sqlite"

type SQLite struct {
	SQLiteDB *sqlite.SQLiteDB
}

func New(sqliteDb *sqlite.SQLiteDB) *SQLite {
	return &SQLite{
		SQLiteDB: sqliteDb,
	}
}

func (sqlite *SQLite) SaveCredentials() {

}

func (sqlite *SQLite) RemoveCredentials() {

}

func (sqlite *SQLite) UpdateCredentials() {

}
