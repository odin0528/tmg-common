package math_tool

import (
	"fmt"
	"math"
	"strconv"

	"github.com/shopspring/decimal"
)

func FloatAdd(fs ...float64) float64 {
	d := decimal.Zero
	for k := range fs {
		d = d.Add(decimal.NewFromFloat(fs[k]))
	}
	f, _ := d.Float64()
	return f
}

func FloatSub(fs ...float64) float64 {
	if len(fs) == 0 {
		return 0
	} else if len(fs) == 1 {
		return fs[0]
	}
	d := decimal.NewFromFloat(fs[0])
	for _, v := range fs[1:] {
		d = d.Sub(decimal.NewFromFloat(v))
	}
	f, _ := d.Float64()
	return f
}

func FloatMul(f1 float64, f2 float64) float64 {
	f, _ := decimal.NewFromFloat(f1).Mul(decimal.NewFromFloat(f2)).Float64()
	return f
}

func FloatDiv(f1 float64, f2 float64) float64 {
	if f2 == 0 {
		return f1
	}

	f, _ := decimal.NewFromFloat(f1).Div(decimal.NewFromFloat(f2)).Float64()
	return f
}

func IsFloatEqual(f1 float64, f2 float64) bool {
	return floatCompare(FLOAT_CMP_EQUAL, f1, f2)
}

func IsFloatGreaterThan(f1 float64, f2 float64) bool {
	return floatCompare(FLOAT_CMP_GREATER_THAN, f1, f2)
}

func IsFloatGreaterThanOrEqual(f1 float64, f2 float64) bool {
	return floatCompare(FLOAT_CMP_GREATER_THAN_OR_EQUAL, f1, f2)
}

func IsFloatLessThan(f1 float64, f2 float64) bool {
	return floatCompare(FLOAT_CMP_LESS_THAN, f1, f2)
}

func IsFloatLessThanOrEqual(f1 float64, f2 float64) bool {
	return floatCompare(FLOAT_CMP_LESS_THAN_OR_EQUAL, f1, f2)
}

func floatCompare(cmp FloatCompare, f1 float64, f2 float64) bool {
	d1 := decimal.NewFromFloat(f1)
	d2 := decimal.NewFromFloat(f2)

	switch cmp {
	case FLOAT_CMP_EQUAL:
		return d1.Equal(d2)
	case FLOAT_CMP_GREATER_THAN:
		return d1.GreaterThan(d2)
	case FLOAT_CMP_GREATER_THAN_OR_EQUAL:
		return d1.GreaterThanOrEqual(d2)
	case FLOAT_CMP_LESS_THAN:
		return d1.LessThan(d2)
	case FLOAT_CMP_LESS_THAN_OR_EQUAL:
		return d1.LessThanOrEqual(d2)
	}

	return false
}

func Cmp(f1 float64, f2 float64) CmpVal {
	/*	-1 if f1 <  f2
	 	 0 if f1 == f2
		+1 if f1 >  f2 */

	d1 := decimal.NewFromFloat(f1)
	d2 := decimal.NewFromFloat(f2)

	return CmpVal(d1.Cmp(d2))
}

func GetMoneyFloatDecimal(retVal float64) float64 {
	if CMP_BIGGER_THAN == Cmp(retVal, 0.0) {
		retVal = GetFloatDecimal(retVal, ROUND_PRECISION)
	} else {
		retVal = GetFloatDecimal(retVal, ROUND_PRECISION+1)
	}

	return retVal
}

func GetFloatDecimal(val float64, place int) float64 {
	var tail float64

	format := "%." + strconv.Itoa(place+1) + "f"
	truncVal := math.Trunc(val)
	tail = val - truncVal
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
