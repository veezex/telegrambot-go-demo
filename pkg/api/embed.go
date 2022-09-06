package api

import _ "embed"

//go:embed v1/openapiv2/api.swagger.json
var SwaggerJson []byte
