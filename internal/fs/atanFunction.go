package fs

import "github.com/yidane/formula/opt"

type AtanFunction struct {
}

func (*AtanFunction) Name() string {
	return "atan"
}

func (*AtanFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewAtanFunction() *AtanFunction {
	return &AtanFunction{}
}
