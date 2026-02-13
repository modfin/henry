package mon

import (
	"errors"
	"testing"
)

func TestOk(t *testing.T) {
	r := Ok(42)
	if !r.Ok() {
		t.Error("Expected Ok() to return true")
	}
	if r.Error() != nil {
		t.Error("Expected Error() to return nil")
	}
	val, err := r.Get()
	if err != nil {
		t.Errorf("Expected Get() to return no error, got %v", err)
	}
	if val != 42 {
		t.Errorf("Expected Get() to return 42, got %v", val)
	}
}

func TestErr(t *testing.T) {
	testErr := errors.New("test error")
	r := Err[int](testErr)
	if r.Ok() {
		t.Error("Expected Ok() to return false")
	}
	if r.Error() != testErr {
		t.Errorf("Expected Error() to return test error, got %v", r.Error())
	}
	val, err := r.Get()
	if err != testErr {
		t.Errorf("Expected Get() to return test error, got %v", err)
	}
	if val != 0 {
		t.Errorf("Expected Get() to return zero value, got %v", val)
	}
}

func TestErrf(t *testing.T) {
	r := Errf[string]("error: %s", "test")
	if r.Ok() {
		t.Error("Expected Ok() to return false")
	}
	if r.Error() == nil {
		t.Error("Expected Error() to return an error")
	}
	if r.Error().Error() != "error: test" {
		t.Errorf("Expected error message 'error: test', got %s", r.Error().Error())
	}
}

func TestTupleToResult(t *testing.T) {
	// Success case
	r := TupleToResult(42, nil)
	if !r.Ok() {
		t.Error("Expected Ok() to return true")
	}
	val, _ := r.Get()
	if val != 42 {
		t.Errorf("Expected value 42, got %v", val)
	}

	// Error case
	testErr := errors.New("test error")
	r = TupleToResult(0, testErr)
	if r.Ok() {
		t.Error("Expected Ok() to return false")
	}
	if r.Error() != testErr {
		t.Errorf("Expected error to match, got %v", r.Error())
	}
}

func TestTry(t *testing.T) {
	// Success case
	r := Try(func() (int, error) {
		return 42, nil
	})
	if !r.Ok() {
		t.Error("Expected Ok() to return true")
	}
	val, _ := r.Get()
	if val != 42 {
		t.Errorf("Expected value 42, got %v", val)
	}

	// Error case
	testErr := errors.New("test error")
	r = Try(func() (int, error) {
		return 0, testErr
	})
	if r.Ok() {
		t.Error("Expected Ok() to return false")
	}
	if r.Error() != testErr {
		t.Errorf("Expected error to match, got %v", r.Error())
	}
}

func TestResultMustGet(t *testing.T) {
	// Success case
	r := Ok(42)
	val := r.MustGet()
	if val != 42 {
		t.Errorf("Expected value 42, got %v", val)
	}

	// Panic case
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected MustGet to panic on error result")
		}
	}()
	testErr := errors.New("test error")
	errResult := Err[int](testErr)
	errResult.MustGet()
}

func TestResultOrElse(t *testing.T) {
	// Success case
	r := Ok(42)
	val := r.OrElse(100)
	if val != 42 {
		t.Errorf("Expected value 42, got %v", val)
	}

	// Error case
	r = Err[int](errors.New("test"))
	val = r.OrElse(100)
	if val != 100 {
		t.Errorf("Expected fallback value 100, got %v", val)
	}
}

func TestResultOrEmpty(t *testing.T) {
	// Success case
	r := Ok(42)
	val := r.OrEmpty()
	if val != 42 {
		t.Errorf("Expected value 42, got %v", val)
	}

	// Error case
	r = Err[int](errors.New("test"))
	val = r.OrEmpty()
	if val != 0 {
		t.Errorf("Expected zero value 0, got %v", val)
	}
}

func TestResultForEach(t *testing.T) {
	// Success case
	called := false
	r := Ok(42)
	r.ForEach(func(val int) {
		called = true
		if val != 42 {
			t.Errorf("Expected value 42, got %v", val)
		}
	})
	if !called {
		t.Error("Expected ForEach to be called")
	}

	// Error case
	called = false
	r = Err[int](errors.New("test"))
	r.ForEach(func(val int) {
		called = true
	})
	if called {
		t.Error("Expected ForEach not to be called on error")
	}
}

