package data

import (
	"database/sql"
	"fmt"
	upperdb "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
	"github.com/upper/db/v4/adapter/sqlite"
	"os"
	"strings"
)

var db *sql.DB
var upperDbSess upperdb.Session

//Models is the wrapper for all database models
type Models struct {
	//any models inserted here (and in the New function)
	//are easily accessible throughout the entire application
	//uncomment the lines below after running make auth commands
	//to add auth models to package
	//Users  User
	//Tokens Token

}

func New(databasePool *sql.DB) Models {
	db = databasePool

	switch strings.ToLower(os.Getenv("DATABASE_TYPE")) {
	case "mysql", "mariadb":
		upperDbSess, _ = mysql.New(databasePool)
	case "postgres", "postgresql":
		upperDbSess, _ = postgresql.New(databasePool)
	case "sqlite":
		upperDbSess, _ = sqlite.New(databasePool)
	default:

	}

	return Models{}
}

func getIdFromInsert(i upperdb.ID) int {
	idType := fmt.Sprintf("%T", i)

	if idType == "int64" {
		return int(i.(int64))
	}

	return i.(int)
}
