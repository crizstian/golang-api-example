package libs

import (
	"crypto/rand"
	"math/big"
	randmin "math/rand"
	"strconv"
	"strings"
	"time"

	hashids "github.com/speps/go-hashids"
)

type Utils struct {
	Pass string
}

func (w Utils) RandString(n int) string {
	const alphanum = "0123456789abcdefghijklmnopqrstuvwxyz"
	symbols := big.NewInt(int64(len(alphanum)))
	states := big.NewInt(0)
	states.Exp(symbols, big.NewInt(int64(n)), nil)
	r, err := rand.Int(rand.Reader, states)
	if err != nil {
		panic(err)
	}
	var bytes = make([]byte, n)
	r2 := big.NewInt(0)
	symbol := big.NewInt(0)
	for i := range bytes {
		r2.DivMod(r, symbols, symbol)
		r, r2 = r2, r
		bytes[i] = alphanum[symbol.Int64()]
	}
	return string(bytes)
}

type BitsData struct {
	Bytes     float64 `json: "bytes"`
	Kilobytes float64 `json: "kilobytes"`
	Megabytes float64 `json: "megabytes"`
	Gigabytes float64 `json: "gigabytes"`
}

func BitsConversion(bits float64) interface{} {
	d := new(BitsData)

	_byte := 8.0
	d.Bytes = bits / _byte
	d.Kilobytes = d.Bytes / 1000.0
	d.Megabytes = d.Kilobytes / 1000.0
	d.Gigabytes = d.Megabytes / 1000.0

	return d
}

// Generate uniq random string
func Entry() (string, error) {

	hd := hashids.NewData()
	hd.Salt = "a9ws1d4frg6a1s2d3f4g5zax8c9vf1g"
	hd.Alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	hd.MinLength = 15

	h, err := hashids.NewWithData(hd)
	timeint := int(time.Now().UnixNano())
	randint := randmin.Intn(9999)

	e, err := h.Encode([]int{timeint, randint})

	if err != nil {
		return "", err
	}

	return e, nil
}

// devuelve un int con el anio y mes en el que se creo una cuenta de usuario
func UserCreatedInt(created_at string) int {
	date := strings.Split(created_at, "-")
	ym := date[0] + date[1]
	yms, _ := strconv.Atoi(ym)
	return yms
}

// calcula el total de dias transcurridos desde que se creo la cuenta al dia de hoy
func DaysTillToday(Date time.Time) int {
	DateParse, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", Date.String())
	Duration := time.Now().Sub(DateParse)
	return int(Duration.Hours() / 24)
}
