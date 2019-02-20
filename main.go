package main

import (
    "fmt"
	"log"
    "os"
    "io/ioutil"
    "bytes"
    "strconv"
    "encoding/json"
	"github.com/jung-kurt/gofpdf"
)

const (
    border = "0"
    fSize = 17
    fSizeS = 12
    font = "Courier"
	H = 6.32
	L1 = 84.3
    L2 = 141.0
    W1 = 110.5
	W2 = 25
)


func embedImage(pdf *gofpdf.Fpdf, imageName string) {
    imgOpt := gofpdf.ImageOptions{"png", true, false}
	data, err := Asset("data/" + imageName)

	if err != nil {
		log.Fatal("Cannot extract document page!", err)
	}

    pdf.RegisterImageOptionsReader(imageName, imgOpt, bytes.NewReader(data))
}

func cell(pdf *gofpdf.Fpdf, x float64, y float64, w float64, h float64, text string) {
	pdf.SetXY(x, y)
	pdf.CellFormat(w, h, text, border, 0, "L", false, 0, "")
}

func cellc(pdf *gofpdf.Fpdf, x float64, y float64, w float64, h float64, text string) {
	pdf.SetXY(x, y)
	pdf.CellFormat(w, h, text, border, 0, "C", false, 0, "")
}

func printArray(pdf *gofpdf.Fpdf, array []int, start float64) {
	pdf.SetFont(font, "", fSize)
    for x:= 0; x < len(array); x++ {
	    cell(pdf, L1 + float64(x) * H, start, H, H, strconv.Itoa(array[x]))
    }
}

func putX(pdf *gofpdf.Fpdf, X, Y, Z float64, cond bool) {
    if cond {
        cellc(pdf, X, Z, H, H, "X")
    } else {
        cellc(pdf, Y, Z, H, H, "X")
    }
}

func putXifMoreThan0(pdf *gofpdf.Fpdf, X1, Y1, X2, Y2 float64, cond int) {
    if (cond > 0) {
	    pdf.SetFont(font, "", fSize)
        cellc(pdf, X1, Y1, H, H, "X")
	    pdf.SetFont(font, "", fSizeS)
        cellc(pdf, X2, Y2, 15, H-3, strconv.Itoa(cond))
    }
}

func putIfMoreThan(pdf *gofpdf.Fpdf, X, Y float64, data int) {
    if data > 0 {
        pdf.SetFont(font, "", fSizeS)
        cellc(pdf, X, Y, 15, H, strconv.Itoa(data))
    }
}

func prepBasePage(pdf *gofpdf.Fpdf, bkg string) {
	pdf.AddPage()
	pdf.Image(bkg, 0, 0, 210, 297, false, "", 0, "")
	pdf.SetFont(font, "", fSize)
}

