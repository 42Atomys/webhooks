package factory

import (
	"fmt"
	"reflect"
)

type compareFactory struct{ Factory }

func (*compareFactory) Name() string {
	return "compare"
}

func (*compareFactory) DefinedInpus() []*Var {
	return []*Var{
		{false, reflect.TypeOf(&InputConfig{}), "first", &InputConfig{}},
		{false, reflect.TypeOf(&InputConfig{}), "second", &InputConfig{}},
	}
}

func (*compareFactory) DefinedOutputs() []*Var {
	return []*Var{
		{false, reflect.TypeOf(false), "result", false},
	}
}

func (c *compareFactory) Func() RunFunc {
	return func(factory *Factory, configRaw map[string]interface{}) error {
		firstVar, ok := factory.Input("first")
		if !ok {
			return fmt.Errorf("missing input first")
		}

		secondVar, ok := factory.Input("second")
		if !ok {
			return fmt.Errorf("missing input second")
		}

		factory.Output("result", c.compareSlice(
			firstVar.Value.(*InputConfig).Get(),
			secondVar.Value.(*InputConfig).Get(),
		))
		return nil
	}
}

func (*Factory) compareSlice(slice1, slice2 []string) bool {
	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			for _, s2 := range slice2 {
				if s1 == s2 {
					return true
				}
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}
	return false
}
