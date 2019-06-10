package validator

import (
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	emptyStr          = ""
	emailPatern       = regexp.MustCompile(".+@.+\\..+")
	dateiso8601Patern = regexp.MustCompile("^(\\d{4})-(\\d{2})-(\\d{2})T(\\d{2}):(\\d{2}):(\\d{2})(Z|(\\+|-)\\d{2}(:?\\d{2})?)$")
)

func NotNil(val interface{}) bool {
	if val == nil {
		return false
	}

	return true
}

func Email(val string) bool {
	if val == "" {
		return false
	}

	if !emailPatern.MatchString(val) {
		return false
	}

	return true
}

func Gender(val string) bool {
	if val != `male` && val != `female` {
		return false
	}

	return true
}

func InString(val string, in []string) bool {
	c := 0
	for _, k := range in {
		if k == val {
			c++
		}
	}
	if c == len(in) {
		return true
	}

	return false
}

func RequiredString(val string) bool {
	if len(strings.TrimSpace(val)) == 0 {
		return false
	}

	return true
}

func RequiredBytes(val []byte) bool {
	if len(val) == 0 {
		return false
	}

	return true
}

func RequiredInt(val int) bool {
	if val == 0 {
		return false
	}

	return true
}

func RequiredInt32(val int32) bool {
	if val == 0 {
		return false
	}

	return true
}

func RequiredInt64(val int64) bool {
	if val == 0 {
		return false
	}

	return true
}

func RequiredFloat32(val float32) bool {
	if val == 0 {
		return false
	}

	return true
}

func RequiredFloat64(val float64) bool {
	if val == 0 {
		return false
	}

	return true
}

func RequiredBool(val bool) bool {
	if !val {
		return false
	}

	return true
}

func RequiredEmail(val string) bool {
	if val == "" {
		return false
	}

	return Email(val)
}

func RequiredTime(val time.Time) bool {
	if val.IsZero() {
		return false
	}

	return true
}

func RequiredArrayString(val []string) bool {
	if len(val) == 0 {
		return false
	}

	return true
}

func MinChar(val string, n int) bool {
	if utf8.RuneCountInString(val) < n {
		return false
	}

	return true
}

func MinInt(val int, minVal int) bool {
	if val > minVal {
		return true
	}

	return false
}

func MinInt32(val int32, minVal int32) bool {
	if val > minVal {
		return true
	}

	return false
}

func MinInt64(val int64, minVal int64) bool {
	if val > minVal {
		return true
	}

	return false
}

func MinFloat32(val float32, minVal float32) bool {
	if val < minVal {
		return true
	}

	return false
}

func MinFloat64(val float64, minVal float64) bool {
	if val < minVal {
		return true
	}

	return false
}

func MaxChar(val string, max int) bool {
	if utf8.RuneCountInString(val) > max {
		return false
	}

	return true
}

func MaxInt(val int, max int) bool {
	if val > max {
		return false
	}

	return true
}

func MaxInt32(val int32, max int32) bool {
	if val > max {
		return false
	}

	return true
}

func MaxInt64(val int64, max int64) bool {
	if val > max {
		return false
	}

	return true
}

func MaxFloat32(val float32, n float32) bool {
	if val > n {
		return true
	}

	return false
}

func MaxFloat64(val float64, n float64) bool {
	if val > n {
		return true
	}

	return false
}

func RangeInt(val, min, max int) bool {
	if val >= min && val <= max {
		return true
	}

	return false
}

func RangeInt32(val, min, max int32) bool {
	if val >= min && val <= max {
		return true
	}

	return false
}

func RangeInt64(val, min, max int64) bool {
	if val >= min && val <= max {
		return true
	}

	return false
}

func RangeFloat32(val, min, max float32) bool {
	if val >= min && val <= max {
		return true
	}

	return false
}

func RangeFloat64(val, min, max float64) bool {
	if val >= min && val <= max {
		return true
	}

	return false
}

func LenArrayString(val []string, nLen int) bool {
	if len(val) == nLen {
		return true
	}

	return false
}

func LenArrayInt(val []int, nLen int) bool {
	if len(val) == nLen {
		return true
	}

	return false
}

func LenArrayInt32(val []int32, nLen int) bool {
	if len(val) == nLen {
		return true
	}

	return false
}

func LenArrayInt64(val []int64, nLen int) bool {
	if len(val) == nLen {
		return true
	}

	return false
}

func LenArrayFloat32(val []float32, nLen int) bool {
	if len(val) == nLen {
		return true
	}

	return false
}

func LenArrayFloat64(val []float64, nLen int) bool {
	if len(val) == nLen {
		return true
	}

	return false
}

func PointerRequiredString(val *string) bool {
	if val == nil {
		return false
	}

	return RequiredString(*val)
}

func PointerRequiredEmail(val *string) bool {
	if val == nil {
		return false
	}

	return RequiredString(*val)
}

func PointerRequiredInt(val *int) bool {
	if val == nil {
		return false
	}

	return RequiredInt(*val)
}

func PointerRequiredInt32(val *int32) bool {
	if val == nil {
		return false
	}

	return RequiredInt32(*val)
}

func PointerRequiredInt64(val *int64) bool {
	if val == nil {
		return false
	}

	return RequiredInt64(*val)
}

func PointerRequiredFloat64(val *float64) bool {
	if val == nil {
		return false
	}

	return RequiredFloat64(*val)
}

func PointerRequiredFloat32(val *float32) bool {
	if val == nil {
		return false
	}

	return RequiredFloat32(*val)
}

func PointerRequiredBool(val *bool) bool {
	if val == nil {
		return false
	}

	return RequiredBool(*val)

}

func PointerRangeInt(val *int, min, max int) bool {
	if val == nil || (*val >= min && *val <= max) {
		return true
	}

	return false
}

func PointerRangeInt32(val *int32, min, max int32) bool {
	if val == nil || *val >= min && *val <= max {
		return true
	}

	return false
}

func PointerRangeInt64(val *int64, min, max int64) bool {
	if val == nil || *val >= min && *val <= max {
		return true
	}

	return false
}

func PointerRangeFloat32(val *float32, min, max float32) bool {
	if val == nil || *val >= min && *val <= max {
		return true
	}

	return false
}

func PointerRangeFloat64(val *float64, min, max float64) bool {
	if val == nil || *val >= min && *val <= max {
		return true
	}

	return false
}
