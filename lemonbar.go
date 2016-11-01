package main

type lemonbarFormatter struct {
	before, after string
}

var bareSingleton formatter

func NewLemonbarFormatter() formatter {
	f := new(lemonbarFormatter)
	f.after = ""
	f.before = ""
	return f
}

func (f *lemonbarFormatter) wrapOuter(before, after string) formatter {
	wrapped := f.clone()
	wrapped.before = before + f.before
	wrapped.after = f.after + after
	return wrapped
}
func (f *lemonbarFormatter) wrapInner(before, after string) formatter {
	wrapped := f.clone()
	wrapped.before = f.before + before
	wrapped.after = after + f.after
	return wrapped
}

func (f *lemonbarFormatter) WrapFgColor(color string) formatter {
	return f.wrapOuter("%{F"+color+"}", "%{F-}")
}
func (f *lemonbarFormatter) WrapBgColor(color string) formatter {
	return f.wrapOuter("%{B"+color+"}", "%{B-}")
}
func (f *lemonbarFormatter) AlignLeft() formatter {
	return f.wrapInner("%{l}", "")
}
func (f *lemonbarFormatter) AlignRight() formatter {
	return f.wrapInner("%{r}", "")
}
func (f *lemonbarFormatter) AlignCenter() formatter {
	return f.wrapInner("%{c}", "")
}
func (f *lemonbarFormatter) WrapUnderlineColor(color string) formatter {
	return f.wrapOuter("%{U"+color+"}", "%{U-}")
}
func (f *lemonbarFormatter) WrapUnderline() formatter {
	return f.wrapOuter("%{+u}", "%{-u}")
}
func (f *lemonbarFormatter) WrapOverline() formatter {
	return f.wrapOuter("%{+o}", "%{-o}")
}
func (f *lemonbarFormatter) WrapBold() formatter {
	return f
}
func (f *lemonbarFormatter) SetMonitor(monitor string) formatter {
	return f.wrapOuter("%{S"+monitor+"}", "")
}
func (f *lemonbarFormatter) GetDefaultColor() string {
	return "-"
}
func (f *lemonbarFormatter) AppendInner(text string) formatter {
	return f.wrapInner("", text)
}
func (f *lemonbarFormatter) PrependInner(text string) formatter {
	return f.wrapInner(text, "")
}
func (f *lemonbarFormatter) AppendOuter(text string) formatter {
	return f.wrapOuter("", text)
}
func (f *lemonbarFormatter) PrependOuter(text string) formatter {
	return f.wrapOuter(text, "")
}
func (f *lemonbarFormatter) clone() *lemonbarFormatter {
	clone := new(lemonbarFormatter)
	clone.before = f.before
	clone.after = f.after
	return clone
}
func (f *lemonbarFormatter) Bare() formatter {
	if bareSingleton == nil {
		bareSingleton = NewLemonbarFormatter()
	}
	return bareSingleton
}
func (f *lemonbarFormatter) Format(text string) string {
	return f.before + text + f.after
}
