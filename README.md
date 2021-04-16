# PDF generator

Generates a viewable PDF of a given size. The output is a PDF with a single page filled with "a"s and is the same size as the `-size` argument.

## Requirements

Requires a Go compliler or Docker and VSCode.

### VSCode + Docker 

VSCode can open the project in a "devcontainer" whcih contains a Go compliler. This removes the need to install Go on your machine. 

## Usage 

### Compile

```bash 
go build
```
That will produce an executable called `pdf_gen` which can then be run like so:

```bash
pdf_gen -size 5MB
```

That will produce a PDF which is 5MB in size. The output file will be `pdfs/5MB.pdf`.

### Compile and run

You can skip the compile step and run it directly like so:

```bash
go run main.go -size 5MB
```
