package mon

import (
	"testing"
)

func TestSome(t *testing.T) {
	o := Some(42)
	if !o.Some() {
		t.Error("Expected Some() to return true")
	}
	if o.None() {
		t.Error("Expected None() to return false")
	}
	val, ok := o.Get()
	if !ok {
		t.Error("Expected Get() to return true")
	}
	if val != 42 {
		t.Errorf("Expected Get() to return 42, got %v", val)
	}
}

func TestNone(t *testing.T) {
	o := None[int]()
	if o.Some() {
		t.Error("Expected Some() to return false")
	}
	if !o.None() {
		t.Error("Expected None() to return true")
	}
	val, ok := o.Get()
	if ok {
		t.Error("Expected Get() to return false")
	}
	if val != 0 {
		t.Errorf("Expected Get() to return zero value, got %v", val)
	}
}

func TestTupleToOption(t *testing.T) {
	// Some case
	o := TupleToOption(42, true)
	if !o.Some() {
		t.Error("Expected Some() to return true")
	}
	val, _ := o.Get()
	if val != 42 {
		t.Errorf("Expected value 42, got %v", val)
	}

	// None case
	o = TupleToOption(42, false)
	if !o.None() {
		t.Error("Expected None() to return true")
	}
}

func TestEmptyableToOption(t *testing.T) {
	// Some case - non-zero value
	o := EmptyableToOption(42)
	if !o.Some() {
		t.Error("Expected Some() to return true for non-zero value")
	}
	val, _ := o.Get()
	if val != 42 {
		t.Errorf("Expected value 42, got %v", val)
	}

	// None case - zero value
	o = EmptyableToOption(0)
	if !o.None() {
		t.Error("Expected None() to return true for zero value")
	}

	// String case - non-empty
	os := EmptyableToOption("hello")
	if !os.Some() {
		t.Error("Expected Some() to return true for non-empty string")
	}

	// String case - empty
	os = EmptyableToOption("")
	if !os.None() {
		t.Error("Expected None() to return true for empty string")
	}
}

func TestPointerToOption(t *testing.T) {
	// Some case - non-nil pointer
	val := 42
	o := PointerToOption(&val)
	if !o.Some() {
		t.Error("Expected Some() to return true for non-nil pointer")
	}
	v, _ := o.Get()
	if v != 42 {
		t.Errorf("Expected value 42, got %v", v)
	}

	// None case - nil pointer
	o = PointerToOption[int](nil)
	if !o.None() {
		t.Error("Expected None() to return true for nil pointer")
	}
}

func TestOptionMustGet(t *testing.T) {
	// Some case
	o := Some(42)
	val := o.MustGet()
	if val != 42 {
		t.Errorf("Expected value 42, got %v", val)
	}

	// None case - panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected MustGet to panic on None")
		}
	}()
	o = None[int]()
	o.MustGet()
}

func TestOptionOrElse(t *testing.T) {
	// Some case
	o := Some(42)
	val := o.OrElse(100)
	if val != 42 {
		t.Errorf("Expected value 42, got %v", val)
	}

	// None case
	o = None[int]()
	val = o.OrElse(100)
	if val != 100 {
		t.Errorf("Expected fallback value 100, got %v", val)
	}
}

func TestOptionOrEmpty(t *testing.T) {
	// Some case
	o := Some(42)
	val := o.OrEmpty()
	if val != 42 {
		t.Errorf("Expected value 42, got %v", val)
	}

	// None case
	o = None[int]()
	val = o.OrEmpty()
	if val != 0 {
		t.Errorf("Expected zero value 0, got %v", val)
	}
}

func TestOptionForEach(t *testing.T) {
	// Some case
	called := false
	o := Some(42)
	o.ForEach(func(val int) {
		called = true
		if val != 42 {
			t.Errorf("Expected value 42, got %v", val)
		}
	})
	if !called {
		t.Error("Expected ForEach to be called")
	}

	// None case
	called = false
	o = None[int]()
	o.ForEach(func(val int) {
		called = true
	})
	if called {
		t.Error("Expected ForEach not to be called on None")
	}
}

