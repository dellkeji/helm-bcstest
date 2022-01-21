package printer

import (
	"fmt"
	"io"

	"github.com/fatih/color"
)

// ColorPrinter 颜色输出
type ColorPrinter struct {
	Success   *color.Color
	SuccessBg *color.Color
	Warning   *color.Color
	WarningBg *color.Color
	Failure   *color.Color
	FailureBg *color.Color
	Highlight *color.Color
	Faint     *color.Color
}

// Printer 输出器
type Printer struct {
	Writer io.Writer
	Colors ColorPrinter
}

// NewPrinter 创建 printer
func NewPrinter(writer io.Writer) *Printer {
	p := &Printer{
		Writer: writer,
		Colors: ColorPrinter{
			Success:   color.New(color.FgGreen),
			SuccessBg: color.New(color.BgGreen, color.BgBlack),
			Failure:   color.New(color.FgRed),
			FailureBg: color.New(color.BgRed, color.FgWhite),
			Warning:   color.New(color.FgYellow),
			WarningBg: color.New(color.BgYellow, color.FgBlack),
			Highlight: color.New(color.Bold),
		},
	}
	return p
}

// Print 输出到 writer, times 用以 tab 空行
func (p *Printer) Print(content string, times int) {
	var indent string
	for i := 0; i < times; i++ {
		indent += "\t"
	}
	fmt.Fprintln(p.Writer, indent+content)
}

// SuccessPrint 成功输出
func (p *Printer) SuccessPrint(format string, a ...interface{}) string {
	return p.Colors.Success.Sprintf(format, a...)
}

// SuccessBgPrint 带有背景色的输出
func (p *Printer) SuccessBgPrint(format string, a ...interface{}) string {
	return p.Colors.SuccessBg.Sprintf(format, a...)
}

// FailurePrint 失败的输出
func (p *Printer) FailurePrint(format string, a ...interface{}) string {
	return p.Colors.Failure.Sprintf(format, a...)
}

// FailureBgPrint 带有背景色的失败输出
func (p *Printer) FailureBgPrint(format string, a ...interface{}) string {
	return p.Colors.FailureBg.Sprintf(format, a...)
}

// WarningPrint 警告输出
func (p *Printer) WarningPrint(format string, a ...interface{}) string {
	return p.Colors.Warning.Sprintf(format, a...)
}

// WarningBgPrint 带有背景的警告输出
func (p *Printer) WarningBgPrint(format string, a ...interface{}) string {
	return p.Colors.Warning.Sprintf(format, a...)
}

// HighlightPrint 高亮输出
func (p *Printer) HighlightPrint(format string, a ...interface{}) string {
	return p.Colors.Highlight.Sprintf(format, a...)
}
