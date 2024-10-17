package enum

type Enum map[int]string

// Code 查询Enum的Code（Key）
func (e Enum) Code(value string) int {
	for k, v := range e {
		if v == value {
			return k
		}
	}
	return -1
}

// Value 查询Enum的Value（Key对应的Value）
func (e Enum) Value(code int) string {
	value, _ := e[code]
	return value
}
