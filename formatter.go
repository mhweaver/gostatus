package main

type formatter interface {
	GetDefaultColor() string
	WrapFgColor(color string) formatter
	WrapBgColor(color string) formatter
	WrapUnderlineColor(color string) formatter

	AlignLeft() formatter
	AlignRight() formatter
	AlignCenter() formatter

	WrapUnderline() formatter
	WrapOverline() formatter

	WrapBold() formatter

	AppendInner(text string) formatter
	PrependInner(text string) formatter
	AppendOuter(text string) formatter
	PrependOuter(text string) formatter
	Bare() formatter
	SetMonitor(monitor string) formatter

	Format(text string) string
}
