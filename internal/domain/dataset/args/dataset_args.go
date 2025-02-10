package args

// DatasetGenerateByDataArgs create dataset
type DatasetGenerateByDataArgs struct {

	// Name the name of the dataset. if the name is empty, the system will generate a name
	Name string

	// Data the user uploads data, the system will parse the data to generate fields
	Data []map[string]any
}
