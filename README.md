# Wymagania
    - mupdf
    - go-bindata
    - jung-kurt/gofpdf

# Budowanie

```bash
$ make reqs
$ make getZ15
$ make exctract
$ make bundle
$ make build
```

# Użycie

- Przygotuj plik `JSON` na podstawie wzoru `indata.json'.
- Uruchom:

```bash
$ cat plik.json | ./z15 > output.pdf
```
