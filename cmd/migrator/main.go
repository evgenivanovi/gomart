package main

import (
	"flag"

	"github.com/evgenivanovi/gomart/pkg/goose"
	_ "github.com/lib/pq"
)

func main() {

	dir := flag.String(
		"dir",
		"./migrations",
		"directory with migration files",
	)

	dsn := flag.String(
		"dsn",
		"postgres://postgres:@localhost:5432/test?sslmode=disable",
		"postgres dsn",
	)

	command := flag.String(
		"command",
		"",
		"goose command",
	)

	driver := "postgres"

	flag.Parse()
	goose.Migrate(*dir, driver, *dsn, *command)

}