func TestOptionMatch(t *testing.T) {
	// Some case
	o := Some(42)
	result := o.Match(
		func(val int) (int, bool) {
			return val * 2, true
		},
		func() (int, bool) {
			return 0, false
		},
	)
	if !result.Some() {
		t.Error("Expected result to be Some")
	}
	val, _ := result.Get()
	if val != 84 {
		t.Errorf("Expected value 84, got %v", val)
	}

	// None case
	o = None[int]()
	result = o.Match(
		func(val int) (int, bool) {
			return val * 2, true
		},
		func() (int, bool) {
			return 100, true
		},
	)
	if !result.Some() {
		t.Error("Expected result to be Some after None handler")
	}
	val, _ = result.Get()
	if val != 100 {
		t.Errorf("Expected value 100, got %v", val)
	}

	// None to None case
	o = None[int]()
	result = o.Match(
		func(val int) (int, bool) {
			return val * 2, true
		},
		func() (int, bool) {
			return 0, false
		},
	)
	if !result.None() {
		t.Error("Expected result to be None")
	}
}

func TestOptionMap(t *testing.T) {
	// Some case
	o := Some(42)
	result := o.Map(func(val int) (int, bool) {
		return val * 2, true
	})
	if !result.Some() {
		t.Error("Expected result to be Some")
	}
	val, _ := result.Get()
	if val != 84 {
		t.Errorf("Expected value 84, got %v", val)
	}

	// Some to None case
	o = Some(42)
	result = o.Map(func(val int) (int, bool) {
		return 0, false
	})
	if !result.None() {
		t.Error("Expected result to be None")
	}

	// None case
	o = None[int]()
	result = o.Map(func(val int) (int, bool) {
		return val * 2, true
	})
	if !result.None() {
		t.Error("Expected result to be None")
	}
}

func TestOptionMapNone(t *testing.T) {
	// Some case - no mapping
	o := Some(42)
	result := o.MapNone(func() (int, bool) {
		return 100, true
	})
	if !result.Some() {
		t.Error("Expected result to be Some")
	}
	val, _ := result.Get()
	if val != 42 {
		t.Errorf("Expected value 42, got %v", val)
	}

	// None case - successful mapping
	o = None[int]()
	result = o.MapNone(func() (int, bool) {
		return 100, true
	})
	if !result.Some() {
		t.Error("Expected result to be Some after MapNone")
	}
	val, _ = result.Get()
	if val != 100 {
		t.Errorf("Expected value 100, got %v", val)
	}

	// None case - mapping returns None
	o = None[int]()
	result = o.MapNone(func() (int, bool) {
		return 0, false
	})
	if !result.None() {
		t.Error("Expected result to be None")
	}
}

func TestOptionFlatMap(t *testing.T) {
	// Some case
	o := Some(42)
	result := o.FlatMap(func(val int) Option[int] {
		return Some(val * 2)
	})
	if !result.Some() {
		t.Error("Expected result to be Some")
	}
	val, _ := result.Get()
	if val != 84 {
		t.Errorf("Expected value 84, got %v", val)
	}

	// Some to None case
	o = Some(42)
	result = o.FlatMap(func(val int) Option[int] {
		return None[int]()
	})
	if !result.None() {
		t.Error("Expected result to be None")
	}

	// None case
	o = None[int]()
	result = o.FlatMap(func(val int) Option[int] {
		return Some(val * 2)
	})
	if !result.None() {
		t.Error("Expected result to be None")
	}
}

func TestOptionToPointer(t *testing.T) {
	// Some case
	o := Some(42)
	ptr := o.ToPointer()
	if ptr == nil {
		t.Error("Expected pointer to be non-nil")
	}
	if *ptr != 42 {
		t.Errorf("Expected value 42, got %v", *ptr)
	}

	// None case
	o = None[int]()
	ptr = o.ToPointer()
	if ptr != nil {
		t.Error("Expected pointer to be nil")
	}
}
