package fs

import "github.com/yidane/formula/opt"

type RoundFunction struct {
}

func (*RoundFunction) Name() string {
	return "round"
}

func (*RoundFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewRoundFunction() *RoundFunction {
	return &RoundFunction{}
}
