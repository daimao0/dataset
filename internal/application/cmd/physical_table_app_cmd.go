package cmd

// PhysicalTableCreateByDataCmd is command for create data source by data
type PhysicalTableCreateByDataCmd struct {

	// Name the name of the data source. if the name is empty, the system will generate a name
	Name string

	// Data from the data source
	Data []map[string]any
}

type PhysicalTableCreateBySQLCmd struct {

	// Name the name of the data source. if the name is empty, the system will generate a name
	Name string

	// SQL the sql of the data source
	SQL string

	// DatabaseType is the Database Type
	DatabaseType string
}
