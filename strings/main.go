package main

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())

type Option struct {
	length    int
	uppcase   bool
	lowerCase bool
	numbers   bool
	specials  bool
}

const _charset = "."
const _charsetLowerCase = "qwertyuıopğüasdfghjklişzxcvbnmöç"
const _charsetUpperCase = "QWERTYUIOPĞÜİŞLKJHGGFDSAZXCVBNMMÖÇ"
const _charsetNumbers = "123456789"
const _charsetSpecial = "*-,./&%+^'!"

func GeneratePassword(length int) string {
	x := make([]byte, length)
	for i := range x {
		x[i] = _charset[source.Int63()&int64(len(_charset))]
	}
	return string(x)
}

func GeneratePasswordAdvanced(opt Option) (string, error) {
	x := make([]byte, opt.length)
	charset2 := "."

	if opt.uppcase {
		charset2 += _charsetUpperCase
	}
	if opt.lowerCase {
		charset2 += _charsetLowerCase
	}
	if opt.numbers {
		charset2 += _charsetNumbers
	}
	if opt.specials {
		charset2 += _charsetSpecial
	}
	if charset2 == "." {
		return "Non Generated Password", errors.New("Type Casting Error!")
	}

	for i := range x {
		x[i] = charset2[source.Int63()%int64(len(charset2))]
	}

	return string(x), nil

}

func main() {

	// Az sayıda string birleştirirken + operatörü iyidir fakat çok fazla string birleştiriliyorsa bu bellek yönetimi için kötüdür.
	str1 := "abc"
	str2 := "def"
	result := str1 + str2
	fmt.Println(result)

	// Çok fazla string birleştirme için buffer yöntemi iyi yöntemlerden birisidir.

	var x bytes.Buffer
	for i := 0; i < 100; i++ {
		x.WriteString(randomString())
	}
	//fmt.Println(x.String())

	// String builder işlemleri

	builder := strings.Builder{}

	for i := 0; i < 10; i++ {
		builder.WriteString("Data ")
	}

	result2 := builder.String()
	// daha önce hiç kullanılmadı hatasını gidermek için yazdık ciddiye almayın.
	result2 = result2
	//fmt.Println(result2)

	// rastgele işlemler
	randomPassword := GeneratePassword(15)
	randomPassword = randomPassword
	//fmt.Println(randomPassword)

	// Gelişmiş rastgele password (String oluşturma)
	passwordAdvanced, err := GeneratePasswordAdvanced(Option{
		length: 15, uppcase: true, lowerCase: true, specials: true, numbers: true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(passwordAdvanced)

}

func randomString() string {
	return "random string"
}
