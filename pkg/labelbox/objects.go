package labelbox

type ObjectEntity struct {
	FeatureID string `json:"featureId"`
	SchemaID  string `json:"schemaId"`
	Color     string `json:"color"`
	Title     string `json:"title"`
	Value     string `json:"value"`
	Version   int    `json:"version"`
	Format    string `json:"format"`
	Data      Data   `json:"data"`
}
type Location struct {
	Start int `json:"start"`
	End   int `json:"end"`
}
type Data struct {
	Location Location `json:"location"`
}
