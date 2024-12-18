package math_tool

import "sync"

type FloatOperation string
type FloatCompare string
type CmpValue int

const (
	FLOAT_OP_ADD FloatOperation = "add"
	FLOAT_OP_SUB FloatOperation = "sub"
	FLOAT_OP_MUL FloatOperation = "mul"
	FLOAT_OP_DIV FloatOperation = "div"

	FLOAT_CMP_EQUAL                 FloatCompare = "equal"
	FLOAT_CMP_GREATER_THAN          FloatCompare = "greater_than"
	FLOAT_CMP_GREATER_THAN_OR_EQUAL FloatCompare = "greater_than_or_equal"
	FLOAT_CMP_LESS_THAN             FloatCompare = "less_than"
	FLOAT_CMP_LESS_THAN_OR_EQUAL    FloatCompare = "less_than_or_equal"

	ROUND_PRECISION         = 6
	ROUND_DISPLAY_PRECISION = 2
	AWARD_PRECISION         = 4

	CMP_LESS_THAN   CmpValue = -1
	CMP_EQUAL       CmpValue = 0
	CMP_BIGGER_THAN CmpValue = 1
)

var randomOnce sync.Once