func preparePages(pdf *gofpdf.Fpdf, data InJson) {
    const (
        L2 = 21.2
        L3 = 147.2
        L4 = 172.3
    )

    prepBasePage(pdf, "page1.png")
	cell(pdf, L1, 182.85, W1, H, data.Page1.ID)
	cell(pdf, L1, 195.55, W1, H, data.Page1.Imie)
	cell(pdf, L1, 205.00, W1, H, data.Page1.Nazwisko)
	cell(pdf, L1, 214.40, W1, H, data.Page1.Ulica)
	cell(pdf, L1, 223.80, W2, H, data.Page1.NrDomu)
	cell(pdf, L1, 233.20, W2, H, data.Page1.KodPocztowy)
	cell(pdf, L3-5, 223.80, W2, H, data.Page1.NrLokalu)
	cell(pdf, L1, 242.80, W1, H, data.Page1.Miejscowosc)
	cell(pdf, L1, 252.20, W1, H, data.Page1.Panstwo)
	cell(pdf, L1, 264.60, W1, H, data.Page1.Telefon)
    printArray(pdf, data.Page1.Pesel, 173.30)

    prepBasePage(pdf, "page2.png")
	cell(pdf, 15.1, 40.95, 180, 10.5, data.Page2.Okres)
	cell(pdf, L1, 80.5, W1, H, data.Page2.IdDziecka)
	cell(pdf, L1, 93.1, W1, H, data.Page2.ImieDziecka)
	cell(pdf, L1, 102.5, W1, H, data.Page2.NazwiskoDziecka)
    cell(pdf, L2, 163.7, 173.7, 12.5, data.Page2.OkresInnyDomownik)
    cell(pdf, L2, 192.1, 173.7, 12.5, data.Page2.GodzinyOpieki)
    putX(pdf, L3, L4, 130.9, data.Page2.OrzeczenieOniepSprw)
    putX(pdf, L3, L4, 154.2, data.Page2.JestDomownik)
    putX(pdf, L3, L4, 182.5, data.Page2.PracujeNaZmiany)
    putX(pdf, L3, L4, 210.0, data.Page2.WspolneGospodarstwo)
    putX(pdf, L3, L4, 223.7, data.Page2.ZmienilesPlatnika)
    printArray(pdf, data.Page2.PeselDziecka, 71.00)
    printArray(pdf, data.Page2.DataUrodzeniaDziecka, 112.00)
    putXifMoreThan0(pdf, L2, 242.6, 64.7, 245.2, data.Page2.DniOpiekiDo8lat)
    putXifMoreThan0(pdf, L2, 251.8, 151.2, 252.6, data.Page2.DniOpiekiPonad14lat)
    putXifMoreThan0(pdf, L2, 261.3, 36.5, 266.0, data.Page2.DniOpieki8_18lat)

    prepBasePage(pdf, "page3.png")
	cell(pdf, L1, 178.9, W1, H, data.Page3.IdMalzonka)
	cell(pdf, L1, 201.3, W1, H, data.Page3.NazwiskoMalzonka)
	cell(pdf, L1, 53.55, W1, H, data.Page3.IdRodzica)
	cell(pdf, L1, 75.6, W1, H, data.Page3.NazwiskoRodzica)
    cell(pdf, L2, 220.0 , 173.6, 10.7, data.Page3.GodzinyOpieki)
    cell(pdf, L2, 94.5, 173.5, 10.5, data.Page3.RodzicGodzinyOpieki)
    cell(pdf, L1, 191.9, W1, H, data.Page3.ImieMalzonka)
    cell(pdf, L1, 66.0, W1, H, data.Page3.ImieRodzica)
    printArray(pdf, data.Page3.PeselRodzica, 43.90)
    printArray(pdf, data.Page3.PeselMalzonka, 169.50)
    putX(pdf, L2, 122.0, 111.3, data.Page3.RodzicOtrzymalZasilek)
    putX(pdf, 122.1, 140.8, 237.0, data.Page3.MalzonekOtrzymalZasilek)
    putX(pdf, 153.7, 172.5, 210.6, data.Page3.MalzonekPracujeNaZmiany)
    putX(pdf, 160.0, 178.6, 84.8, data.Page3.RodzicPracujeNaZmiany)
    putX(pdf, 59.2, 78.1, 210.6, data.Page3.MalzonekPracuje)
    putX(pdf, 65.4, 84.1, 84.8, data.Page3.RodzicPracuje)
    putXifMoreThan0(pdf, L2, 123.8, 64.7, 126.5, data.Page3.RodzicOpiekaDo8Lat)
    putXifMoreThan0(pdf, L2, 133.5, 151.2, 134.3, data.Page3.RodzicOpiekaPonad14Lat)
    putXifMoreThan0(pdf, L2, 142.8, 36.5, 147.5, data.Page3.RodzicOpieka8_18Lat)
    putXifMoreThan0(pdf, L2, 249.8, 64.7, 252.3, data.Page3.MalzonekOpiekaDo8Lat)
    putXifMoreThan0(pdf, L2, 258.8, 151.2, 260.2, data.Page3.MalzonekOpiekaPonad14Lat)
    putXifMoreThan0(pdf, L2, 268.7, 36.5, 273.3, data.Page3.MalzonekOpieka8_18Lat)

    prepBasePage(pdf, "page4.png")
    printArray(pdf, data.Page4.PeselInnego, 53.30);
    printArray(pdf, data.Page4.PeselInnego2, 110.50)
    pdf.SetFont(font, "", fSizeS)
	cell(pdf, 60.0, 174.0, 69, H-2, data.Page4.OpiekaSprawowanaPrzez)
	cell(pdf, 134.0, 174.0, 59, H-2, data.Page4.OpiekaNad)
    pdf.SetFont(font, "", fSize)
	cell(pdf, L1, 62.8, W1, H, data.Page4.IdInnego)
    cell(pdf, L1, 75.0, W1, H, data.Page4.ImieInnego)
	cell(pdf, L1, 84.5, W1, H, data.Page4.NazwiskoInnego)
	cell(pdf, L1, 119.9, W1, H, data.Page4.IdInnego2)
    cell(pdf, L1, 132.2, W1, H, data.Page4.ImieInnego2)
	cell(pdf, L1, 141.6, W1, H, data.Page4.NazwiskoInnego2)
	cell(pdf, 15.0, 190.5, 180, 15, data.Page4.Uwagi)
    putIfMoreThan(pdf, 65.0, 90.3, data.Page4.ZasilekZaXdni)
    putIfMoreThan(pdf, 23.0, 155.6, data.Page4.OpiekaDo14)
    putIfMoreThan(pdf, 137.5, 159.6, data.Page4.OpiekaPonad14)
    putIfMoreThan(pdf, 163.0, 167.5, data.Page4.Opieka8_18)

    for x:=0; x < len(data.Page4.NrRachunku); x++ {
        cellc(pdf, 15 + float64(x) * H*0.999, 217.10, H, H,
              strconv.Itoa(data.Page4.NrRachunku[x]))
    }

    for x:=0; x < len(data.Page4.Data); x++ {
        cellc(pdf, 27.8 + float64(x) * H*0.999, 247.60, H, H,
              strconv.Itoa(data.Page4.Data[x]))
    }
}

func makePdf(data InJson) {
	pdf := gofpdf.New("P", "mm", "A4", "")

    for i:=1; i<5; i++ {
        embedImage(pdf, fmt.Sprintf("page%d.png", i))
    }

	preparePages(pdf, data)
    err := pdf.Output(os.Stdout)
	if err != nil {
		log.Fatal("Cannot create PDF file!", err)
	}
}


func main() {
    var data InJson
    temp, _ := ioutil.ReadAll(os.Stdin)

    err := json.Unmarshal(temp, &data)
    if err != nil {
        log.Fatal("There was an error reading JSON data: ", err)
    }

	makePdf(data)
}
