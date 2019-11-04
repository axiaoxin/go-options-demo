package main

import "fmt"

// OptionItem fields are option items
type OptionItem struct {
	item1 int
	item2 string
}

// OptionItemApplyFunc set new item function
type OptionItemApplyFunc func(*OptionItem)

// Item1Option set new item1
func Item1Option(item1 int) OptionItemApplyFunc {
	return func(o *OptionItem) {
		o.item1 = item1
	}
}

// Item2Option set new item2
func Item2Option(item2 string) OptionItemApplyFunc {
	return func(o *OptionItem) {
		o.item2 = item2
	}
}

// New mock a init with option items
func New(options ...OptionItemApplyFunc) *OptionItem {
	defaultOpt := &OptionItem{}
	for _, opt := range options {
		opt(defaultOpt)
	}
	return defaultOpt
}

func main() {
	opt := New()
	fmt.Printf("options f1:%d, f2:%s\n", opt.item1, opt.item2)

	opt = New(Item2Option("good"))
	fmt.Printf("options f1:%d, f2:%s\n", opt.item1, opt.item2)

	opt = New(Item1Option(999), Item2Option("xyz"))
	fmt.Printf("options f1:%d, f2:%s\n", opt.item1, opt.item2)
}
