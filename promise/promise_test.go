package promise

import (
	"context"
	"testing"
	"time"
)

func TestPromise_Await1(t *testing.T) {

	ctx, _ := context.WithTimeout(context.Background(), 500*time.Millisecond)

	p := Async(ctx, func(_ context.Context) (string, error) {
		time.Sleep(time.Millisecond)
		return "Hello", nil
	})

	s, err := p.Await()
	if err != nil {
		t.Logf("expected no error, got %v", err)
		t.Fail()
	}
	if s != "Hello" {
		t.Logf("expected return Hello, got %v", s)
		t.Fail()
	}
}

func TestPromise_Await2(t *testing.T) {

	ctx, _ := context.WithTimeout(context.Background(), 500*time.Millisecond)

	p := Async(ctx, func(_ context.Context) (string, error) {
		time.Sleep(time.Millisecond)
		return "Hello", nil
	})

	s, err := Await(p)
	if err != nil {
		t.Logf("expected no error, got %v", err)
		t.Fail()
	}
	if s != "Hello" {
		t.Logf("expected return Hello, got %v", s)
		t.Fail()
	}
}

func TestPromise_AwaitAll(t *testing.T) {

	ctx, _ := context.WithTimeout(context.Background(), 500*time.Millisecond)

	p1 := Async(ctx, func(_ context.Context) (string, error) {
		time.Sleep(time.Millisecond)
		return "Hello 1", nil
	})
	p2 := Async(ctx, func(_ context.Context) (string, error) {
		time.Sleep(time.Millisecond)
		return "Hello 2", nil
	})

	err := AwaitAll(ctx, p1, p2)

	if err != nil {
		t.Logf("expected no error, got %v", err)
		t.Fail()
	}

	r1, e1 := p1.Await()
	if e1 != nil {
		t.Logf("expected no error, got %v", e1)
		t.Fail()
	}
	if r1 != "Hello 1" {
		t.Logf("expected return Hello 1, got %v", r1)
		t.Fail()
	}
	r2, e2 := p2.Await()
	if e2 != nil {
		t.Logf("expected no error, got %v", e2)
		t.Fail()
	}
	if r2 != "Hello 2" {
		t.Logf("expected return Hello 2, got %v", r2)
		t.Fail()
	}
}

func TestPromise_Await_Then(t *testing.T) {

	ctx, _ := context.WithTimeout(context.Background(), 500*time.Millisecond)

	p := Async(ctx, func(_ context.Context) (string, error) {
		time.Sleep(time.Millisecond)
		return "Hello", nil
	}).Then(func(s string) {
		if s != "Hello" {
			t.Logf("expected return Hello, got %v", s)
			t.Fail()
		}
	})

	s, err := Await(p)
	if err != nil {
		t.Logf("expected no error, got %v", err)
		t.Fail()
	}
	if s != "Hello" {
		t.Logf("expected return Hello, got %v", s)
		t.Fail()
	}
}
