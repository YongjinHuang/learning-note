package fz

type FizzBuzz struct {
	n   int
	f   chan struct{}
	b   chan struct{}
	fb  chan struct{}
	chn chan struct{}
}

func NewFizzBuzz(n int) FizzBuzz {
	return FizzBuzz{
		n:   n,
		f:   make(chan struct{}),
		b:   make(chan struct{}),
		fb:  make(chan struct{}),
		chn: make(chan struct{}, 1),
	}
}

func (fizzBuzz *FizzBuzz) Fizz(pf func()) {
	for i := 1; i <= fizzBuzz.n; i++ {
		switch getFbType(i) {
		case typeFizz:
			fizzBuzz.f <- struct{}{}
			pf()
			fizzBuzz.transfer(i + 1)
		default:
			continue
		}
	}
}

func (fizzBuzz *FizzBuzz) transfer(i int) {
	if i > fizzBuzz.n {
		return
	}
	switch getFbType(i) {
	case typeFizz:
		<-fizzBuzz.f
	case typeBuzz:
		<-fizzBuzz.b
	case typeFizzBuzz:
		<-fizzBuzz.fb
	default:
		<-fizzBuzz.chn
	}
}

func (fizzBuzz *FizzBuzz) Buzz(pb func()) {
	for i := 1; i <= fizzBuzz.n; i++ {
		switch getFbType(i) {
		case typeBuzz:
			fizzBuzz.b <- struct{}{}
			pb()
			fizzBuzz.transfer(i + 1)
		default:
			continue
		}
	}
}

func (fizzBuzz *FizzBuzz) FizzBuzz(pfb func()) {
	for i := 1; i <= fizzBuzz.n; i++ {
		switch getFbType(i) {
		case typeFizzBuzz:
			fizzBuzz.fb <- struct{}{}
			pfb()
			fizzBuzz.transfer(i + 1)
		default:
			continue
		}
	}
}

func (fizzBuzz *FizzBuzz) Number(pn func(int)) {
	for i := 1; i <= fizzBuzz.n; i++ {
		switch getFbType(i) {
		case typeNumber:
			fizzBuzz.chn <- struct{}{}
			pn(i)
			fizzBuzz.transfer(i + 1)
		default:
			continue
		}
	}
}

type fBType uint8

const (
	typeFizz fBType = iota + 1
	typeBuzz
	typeFizzBuzz
	typeNumber
)

func getFbType(n int) fBType {
	if n%3 == 0 && n%5 == 0 {
		return typeFizzBuzz
	}
	if n%3 == 0 {
		return typeFizz
	}
	if n%5 == 0 {
		return typeBuzz
	}
	return typeNumber
}
