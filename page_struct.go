package main

type InJson struct {
    Page1 struct {
        ID          string `json:"ID"`
        Imie        string `json:"Imie"`
        Nazwisko    string `json:"Nazwisko"`
        Ulica       string `json:"Ulica"`
        NrDomu      string `json:"NrDomu"`
        NrLokalu    string `json:"NrLokalu"`
        KodPocztowy string `json:"KodPocztowy"`
        Miejscowosc string `json:"Miejscowosc"`
        Panstwo     string `json:"Panstwo"`
        Telefon     string `json:"Telefon"`
        Pesel       []int  `json:"Pesel"`
    } `json:"page1"`

    Page2 struct {
        Okres                string `json:"Okres"`
        IdDziecka            string `json:"IdDziecka"`
        ImieDziecka          string `json:"ImieDziecka"`
        NazwiskoDziecka      string `json:"NazwiskoDziecka"`
        OkresInnyDomownik    string `json:"OkresInnyDomownik"`
        GodzinyOpieki        string `json:"GodzinyOpieki"`
        OrzeczenieOniepSprw  bool   `json:"OrzeczenieOniepSprw"`
        JestDomownik         bool   `json:"JestDomownik"`
        PracujeNaZmiany      bool   `json:"PracujeNaZmiany"`
        WspolneGospodarstwo  bool   `json:"WspolneGospodarstwo"`
        ZmienilesPlatnika    bool   `json:"ZmienilesPlatnika"`
        DniOpiekiDo8lat      int    `json:"DniOpiekiDo8lat"`
        DniOpiekiPonad14lat  int    `json:"DniOpiekiPonad14lat"`
        DniOpieki8_18lat     int    `json:"DniOpieki8_18lat"`
        PeselDziecka         []int  `json:"PeselDziecka"`
        DataUrodzeniaDziecka []int  `json:"DataUrodzeniaDziecka"`
    } `json:"page2"`

    Page3 struct {
        IdRodzica                string `json:"IdRodzica"`
        ImieRodzica              string `json:"ImieRodzica"`
        NazwiskoRodzica          string `json:"NazwiskoRodzica"`
        RodzicGodzinyOpieki      string `json:"RodzicGodzinyOpieki"`
        IdMalzonka               string `json:"IdMalzonka"`
        ImieMalzonka             string `json:"ImieMalzonka"`
        NazwiskoMalzonka         string `json:"NazwiskoMalzonka"`
        GodzinyOpieki            string `json:"GodzinyOpieki"`
        RodzicPracuje            bool   `json:"RodzicPracuje"`
        RodzicPracujeNaZmiany    bool   `json:"RodzicPracujeNaZmiany"`
        RodzicOtrzymalZasilek    bool   `json:"RodzicOtrzymalZasilek"`
        MalzonekPracuje          bool   `json:"MalzonekPracuje"`
        MalzonekPracujeNaZmiany  bool   `json:"MalzonekPracujeNaZmiany"`
        MalzonekOtrzymalZasilek  bool   `json:"MalzonekOtrzymalZasilek"`
        RodzicOpiekaDo8Lat       int    `json:"RodzicOpiekaDo8lat"`
        RodzicOpiekaPonad14Lat   int    `json:"RodzicOpiekaPonad14lat"`
        RodzicOpieka8_18Lat      int    `json:"RodzicOpieka8_18lat"`
        MalzonekOpiekaDo8Lat     int    `json:"MalzonekOpiekaDo8lat"`
        MalzonekOpiekaPonad14Lat int    `json:"MalzonekOpiekaPonad14lat"`
        MalzonekOpieka8_18Lat    int    `json:"MalzonekOpieka8_18lat"`
        PeselRodzica             []int  `json:"PeselRodzica"`
        PeselMalzonka            []int  `json:"PeselMalzonka"`
    }


    Page4 struct {
        IdInnego              string `json:"IdInnego"`
        ImieInnego            string `json:"ImieInnego"`
        NazwiskoInnego        string `json:"NazwiskoInnego"`
        IdInnego2             string `json:"IdInnego2"`
        ImieInnego2           string `json:"ImieInnego2"`
        NazwiskoInnego2       string `json:"NazwiskoInnego2"`
        OpiekaSprawowanaPrzez string `json:"OpiekaSprawowanaPrzez"`
        OpiekaNad             string `json:"OpiekaNad"`
        Uwagi                 string `json:"Uwagi"`
        OpiekaDo14            int    `json:"OpiekaDo14"`
        OpiekaPonad14         int    `json:"OpiekaPonad14"`
        Opieka8_18            int    `json:"Opieka8_18"`
        ZasilekZaXdni         int    `json:"ZasilekZaXdni"`
        PeselInnego           []int  `json:"PeselInnego"`
        PeselInnego2          []int  `json:"PeselInnego2"`
        NrRachunku            []int  `json:"NrRachunku"`
        Data                  []int  `json:"Data"`
    }
}
