package fs

import "github.com/yidane/formula/opt"

type MultiplyFunction struct {
}

func (*MultiplyFunction) Name() string {
	return "*"
}

func (*MultiplyFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewMultiplyFunction() *MultiplyFunction {
	return &MultiplyFunction{}
}
