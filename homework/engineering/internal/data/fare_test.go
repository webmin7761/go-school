package data

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/webmin7761/go-school/homework/engineering/internal/data/ent/enttest"
)

func TestXXX(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
}
