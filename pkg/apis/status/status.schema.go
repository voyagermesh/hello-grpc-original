package status

// Auto-generated. DO NOT EDIT.
import (
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var statusRequestSchema *gojsonschema.Schema

func init() {
	var err error
	statusRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *StatusRequest) Valid() (*gojsonschema.Result, error) {
	return statusRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *StatusRequest) IsRequest() {}
