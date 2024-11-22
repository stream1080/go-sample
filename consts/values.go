package consts

type Mapping struct {
	Value   any    `json:"value"`
	Display string `json:"display"`
}

type Field struct {
	Field    string    `json:"field"`
	Display  string    `json:"display"`
	Mappings []Mapping `json:"mappings"`
}

var Values = map[string]Field{
	"status": {
		Field:   "status",
		Display: "状态",
		Mappings: []Mapping{
			{Value: "enable", Display: "启用"},
			{Value: "disable", Display: "禁用"},
		},
	},
}
