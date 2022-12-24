package util

import (
	"fmt"
	"math"
	"strconv"
)

const (
	// If the absolute value of the floating-point number is less than this number, it can be regarded as 0
	floatZero = 0.000001
)

// InterfaceToInt64 Convert interface{} to int64
func InterfaceToInt64(n interface{}) int64 {
	var r int64
	var err error
	switch v := n.(type) {
	case int:
		r = int64(int(v))
	case int64:
		r = int64(v)
	case float32:
		r = int64(float32(v))
	case float64:
		r = int64(float64(v))
	case string:
		r, err = strconv.ParseInt(string(v), 10, 64)
	default:
		err = fmt.Errorf("cannot recognize interface type")
	}
	if err != nil {
		return 0
	}
	return r
}

// InterfaceToFloat64 Convert interface{} to float64
func InterfaceToFloat64(n interface{}) float64 {
	var r float64
	var err error
	switch n := n.(type) {
	case int:
		r = float64(int(n))
	case int64:
		r = float64(int64(n))
	case float32:
		r = float64(float32(n))
	case float64:
		r = float64(n)
	case string:
		r, err = strconv.ParseFloat(string(n), 64)
	default:
		err = fmt.Errorf("cannot recognize interface type")
	}
	if err != nil || math.IsInf(r, 0) || math.IsNaN(r) {
		return 0
	}
	return r
}

// FmtFloatList format float list
func FmtFloatList(nums []float64) {
	for i := range nums {
		if math.IsInf(nums[i], 0) || math.IsNaN(nums[i]) {
			nums[i] = 0
		}
	}
}

// InterfaceToString Convert interface{} to string
func InterfaceToString(n interface{}) string {
	var s string
	var err error
	switch n := n.(type) {
	case int:
		s = strconv.Itoa(int(n))
	case int64:
		s = strconv.FormatInt(int64(n), 10)
	case float32:
		s = strconv.FormatFloat(float64(float32(n)), 'f', -1, 64)
	case float64:
		s = strconv.FormatFloat(float64(n), 'f', -1, 64)
	case string:
		s = string(n)
	case nil:
		s = ""
	default:
		err = fmt.Errorf("cannot recognize interface type")
	}
	if err != nil {
		return ""
	}
	return s
}

// FloatToString Convert float64 to string
func FloatToString(n float64) string {
	return strconv.FormatFloat(n, 'f', -1, 64)
}

// StringToFloat Convert string to float64, if the conversion fails, return 0
func StringToFloat(s string) float64 {
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return n
}

// StringToInt Convert string to int, if the conversion fails, return 0
func StringToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}

// StringToInt64 Convert string to int64, if the conversion fails, return 0
func StringToInt64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return 0
	}
	return n
}

// BoolToInt bool converted to int return 0/1
func BoolToInt(n bool) int {
	if n {
		return 1
	}
	return 0
}

// IntToBool integer conversion to bool
func IntToBool(n int) bool {
	return n != 0
}

// SafeDivideFloat64 Safe division
// zeroRtn defines what to return if the divisor is 0
func SafeDivideFloat64(a, b, zeroRtn float64) float64 {
	if math.Abs(b) < floatZero {
		return zeroRtn
	}
	return a / b
}

// MinInt Output the smallest of the list
func MinInt(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	mi := nums[0]
	for _, n := range nums {
		if n < mi {
			mi = n
		}
	}
	return mi
}

// MinInt64 Output the smallest of the list
func MinInt64(nums ...int64) int64 {
	if len(nums) == 0 {
		return 0
	}
	mi := nums[0]
	for _, n := range nums {
		if n < mi {
			mi = n
		}
	}
	return mi
}

// MaxInt Return the maximum value
func MaxInt(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	mx := nums[0]
	for _, n := range nums {
		if n > mx {
			mx = n
		}
	}
	return mx
}

// MaxInt64 Return the maximum value
func MaxInt64(nums ...int64) int64 {
	if len(nums) == 0 {
		return 0
	}
	mx := nums[0]
	for _, n := range nums {
		if n > mx {
			mx = n
		}
	}
	return mx
}

// MinFloat64 Output the smallest of the list
func MinFloat64(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	mi := nums[0]
	for _, n := range nums {
		if n < mi {
			mi = n
		}
	}
	return mi
}

// MaxFloat64 Return the maximum value
func MaxFloat64(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	mx := nums[0]
	for _, n := range nums {
		if n > mx {
			mx = n
		}
	}
	return mx
}
