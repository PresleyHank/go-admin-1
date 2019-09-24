// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sqlite

import (
	"database/sql"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/db/performer"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

type Sqlite struct {
	DbList map[string]*sql.DB
	Once   sync.Once
}

func GetSqliteDB() *Sqlite {
	return &DB
}

func (db *Sqlite) GetName() string {
	return "sqlite"
}

func (db *Sqlite) QueryWithConnection(con string, query string, args ...interface{}) ([]map[string]interface{}, *sql.Rows) {
	return performer.Query(db.DbList[con], query, args...)
}

func (db *Sqlite) ExecWithConnection(con string, query string, args ...interface{}) sql.Result {
	return performer.Exec(db.DbList[con], query, args...)
}

func (db *Sqlite) Query(query string, args ...interface{}) ([]map[string]interface{}, *sql.Rows) {
	return performer.Query(db.DbList["default"], query, args...)
}

func (db *Sqlite) Exec(query string, args ...interface{}) sql.Result {
	return performer.Exec(db.DbList["default"], query, args...)
}

func (db *Sqlite) InitDB(cfgList map[string]config.Database) {
	db.Once.Do(func() {
		var (
			sqlDB *sql.DB
			err   error
		)

		for conn, cfg := range cfgList {
			sqlDB, err = sql.Open("sqlite3", cfg.File)

			if err != nil {
				panic(err)
			} else {
				db.DbList[conn] = sqlDB
			}
		}
	})
}

var DB = Sqlite{
	DbList: map[string]*sql.DB{},
}
