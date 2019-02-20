URL="http://www.zus.pl/documents/10182/788036/Z-15A+%2B+Informacja_stycze%C5%84+2018.pdf/478d8394-3ecb-444f-b4c0-50608695bddf"

getZ15:
	wget -v -O z15.pdf $(URL)

reqs:
	go get -u -v github.com/jung-kurt/gofpdf/...
	go get -u github.com/jteeuwen/go-bindata/...

extract:
	mkdir -p data
	mutool draw -o "page%d.png" -F png -r 600 z15.pdf 1,2,3,4
	mv page*.png data/.

bundle:
	go-bindata data/

build:
	go build -o z15 *.go

clean:
	rm -f z15.pdf

test: build
	cat indata.json | ./z15 > test.pdf
