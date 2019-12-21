package parser

// Parser :
type Parser interface {
	ParseQuery(query map[string]interface{}, assign interface{}) error
}
