package foo

import "learn-git/bar"

type Foo struct {
}

func (f *Foo) Foo() {
	bar.Bar()
}