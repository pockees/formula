package fs

import "github.com/yidane/formula/opt"

type ExpFunction struct {
}

func (*ExpFunction) Name() string {
	panic("implement me")
}

func (*ExpFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewExpFunction() *ExpFunction  {
	return &ExpFunction{}
}

