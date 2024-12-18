package math_tool

import (
	"fmt"
	"math"
	"strconv"

	"github.com/shopspring/decimal"
)

func FloatAdd(terms ...float64) float64 {
	temp := decimal.Zero
	for _, term := range terms {
		temp = temp.Add(decimal.NewFromFloat(term))
	}

	result, _ := temp.Float64()
	return result
}

func FloatSub(terms ...float64) float64 {
	if len(terms) == 0 {
		return 0
	} else if len(terms) == 1 {
		return terms[0]
	}

	temp := decimal.NewFromFloat(terms[0])
	for _, term := range terms[1:] {
		temp = temp.Sub(decimal.NewFromFloat(term))
	}

	result, _ := temp.Float64()
	return result
}

func FloatMul(multiplier float64, multiplicand float64) float64 {
	result, _ := decimal.NewFromFloat(multiplier).Mul(decimal.NewFromFloat(multiplicand)).Float64()
	return result
}

func FloatDiv(dividend float64, divisor float64) float64 {
	if divisor == 0 {
		return dividend
	}

	result, _ := decimal.NewFromFloat(dividend).Div(decimal.NewFromFloat(divisor)).Float64()
	return result
}

func IsFloatEqual(src float64, dst float64) bool {
	return floatCompare(FLOAT_CMP_EQUAL, src, dst)
}

func IsFloatGreaterThan(src float64, dst float64) bool {
	return floatCompare(FLOAT_CMP_GREATER_THAN, src, dst)
}

func IsFloatGreaterThanOrEqual(src float64, dst float64) bool {
	return floatCompare(FLOAT_CMP_GREATER_THAN_OR_EQUAL, src, dst)
}

func IsFloatLessThan(src float64, dst float64) bool {
	return floatCompare(FLOAT_CMP_LESS_THAN, src, dst)
}

func IsFloatLessThanOrEqual(src float64, dst float64) bool {
	return floatCompare(FLOAT_CMP_LESS_THAN_OR_EQUAL, src, dst)
}

func floatCompare(cmp FloatCompare, src float64, dst float64) bool {
	srcDecimal := decimal.NewFromFloat(src)
	dstDecimal := decimal.NewFromFloat(dst)

	switch cmp {
	case FLOAT_CMP_EQUAL:
		return srcDecimal.Equal(dstDecimal)
	case FLOAT_CMP_GREATER_THAN:
		return srcDecimal.GreaterThan(dstDecimal)
	case FLOAT_CMP_GREATER_THAN_OR_EQUAL:
		return srcDecimal.GreaterThanOrEqual(dstDecimal)
	case FLOAT_CMP_LESS_THAN:
		return srcDecimal.LessThan(dstDecimal)
	case FLOAT_CMP_LESS_THAN_OR_EQUAL:
		return srcDecimal.LessThanOrEqual(dstDecimal)
	}

	return false
}

func Cmp(src float64, dst float64) CmpValue {
	srcDecimal := decimal.NewFromFloat(src)
	dstDecimal := decimal.NewFromFloat(dst)

	return CmpValue(srcDecimal.Cmp(dstDecimal))
}

func GetMoneyFloatDecimal(value float64) float64 {
	if CMP_BIGGER_THAN == Cmp(value, 0.0) {
		value = GetFloatDecimal(value, ROUND_PRECISION)
	} else {
		value = GetFloatDecimal(value, ROUND_PRECISION+1)
	}

	return value
}

func GetAwardMoneyFloatDecimal(value float64) float64 {
	if CMP_BIGGER_THAN == Cmp(value, 0.0) {
		value = GetFloatDecimal(value, AWARD_PRECISION)
	} else {
		value = GetFloatDecimal(value, AWARD_PRECISION+1)
	}

	return value
}

func GetFloatDecimal(value float64, place int) float64 {
	var tail float64

	format := "%." + strconv.Itoa(place+1) + "f"
	truncVal := math.Trunc(value)
	tail = value - truncVal
	strDecimal := fmt.Sprintf(format, tail)
	strDecimal = strDecimal[:place+2]
	tail, _ = strconv.ParseFloat(strDecimal, 64)

	srcDecimal := decimal.NewFromFloat(truncVal)
	opDecimal := decimal.NewFromFloat(tail)

	srcDecimal = srcDecimal.Add(opDecimal)

	result, _ := srcDecimal.Float64()

	return result
}

func Abs(value float64) float64 {
	if CMP_LESS_THAN == Cmp(value, 0.0) {
		value = FloatSub(0, value)
	}

	return value
}
func GetDisplayFloatDecimalUp(value float64) float64 {

	return makeDisplayRoundUp(value, ROUND_DISPLAY_PRECISION)
}
func GetDisplayFloatDecimalDown(value float64) float64 {

	return makeDisplayRoundDown(value, ROUND_DISPLAY_PRECISION)
}

func makeDisplayRoundUp(value float64, n int32) float64 {
	decimal := decimal.NewFromFloat(value)
	v, _ := decimal.RoundUp(n).Float64()
	return v
}

func makeDisplayRoundDown(value float64, n int32) float64 {
	decimal := decimal.NewFromFloat(value)
	v, _ := decimal.RoundDown(n).Float64()
	return v
}
