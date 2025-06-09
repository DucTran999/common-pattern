package bridge

type Document struct {
	Printer Printer
}

func (d *Document) Print() {
	d.Printer.PrintFile()
}
