package main

import "fmt"

func main() {

	// Maps için default tanım.
	myMap := make(map[int]string)
	fmt.Println(myMap)
	myMap[43] = "Foo"
	myMap[12] = "Bar"
	fmt.Println(myMap)

	// Farklı bir yöntem

	states := make(map[string]string)
	states["IST"] = "ISTANBUL"
	states["IZM"] = "IZMIR"
	states["ANK"] = "ANKARA"
	states["ANT"] = "ANTALYA"
	fmt.Println(states)

	// Şehir listesinde ANT anahtar adına sahip veriyi elde et

	antalya := states["ANT"]
	fmt.Println("Şehir=>", antalya)

	// ANK anahtar adına sahip veriyi silmek
	delete(states, "ANK")
	fmt.Println(states)

	// Değer ekleme
	states["ERC"] = "Erzurum"

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	// states map nesnesinin elemanı adadince kapasiteli bir key listesi oluştur
	keys := make([]string, len(states))
	i := 0
	for k, v := range states {
		keys[i] = k
		fmt.Println(v)
		i++
	}
	fmt.Println(keys)

	// Key Listesinde ki index değerlerine göre states nesnesinde bulunan şehirleri yazdır
	for i := range keys {
		fmt.Println(states[keys[i]])
	}

}
