package types
// adaptation of osmomath
import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var powPrecision, _ = sdk.NewDecFromStr("0.00000001")

// singletons
var zero sdk.Dec = sdk.ZeroDec()

var (
	one_half sdk.Dec = sdk.MustNewDecFromStr("0.5")
	one sdk.Dec = sdk.OneDec()
	two sdk.Dec = sdk.MustNewDecFromStr("2")
)

// returns internal power precision
func GetPowPrecision() sdk.Dec {
	return powPrecision.Clone()
}

// absolute diff w/ sign
func AbsDifferenceWithSign(a, b sdk.Dec) (sdk.Dec, bool) {
	if a.GTE(b){
		return a.SubMut(b), false
	} else {
		return a.NegMut().AddMut(b), true
	}
}

func Pow(base sdk.Dec, exp sdk.Dec) sdk.Dec {
	if !base.IsPositive() {
		panic(fmt.Errorf("base must be positive"))
	}
	if base.GTE(two) {
		panic(fmt.Errorf("base must be less than two"))
	}
	integer := exp.TruncateDec()
	fractional := exp.Sub(integer)

	integerPow := base.Power(uint64(integer.TruncateInt64()))

	if fractional.IsZero(){
		return integerPow
	}

	fractionalPow := PowApprox(base, fractional, powPrecision)

	return integerPow.Mul(fractionalPow)
}

func PowApprox(base sdk.Dec, exp sdk.Dec, precision sdk.Dec) sdk.Dec {
	if exp.IsZero(){
		return sdk.ZeroDec()
	}
	if exp.Equal(one_half) {
		output, err := base.ApproxSqrt()
		if err != nil {
			panic(err)
		}
		return output
	}

	base = base.Clone()
	x, xneg := AbsDifferenceWithSign(base, one)
	term := sdk.OneDec()
	sum := sdk.OneDec()
	negative := false
	a := exp.Clone()
	bigK := sdk.NewDec(0)
	
	for i := int64(1); term.GTE(precision); i++ {
		c, cneg := AbsDifferenceWithSign(a, bigK)
		bigK.Set(sdk.NewDec(i))
		term.MulMut(c).MulMut(x).QuoMut(bigK)
		a.Set(exp)
		if term.IsZero() {
			break
		}
		if xneg {
			negative = !negative
		}
		if cneg {
			negative = !negative
		}
		if negative {
			sum.SubMut(term)
		} else {
			sum.AddMut(term)
		}
	} 
	return sum
}