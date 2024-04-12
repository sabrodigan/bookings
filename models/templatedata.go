package models

// TemplateData this is important for the 'boxes' app
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} // if you are unsure on the data types use interface
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
