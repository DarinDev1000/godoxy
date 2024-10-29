package error

import (
	"fmt"
	"sync"
)

type Builder struct {
	about string
	errs  []error
	sync.Mutex
}

func NewBuilder(about string) *Builder {
	return &Builder{about: about}
}

func (b *Builder) About() string {
	if !b.HasError() {
		return ""
	}
	return b.about
}

//go:inline
func (b *Builder) HasError() bool {
	return len(b.errs) > 0
}

func (b *Builder) Error() Error {
	if !b.HasError() {
		return nil
	}
	if len(b.errs) == 1 {
		return From(b.errs[0])
	}
	return &nestedError{Err: New(b.about), Extras: b.errs}
}

func (b *Builder) String() string {
	if !b.HasError() {
		return ""
	}
	return (&nestedError{Err: New(b.about), Extras: b.errs}).Error()
}

// Add adds an error to the Builder.
//
// adding nil is no-op,
func (b *Builder) Add(err error) *Builder {
	if err == nil {
		return b
	}

	b.Lock()
	defer b.Unlock()

	switch err := err.(type) {
	case *baseError:
		b.errs = append(b.errs, err.Err)
	case *nestedError:
		if err.Err == nil {
			b.errs = append(b.errs, err.Extras...)
		} else {
			b.errs = append(b.errs, err)
		}
	default:
		b.errs = append(b.errs, err)
	}

	return b
}

func (b *Builder) Adds(err string) *Builder {
	b.Lock()
	defer b.Unlock()
	b.errs = append(b.errs, newError(err))
	return b
}

func (b *Builder) Addf(format string, args ...any) *Builder {
	if len(args) > 0 {
		b.Lock()
		defer b.Unlock()
		b.errs = append(b.errs, fmt.Errorf(format, args...))
	} else {
		b.Adds(format)
	}

	return b
}

func (b *Builder) AddRange(errs ...error) *Builder {
	b.Lock()
	defer b.Unlock()

	for _, err := range errs {
		if err != nil {
			b.errs = append(b.errs, err)
		}
	}

	return b
}
