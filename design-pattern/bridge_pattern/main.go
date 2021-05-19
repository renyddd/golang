package main

import "fmt"

// https://refactoringguru.cn/design-patterns/bridge/go/example

type computer interface {
	print()
	setPrinter(printer)
}

type mac struct {
	printer printer
}

func (m *mac) print() {
	fmt.Println("Print request for mac.")
	m.printer.printFile()
}

func (m *mac) setPrinter(p printer) {
	m.printer = p
}

type windows struct {
	printer printer
}

func (w *windows) print() {
	fmt.Println("Print request for windows.")
	w.printer.printFile()
}

type printer interface {
	printFile()
}

type epsonP struct {
}

func (p *epsonP) printFile() {
	fmt.Println("Printing by epson Printer.")
}

type hpP struct {
}

func (p *hpP) printFile() {
	fmt.Println("Printing by hp Printer.")
}

func main() {
	hp := &hpP{}
	macPc := &mac{}

	macPc.setPrinter(hp)
	macPc.print()

	epson := &epsonP{}
	macPc.setPrinter(epson)
	macPc.print()
}
