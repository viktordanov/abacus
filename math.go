package main

import (
	"math"
	"math/big"
)

func Exp(z *big.Float) *big.Float {
	if z.Sign() == 0 {
		return big.NewFloat(1).SetPrec(z.Prec())
	}
	if z.IsInf() && z.Sign() > 0 {
		return big.NewFloat(math.Inf(+1)).SetPrec(z.Prec())
	}
	if z.IsInf() && z.Sign() < 0 {
		return big.NewFloat(0).SetPrec(z.Prec())
	}

	guess := new(big.Float)

	zf, _ := z.Float64()
	if zfs := math.Exp(zf); zfs == math.Inf(+1) || zfs == 0 {

		halfZ := new(big.Float).Mul(z, big.NewFloat(0.5))
		halfExp := Exp(halfZ.SetPrec(z.Prec() + 64))
		return new(big.Float).Mul(halfExp, halfExp).SetPrec(z.Prec())
	} else {
		guess.SetFloat64(zfs)
	}

	f := func(t *big.Float) *big.Float {
		x := new(big.Float)
		x.Sub(Log(t), z)
		return x.Mul(x, t)
	}

	x := newton(f, guess, z.Prec())

	return x
}

func Log(z *big.Float) *big.Float {
	if z.Sign() == -1 {
		panic("Log: argument is negative")
	}
	if z.Sign() == 0 {
		return big.NewFloat(math.Inf(-1)).SetPrec(z.Prec())
	}

	prec := z.Prec() + 64

	one := big.NewFloat(1).SetPrec(prec)
	two := big.NewFloat(2).SetPrec(prec)
	four := big.NewFloat(4).SetPrec(prec)

	if z.Cmp(one) == 0 {
		return big.NewFloat(0).SetPrec(z.Prec())
	}
	if z.IsInf() {
		return big.NewFloat(math.Inf(+1)).SetPrec(z.Prec())
	}

	x := new(big.Float).SetPrec(prec)

	var neg bool
	if z.Cmp(one) < 0 {
		x.Quo(one, z)
		neg = true
	} else {
		x.Set(z)
	}

	lim := new(big.Float)
	lim.SetMantExp(two, int(prec/2))

	k := 0
	for x.Cmp(lim) < 0 {
		x.Mul(x, x)
		k++
	}

	pi := pi(prec)
	agm := agm(one, x.Quo(four, x)) // agm = AGM(1, 4/x)

	x.Quo(pi, x.Mul(two, agm)) // reuse x, we don't need it

	if neg {
		x.Neg(x)
	}

	x.Mul(x, lim.SetMantExp(one, -k))

	return x.SetPrec(z.Prec())
}

func Abs(a *big.Float) *big.Float {
	return Zero().Abs(a)
}

func New(f float64) *big.Float {
	r := big.NewFloat(f)
	r.SetPrec(256)
	return r
}

func Div(a, b *big.Float) *big.Float {
	return Zero().Quo(a, b)
}

func Zero() *big.Float {
	r := big.NewFloat(0.0)
	r.SetPrec(256)
	return r
}

func Mul(a, b *big.Float) *big.Float {
	return Zero().Mul(a, b)
}

func Add(a, b *big.Float) *big.Float {
	return Zero().Add(a, b)
}

func Sub(a, b *big.Float) *big.Float {
	return Zero().Sub(a, b)
}

func Lesser(x, y *big.Float) bool {
	return x.Cmp(y) == -1
}

