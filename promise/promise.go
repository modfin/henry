package promise

import "context"

type Completable interface {
	Done() <-chan struct{}
}

func Await[A any](p *Promise[A]) (A, error) {
	return p.Await()
}

func AwaitAll(ctx context.Context, promises ...Completable) error {
	for _, p := range promises {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-p.Done():
		}
	}
	return nil
}

type Promise[A any] struct {
	done   chan struct{}
	ctx    context.Context
	result A
	err    error

	thens []func(A)
	errs  []func(error)
	all   []func(A, error)
}

func Async[A any](ctx context.Context, exec func(context.Context) (A, error)) *Promise[A] {
	var p = &Promise[A]{
		done: make(chan struct{}),
		ctx:  ctx,
	}

	go func() {
		p.result, p.err = exec(p.ctx)

		if p.err == nil {
			for _, t := range p.thens {
				t(p.result)
			}
		}
		if p.err != nil {
			for _, t := range p.errs {
				t(p.err)
			}
		}
		if p.all != nil {
			for _, t := range p.all {
				t(p.result, p.err)
			}
		}

		close(p.done)
	}()

	return p
}

func (p *Promise[A]) Await() (A, error) {
	select {
	case <-p.ctx.Done():
		return p.result, p.ctx.Err()
	case <-p.done:
		return p.result, p.err
	}
}

func (p *Promise[A]) Then(onSuccess func(A)) *Promise[A] {
	p.thens = append(p.thens, onSuccess)
	return p
}
func (p *Promise[A]) Error(onError func(error)) *Promise[A] {
	p.errs = append(p.errs, onError)
	return p
}
func (p *Promise[A]) Finally(onAll func(A, error)) *Promise[A] {
	p.all = append(p.all, onAll)
	return p
}

func (p *Promise[A]) Done() <-chan struct{} {
	return p.done
}
