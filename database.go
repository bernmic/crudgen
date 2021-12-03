package crudgen

type DatabaseDefinition interface {
	GenerateCRUDFilesForTable(tablename string) (string, string, error)
}
