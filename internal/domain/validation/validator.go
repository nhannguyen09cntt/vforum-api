package validation

var ValidationStructs = make(map[string]interface{})

type Validator interface {
	Validate(params map[string]interface{}, context string) error
}
