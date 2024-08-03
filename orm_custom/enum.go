package orm_custom

type Enum struct {
	Name  string `boil:"enum_name" json:"enumName" toml:"enumName" yaml:"enumName"`
	Value string `boil:"enum_value" json:"enumValue" toml:"enumValue" yaml:"enumValue"`
}
