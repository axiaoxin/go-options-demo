// options by interface

package main

import "fmt"

// OptionItem fields are option items
type OptionItem struct {
	item1 int
	item2 string
}

// OptionApplyer interface
type OptionApplyer interface {
	apply(*OptionItem)
}

// OptionItemApplyFunc set new item function
type OptionItemApplyFunc func(*OptionItem)

func (f OptionItemApplyFunc) apply(opt *OptionItem) {
	f(opt)
}

// Item1Option set new item1
func Item1Option(item1 int) OptionApplyer {
	return OptionItemApplyFunc(func(opt *OptionItem) {
		opt.item1 = item1
	})
}

// Item2Option set new item2
func Item2Option(item2 string) OptionApplyer {
	return OptionItemApplyFunc(func(opt *OptionItem) {
		opt.item2 = item2
	})
}

// New mock init with options
func New(options ...OptionApplyer) *OptionItem {
	defaultOpt := &OptionItem{}
	for _, opt := range options {
		opt.apply(defaultOpt)
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