func agm(a, b *big.Float) *big.Float {
	if a.Prec() != b.Prec() {
		panic("agm: different precisions")
	}
	prec := a.Prec()

	a2 := new(big.Float).Copy(a).SetPrec(prec + 64)
	b2 := new(big.Float).Copy(b).SetPrec(prec + 64)

	if a2.Cmp(b2) == -1 {
		a2, b2 = b2, a2
	}

	lim := new(big.Float)
	lim.SetMantExp(big.NewFloat(1).SetPrec(prec+64), -int(prec+1))

	half := big.NewFloat(0.5)
	t := new(big.Float)

	for t.Sub(a2, b2).Cmp(lim) != -1 {
		t.Copy(a2)
		a2.Add(a2, b2).Mul(a2, half)
		b2 = Sqrt(b2.Mul(b2, t))
	}

	return a2.SetPrec(prec)
}

var piCache *big.Float
var eCache *big.Float
var phiCache *big.Float
var piCachePrec uint
var enablePiCache = true

func init() {
	piCache, _, _ = new(big.Float).SetPrec(1024).Parse("3."+
		"14159265358979323846264338327950288419716939937510"+
		"58209749445923078164062862089986280348253421170679"+
		"82148086513282306647093844609550582231725359408128"+
		"48111745028410270193852110555964462294895493038196"+
		"44288109756659334461284756482337867831652712019091"+
		"45648566923460348610454326648213393607260249141273"+
		"72458700660631558817488152092096282925409171536444", 10)

	piCachePrec = 1024

	eCache, _, _ = new(big.Float).SetPrec(1024).Parse("2."+
		"71828182845904523536028747135266249775724709369995"+
		"95749669676277240766303535475945713821785251664274"+
		"27466391932003059921817413596629043572900334295260"+
		"59563073813232862794349076323382988075319525101901"+
		"15738341879307021540891499348841675092447614606680"+
		"82264800168477411853742345442437107539077744992069"+
		"55170276183860626133138458300075204493382656029760", 10)

	phiCache, _, _ = new(big.Float).SetPrec(1024).Parse("1."+
		"61803398874989484820458683436563811772030917980576"+
		"28621354486227052604628189024497072072041893911374"+
		"84754088075386891752126633862223536931793180060766"+
		"72635443338908659593958290563832266131992829026788"+
		"06752087668925017116962070322210432162695486262963"+
		"13614438149758701220340805887954454749246185695364"+
		"86444924104432077134494704956584678850987433944221", 10)

}

func e(prec uint) *big.Float {
	return new(big.Float).Copy(eCache).SetPrec(prec)
}

func phi(prec uint) *big.Float {
	return new(big.Float).Copy(phiCache).SetPrec(prec)
}

func pi(prec uint) *big.Float {

	if prec <= piCachePrec && enablePiCache {
		return new(big.Float).Copy(piCache).SetPrec(prec)
	}

	// Following R. P. Brent, Multiple-precision zero-finding
	// methods and the complexity of elementary function evaluation,
	// in Analytic Computational Complexity, Academic Press,
	// New York, 1975, Section 8.
	half := big.NewFloat(0.5)
	two := big.NewFloat(2).SetPrec(prec + 64)

	a := big.NewFloat(1).SetPrec(prec + 64)    // a = 1
	b := new(big.Float).Mul(Sqrt(two), half)   // b = 1/√2
	t := big.NewFloat(0.25).SetPrec(prec + 64) // t = 1/4
	x := big.NewFloat(1).SetPrec(prec + 64)    // x = 1

	lim := new(big.Float)
	lim.SetMantExp(big.NewFloat(1).SetPrec(prec+64), -int(prec+1))

	y := new(big.Float)
	for y.Sub(a, b).Cmp(lim) != -1 { // assume a > b
		y.Copy(a)
		a.Add(a, b).Mul(a, half) // a = (a+b)/2
		b = Sqrt(b.Mul(b, y))    // b = √(ab)

		y.Sub(a, y)           // y = a - y
		y.Mul(y, y).Mul(y, x) // y = x(a-y)²
		t.Sub(t, y)           // t = t - x(a-y)²
		x.Mul(x, two)         // x = 2x
	}

	a.Mul(a, a).Quo(a, t) // π = a² / t
	a.SetPrec(prec)

	if enablePiCache {
		piCache.Copy(a)
		piCachePrec = prec
	}

	return a
}

