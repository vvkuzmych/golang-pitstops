package main

import "fmt"

// Document represents the structure of a document.
type Document struct {
	Name string
}

// Machine defines an interface with all functionalities (violates ISP if not all functions are needed).
type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

// MultiFunctionPrinter implements the full Machine interface.
type MultiFunctionPrinter struct{}

func (m MultiFunctionPrinter) Print(d Document) {
	fmt.Println("Printing document:", d.Name)
}

func (m MultiFunctionPrinter) Fax(d Document) {
	fmt.Println("Faxing document:", d.Name)
}

func (m MultiFunctionPrinter) Scan(d Document) {
	fmt.Println("Scanning document:", d.Name)
}

// OldFashionedPrinter implements only Print but other methods are not supported.
type OldFashionedPrinter struct{}

func (o OldFashionedPrinter) Print(d Document) {
	fmt.Println("Printing document:", d.Name)
}

func (o OldFashionedPrinter) Fax(d Document) {
	panic("operation not supported")
}

func (o OldFashionedPrinter) Scan(d Document) {
	panic("operation not supported")
}

// Following ISP: Split into separate interfaces.
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type FaxMachine interface {
	Fax(d Document)
}

// MyPrinter implements only the Printer interface.
type MyPrinter struct{}

func (m MyPrinter) Print(d Document) {
	fmt.Println("Printing document:", d.Name)
}

// Photocopier implements both Printer and Scanner interfaces.
type Photocopier struct{}

func (p Photocopier) Print(d Document) {
	fmt.Println("Printing document in Photocopier:", d.Name)
}

func (p Photocopier) Scan(d Document) {
	fmt.Println("Scanning document in Photocopier:", d.Name)
}

// MultiFunctionDevice combines Printer and Scanner interfaces.
type MultiFunctionDevice interface {
	Printer
	Scanner
}

// MultiFunctionMachine acts as a decorator for MultiFunctionDevice.
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

// Demonstrates the principle in action.
func interfacePrinciple() {
	doc := Document{Name: "Report.pdf"}

	// Using MultiFunctionPrinter
	mfp := MultiFunctionPrinter{}
	mfp.Print(doc)
	mfp.Fax(doc)
	mfp.Scan(doc)

	// Using OldFashionedPrinter
	ofp := OldFashionedPrinter{}
	ofp.Print(doc)
	// Uncommenting these lines will panic
	// ofp.Fax(doc)
	// ofp.Scan(doc)

	// Using MyPrinter
	myPrinter := MyPrinter{}
	myPrinter.Print(doc)

	// Using Photocopier
	photocopier := Photocopier{}
	photocopier.Print(doc)
	photocopier.Scan(doc)

	// Using MultiFunctionMachine
	mfm := MultiFunctionMachine{
		printer: myPrinter,
		scanner: photocopier,
	}
	mfm.Print(doc)
	mfm.Scan(doc)
}

func main() {
	interfacePrinciple()
}
