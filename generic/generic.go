package generic

type Type interface{}

// Number var GenericType generic.Number
type Number float64

func ToNumber(value interface{}) Number {
	switch v := value.(type) {

	case Number:
		return v
	case int:
		return Number(v)
	case int8:
		return Number(v)
	case int16:
		return Number(v)
	case int32:
		return Number(v)
	case int64:
		return Number(v)
	case uint:
		return Number(v)
	case uint8:
		return Number(v)
	case uint16:
		return Number(v)
	case uint32:
		return Number(v)
	case uint64:
		return Number(v)
	case float32:
		return Number(v)
	case float64:
		return Number(v)
	default:
		panic("cannot convert value to Number(float64)")
	}
}
