package pac

import (
	"log"

	"github.com/uptrace/bun"
)

func GetBunDB(s any) *bun.DB {
	db, ok := s.(*bun.DB)

	if !ok {
		log.Fatal("GetBunDB: cannot get Bun DB due to given parameter is not *bun.DB")
	}

	return db
}