func newton(fOverDf func(z *big.Float) *big.Float, guess *big.Float, dPrec uint) *big.Float {
	prec, guard := guess.Prec(), uint(64)
	guess.SetPrec(prec + guard)

	for prec < 2*dPrec {
		guess.Sub(guess, fOverDf(guess))
		prec *= 2
		guess.SetPrec(prec + guard)
	}

	return guess.SetPrec(dPrec)
}

func Pow(z *big.Float, w *big.Float) *big.Float {
	if z.Sign() < 0 {
		panic("Pow: negative base")
	}
	if w.Sign() == 0 {
		return big.NewFloat(1).SetPrec(z.Prec())
	}
	if w.Cmp(big.NewFloat(1)) == 0 || z.IsInf() {
		return new(big.Float).Copy(z)
	}

	if w.Sign() < 0 {
		x := new(big.Float)
		zExt := new(big.Float).Copy(z).SetPrec(z.Prec() + 64)
		wNeg := new(big.Float).Neg(w)
		return x.Quo(big.NewFloat(1), Pow(zExt, wNeg)).SetPrec(z.Prec())
	}

	if false && w.IsInt() {
		wi, _ := w.Int64()
		return powInt(z, int(wi))
	}

	x := new(big.Float).SetPrec(z.Prec() + 64)
	logZ := Log(new(big.Float).Copy(z).SetPrec(z.Prec() + 64))
	x.Mul(w, logZ)
	x = Exp(x)
	return x.SetPrec(z.Prec())

}

func powInt(z *big.Float, w int) *big.Float {
	mant := new(big.Float)
	exp := z.MantExp(mant)

	exp = exp * w

	x := big.NewFloat(1).SetPrec(z.Prec())

	for w > 0 {
		if w%2 == 1 {
			x.Mul(x, mant)
		}
		w >>= 1
		mant.Mul(mant, mant)
	}

	return new(big.Float).SetMantExp(x, exp)
}

func Sqrt(z *big.Float) *big.Float {
	if z.Sign() == -1 {
		panic("Sqrt: argument is negative")
	}
	if z.Sign() == 0 {
		return big.NewFloat(float64(z.Sign()))
	}
	if z.IsInf() {
		return big.NewFloat(math.Inf(+1))
	}

	mant := new(big.Float)
	exp := z.MantExp(mant)
	switch exp % 2 {
	case 1:
		mant.Mul(big.NewFloat(2), mant)
	case -1:
		mant.Mul(big.NewFloat(0.5), mant)
	}

	var x *big.Float
	if z.Prec() <= 128 {
		x = sqrtDirect(mant)
	} else {
		x = sqrtInverse(mant)
	}

	return x.SetMantExp(x, exp/2)
}

func sqrtDirect(z *big.Float) *big.Float {
	half := big.NewFloat(0.5)
	f := func(t *big.Float) *big.Float {
		x := new(big.Float).Mul(t, t)
		x.Sub(x, z)
		x.Mul(half, x)
		return x.Quo(x, t)
	}

	zf, _ := z.Float64()
	guess := big.NewFloat(math.Sqrt(zf))

	return newton(f, guess, z.Prec())
}

func sqrtInverse(z *big.Float) *big.Float {
	nhalf := big.NewFloat(-0.5)
	one := big.NewFloat(1)
	f := func(t *big.Float) *big.Float {
		u := new(big.Float)
		u.Mul(t, t)
		u.Mul(u, z)
		u.Sub(one, u)
		u.Mul(u, nhalf)
		return new(big.Float).Mul(t, u)
	}

	zf, _ := z.Float64()
	guess := big.NewFloat(1 / math.Sqrt(zf))

	x := newton(f, guess, z.Prec()+32)
	return x.Mul(z, x).SetPrec(z.Prec())
}
