//
// GENERATED BY EASYLANG COMPILER.
// !!!! DON'T MODIFY IT!!!!!!
//

package native

import (
	. "github.com/stephenlyu/goformula/stockfunc/function"
	. "github.com/stephenlyu/goformula/function"
	. "github.com/stephenlyu/goformula/formulalibrary/base/formula"
	. "github.com/stephenlyu/goformula/formulalibrary/native/nativeformulas"
)

type ma struct {
	BaseNativeFormula

	// Data of all referenced period


	// Referenced Formulas


	// Vectors
    m1 Value
    m2 Value
    m3 Value
    m4 Value
    var1 Value
    var2 Value
    ma1 Value
    var3 Value
    ma2 Value
    var4 Value
    ma3 Value
    var5 Value
    ma4 Value
}

var (
	MA_META = &FormulaMetaImpl{
		Name: "MA",
		ArgNames: []string{"m1", "m2", "m3", "m4"},
		ArgMeta: []Arg {
			Arg{5.000000, 0.000000, 1000.000000},
			Arg{10.000000, 0.000000, 1000.000000},
			Arg{20.000000, 0.000000, 1000.000000},
			Arg{60.000000, 0.000000, 1000.000000},
		},
		Flags: []int{0x00000000, 0x00000000, 0x00000000, 0x00000000},
		Colors: []*Color{{Red:255, Green:255, Blue:255}, {Red:255, Green:0, Blue:0}, {Red:0, Green:255, Blue:255}, {Red:0, Green:255, Blue:0}},
		LineThicks: []int{1, 1, 1, 1},
		LineStyles: []int{0, 0, 0, 0},
		GraphTypes: []int{1, 1, 1, 1},
		Vars: []string{"MA1", "MA2", "MA3", "MA4"},
	}
)

func NewMA(data *RVector, args []float64) Formula {
	o := &ma{
		BaseNativeFormula: BaseNativeFormula{
			FormulaMetaImpl: MA_META,
			Data__: data,
		},
	}

	// Data of all referenced period


	// Referenced Formulas


	// Vectors
    o.m1 = Scalar(args[0])
    o.m2 = Scalar(args[1])
    o.m3 = Scalar(args[2])
    o.m4 = Scalar(args[3])
    o.var1 = CLOSE(o.Data__)
    o.var2 = MA(o.var1, o.m1)
    o.ma1 = o.var2
    o.var3 = MA(o.var1, o.m2)
    o.ma2 = o.var3
    o.var4 = MA(o.var1, o.m3)
    o.ma3 = o.var4
    o.var5 = MA(o.var1, o.m4)
    o.ma4 = o.var5

	// Actions

    o.DrawActions__ = []DrawAction{

    }

	o.RefValues__ = []Value {o.ma1, o.ma2, o.ma3, o.ma4}
	return o
}

func (this *ma) UpdateLastValue() {
    this.var1.UpdateLastValue()
    this.var2.UpdateLastValue()
    this.var3.UpdateLastValue()
    this.var4.UpdateLastValue()
    this.var5.UpdateLastValue()
}

func init() {
	RegisterNativeFormula(NewMA, MA_META)
}

	