package main

import (
	"fmt"
	"log"
	"strings"

	"flag"

	"github.com/docker/go-units"
	"github.com/jung-kurt/gofpdf"
)

var humanSize = flag.String("size", "1MB", "Set the size of the PDF to generate")

func main() {
	flag.Parse()

	size, err := units.FromHumanSize(*humanSize)
	if err != nil {
		log.Fatalf("Unable to parse size: %s because: %v", *humanSize, err)
	}

	gofpdf.SetDefaultCompression(false)

	if err := generatePdf(size); err != nil {
		fmt.Printf("Got error: %v\n", err)
	}

}

const fontSize = 16
const pdfOverhead = 6722

func generatePdf(size int64) error {
	pdf := gofpdf.New("P", "mm", "A4", "")

	fullPageOfText(pdf)

	pdf.Cell(0, 0, generateTextWithSize(size-pdfOverhead))

	filename := fmt.Sprintf("pdfs/%s.pdf", units.HumanSize(float64(size)))
	return pdf.OutputFileAndClose(filename)
}

func fullPageOfText(pdf gofpdf.Pdf) {
	pdf.AddPage()
	pdf.SetFont("Arial", "", fontSize)
	pdf.SetCompression(false)

	pdf.Write(4, strings.Repeat("a", 3894))
}

func generateTextWithSize(size int64) string {
	if size < 0 {
		return ""
	}

	return strings.Repeat("a", int(size))
}
