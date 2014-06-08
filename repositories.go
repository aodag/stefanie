package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/coopernurse/gorp"
)

func Open(dbUrl string) *sql.DB {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func InitDB(dbUrl string) *gorp.DbMap {
	db := Open(dbUrl)
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	dbmap.AddTableWithName(Project{}, "projects").SetKeys(true, "Id")

	err := dbmap.CreateTablesIfNotExists()
	if err != nil {
		panic(err.Error())
	}
	return dbmap
}

type GorpProjectRepository struct {
	dbmap *gorp.DbMap
}

func (repo *GorpProjectRepository) Add(project *Project) {
	repo.dbmap.Insert(project)
}

func NewGorpProjectRepository(dbmap *gorp.DbMap) *GorpProjectRepository {
	repo := new(GorpProjectRepository)
	repo.dbmap = dbmap
	return repo
}
