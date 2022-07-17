package foo

import "learn-git/bar"

type Foo struct {
}

func NewFoo() *Foo {
	return &Foo{}
}

func (f *Foo) Foo() {
	bar.Bar()
}