func TestResultMatch(t *testing.T) {
	// Success case
	r := Ok(42)
	result := r.Match(
		func(val int) (int, error) {
			return val * 2, nil
		},
		func(err error) (int, error) {
			return 0, err
		},
	)
	if !result.Ok() {
		t.Error("Expected result to be Ok")
	}
	val, _ := result.Get()
	if val != 84 {
		t.Errorf("Expected value 84, got %v", val)
	}

	// Error case
	testErr := errors.New("test error")
	r = Err[int](testErr)
	result = r.Match(
		func(val int) (int, error) {
			return val * 2, nil
		},
		func(err error) (int, error) {
			return 100, nil
		},
	)
	if !result.Ok() {
		t.Error("Expected result to be Ok after error handler")
	}
	val, _ = result.Get()
	if val != 100 {
		t.Errorf("Expected value 100, got %v", val)
	}
}

func TestResultMap(t *testing.T) {
	// Success case
	r := Ok(42)
	result := r.Map(func(val int) (int, error) {
		return val * 2, nil
	})
	if !result.Ok() {
		t.Error("Expected result to be Ok")
	}
	val, _ := result.Get()
	if val != 84 {
		t.Errorf("Expected value 84, got %v", val)
	}

	// Error case - original error
	testErr := errors.New("test error")
	r = Err[int](testErr)
	result = r.Map(func(val int) (int, error) {
		return val * 2, nil
	})
	if result.Ok() {
		t.Error("Expected result to be Err")
	}
	if result.Error() != testErr {
		t.Errorf("Expected error to match, got %v", result.Error())
	}

	// Error case - mapper error
	mapErr := errors.New("map error")
	r = Ok(42)
	result = r.Map(func(val int) (int, error) {
		return 0, mapErr
	})
	if result.Ok() {
		t.Error("Expected result to be Err")
	}
	if result.Error() != mapErr {
		t.Errorf("Expected error to match, got %v", result.Error())
	}
}

func TestResultMapErr(t *testing.T) {
	// Success case - no mapping
	r := Ok(42)
	result := r.MapErr(func(err error) (int, error) {
		return 100, nil
	})
	if !result.Ok() {
		t.Error("Expected result to be Ok")
	}
	val, _ := result.Get()
	if val != 42 {
		t.Errorf("Expected value 42, got %v", val)
	}

	// Error case - successful mapping
	testErr := errors.New("test error")
	r = Err[int](testErr)
	result = r.MapErr(func(err error) (int, error) {
		return 100, nil
	})
	if !result.Ok() {
		t.Error("Expected result to be Ok after MapErr")
	}
	val, _ = result.Get()
	if val != 100 {
		t.Errorf("Expected value 100, got %v", val)
	}

	// Error case - mapping returns error
	newErr := errors.New("new error")
	r = Err[int](testErr)
	result = r.MapErr(func(err error) (int, error) {
		return 0, newErr
	})
	if result.Ok() {
		t.Error("Expected result to be Err")
	}
	if result.Error() != newErr {
		t.Errorf("Expected error to match, got %v", result.Error())
	}
}

func TestResultFlatMap(t *testing.T) {
	// Success case
	r := Ok(42)
	result := r.FlatMap(func(val int) Result[int] {
		return Ok(val * 2)
	})
	if !result.Ok() {
		t.Error("Expected result to be Ok")
	}
	val, _ := result.Get()
	if val != 84 {
		t.Errorf("Expected value 84, got %v", val)
	}

	// Success to error case
	r = Ok(42)
	testErr := errors.New("test error")
	result = r.FlatMap(func(val int) Result[int] {
		return Err[int](testErr)
	})
	if result.Ok() {
		t.Error("Expected result to be Err")
	}
	if result.Error() != testErr {
		t.Errorf("Expected error to match, got %v", result.Error())
	}

	// Error case - no mapping
	r = Err[int](testErr)
	result = r.FlatMap(func(val int) Result[int] {
		return Ok(val * 2)
	})
	if result.Ok() {
		t.Error("Expected result to be Err")
	}
	if result.Error() != testErr {
		t.Errorf("Expected error to match, got %v", result.Error())
	}
}
