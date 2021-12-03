package main

import (
	"flag"
	"fmt"
)

const (
	defaultPackageName      = "main"
	usagePackageName        = "package name for the generated files"
	defaultDatabaseType     = "mysql"
	usageDatabaseType       = "database type (mysql, postgres)"
	defaultDatabaseHost     = "localhost"
	usageDatabaseHost       = "database hostname"
	defaultDatabasePort     = "3306"
	usageDatabasePort       = "database port"
	defaultDatabaseName     = ""
	usageDatabaseName       = "database name"
	defaultDatabaseSchema   = ""
	usageDatabaseSchema     = "database schema"
	defaultDatabaseUser     = "root"
	usageDatabaseUser       = "database user"
	defaultDatabasePassword = ""
	usageDatabasePassword   = "database password"
	defaultDatabaseTable    = ""
	usageDatabaseTable      = "database table"
)

var (
	version string = "0.0.0"
)

type Parameters struct {
	packageName      string
	databaseType     string
	databaseHost     string
	databasePort     string
	databaseName     string
	databaseSchema   string
	databaseUser     string
	databasePassword string
	databaseTable    string
}

func main() {
	parameters := parseParameters()
	fmt.Println("crudgen " + version)
	fmt.Printf("Generating CRUD for %s in package %s\n", parameters.databaseType, parameters.packageName)
}

func parseParameters() Parameters {
	var parameters Parameters
	flag.StringVar(&parameters.packageName, "package", defaultPackageName, usagePackageName)
	flag.StringVar(&parameters.packageName, "P", defaultPackageName, usagePackageName+" (shorthand)")
	flag.StringVar(&parameters.databaseType, "database", defaultDatabaseType, usageDatabaseType)
	flag.StringVar(&parameters.databaseType, "d", defaultDatabaseType, usageDatabaseType+" (shorthand)")
	flag.StringVar(&parameters.databaseHost, "host", defaultDatabaseHost, usageDatabaseHost)
	flag.StringVar(&parameters.databaseHost, "h", defaultDatabaseHost, usageDatabaseHost+" (shorthand)")
	flag.StringVar(&parameters.databasePort, "port", defaultDatabasePort, usageDatabasePort)
	flag.StringVar(&parameters.databasePort, "p", defaultDatabasePort, usageDatabasePort+" (shorthand)")
	flag.StringVar(&parameters.databaseSchema, "schema", defaultDatabaseSchema, usageDatabaseSchema)
	flag.StringVar(&parameters.databaseSchema, "s", defaultDatabaseSchema, usageDatabaseSchema+" (shorthand)")
	flag.StringVar(&parameters.databaseName, "name", defaultDatabaseName, usageDatabaseName)
	flag.StringVar(&parameters.databaseName, "n", defaultDatabaseSchema, usageDatabaseName+" (shorthand)")
	flag.StringVar(&parameters.databaseUser, "user", defaultDatabaseUser, usageDatabaseUser)
	flag.StringVar(&parameters.databaseUser, "u", defaultDatabaseUser, usageDatabaseUser+" (shorthand)")
	flag.StringVar(&parameters.databasePassword, "password", defaultDatabasePassword, usageDatabasePassword)
	flag.StringVar(&parameters.databasePassword, "pw", defaultDatabasePassword, usageDatabasePassword+" (shorthand)")
	flag.StringVar(&parameters.databaseTable, "table", defaultDatabaseTable, usageDatabaseTable)
	flag.StringVar(&parameters.databaseTable, "t", defaultDatabaseTable, usageDatabaseTable+" (shorthand)")
	flag.Parse()

	return parameters
}
