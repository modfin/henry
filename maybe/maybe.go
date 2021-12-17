package maybe

func Some[A any](val A) Value[A] {
	return Value[A]{
		val: &val,
	}
}

func None[A any](_ A) Value[A] {
	return Value[A]{
		val: nil,
	}
}

type Value[A any] struct {
	val *A
}

func (o Value[A]) Valid() bool {
	return o.val != nil
}

func (o Value[A]) Get() (A, bool) {
	if o.Valid() {
		return *o.val, true
	}
	var zero A
	return zero, false
}
func (o Value[A]) Or(fallback A) A {
	a, ok := o.Get()
	if ok {
		return a
	}
	return fallback
}

func (o Value[A]) OrGet(getter func() A) A {
	a, ok := o.Get()
	if ok {
		return a
	}
	return getter()
}
