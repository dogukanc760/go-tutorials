package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	newFile  *os.File
	fileInfo *os.FileInfo
	err      error
)

func main() {
	// Temel olarak dosya oluşturma

	newFile, err = os.Create("demo.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Dosya bilgisini döndürür.
	// Eğer dosya yoksa hata döner!
	fileInfo, err := os.Stat("demo.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("Can't found existing file")
		}
	}
	if err != nil {
		log.Fatal(err)
	}

	// Dosyanın var olup olmadığını kontrol etmek

	fmt.Println(fileInfo)
	// Dosya boyutu
	fmt.Printf("Size in bytes: %d\n", fileInfo.Size())
	// Dosya permleri
	fmt.Printf("Permissions: %d\n", fileInfo.Mode())
	// Son Güncelleme Tarihi
	fmt.Printf("Last Modified: %d\n", fileInfo.ModTime())
	// Bu Path bir dizin mi dosya mı ?
	fmt.Printf("Is Dictionary? %d\n", fileInfo.IsDir())
	// Sistem ARayüz Tipi
	fmt.Printf("System İnterface Type: %d\n", fileInfo.Sys())
	// Sistem İnfo
	fmt.Printf("System Info: %d\n", fileInfo.Sys())

	// Dosyayı yeniden adlandırma ve taşımak

	// originalPath := "demo.txt"
	// newPath := "./moved/test.txt"
	// errs := os.Rename(originalPath, newPath)
	// if err != nil {
	// 	log.Fatal(errs)
	// }

	// ikinci parametre dosya açılış amacını ayarlar
	// üçüncü parametre dosya izinlerini belirler.
	file, errF := os.OpenFile("demo.txt", os.O_APPEND, 0666)
	// başına defer yazmamızın sebebi scope sonunda yazmasak bile bu kod scope sonunda dosyayı kapatacak threadi tetikler.
	defer file.Close()
	if errF != nil {
		log.Fatal(errF)
	}

	file2, err2 := os.OpenFile("demo.txt", os.O_RDONLY, 0666)
	// başına defer yazmamızın sebebi scope sonunda yazmasak bile bu kod scope sonunda dosyayı kapatacak threadi tetikler.
	defer file2.Close()
	if err2 != nil {
		if os.IsPermission(err2) {
			log.Fatal("Dosya için izne sahip değilsiniz!")
		}
	}

	originalFile, err3 := os.Open("demo.txt")
	if err3 != nil {
		log.Fatal(err3)
	}

	defer originalFile.Close()

	// Yeni bir dosya oluştur
	newFile, err4 := os.Create("demo-copy.txt")
	if err4 != nil {
		log.Fatal(err4)
	}
	defer newFile.Close()

	// Byte'ları kaynaktan hedefe kopyala
	bytesWritten, err5 := io.Copy(newFile, originalFile)
	if err5 != nil {
		log.Fatal(err5)
	}
	log.Printf("Copied bytes are: %v", bytesWritten)

	// Dosya içeriğini işle - yani commit et
	// Belleği diske boşalt
	err6 := newFile.Sync()
	if err6 != nil {
		log.Fatal(err6)
	}

	// Bir dosyaya veri yazma
	// Birden fazla dosya işlemi yapılabilir, örneğin aşağıda eğer dosya yoksa önce dosyayı create eder.
	file4, err7 := os.OpenFile("demo.txt", os.O_WRONLY, 0666)
	if err7 != nil {
		log.Fatal(err7)
	}
	defer file4.Close()

	byteSlice := []byte("Bu veri yazıldı! \n")
	bytesWritten, err8 := file4.Write(byteSlice)
	if err8 != nil {
		log.Fatal(err8)
	}
	log.Printf("Wrote %d bytes \n", bytesWritten)

	// Geçici dosyalar ve dizinler
	tempDirPath, err9 := ioutil.TempDir("", "geciciKlasor")
	if err9 != nil {
		log.Fatal(err9)
	}
	fmt.Println("Gecici klasor olusturuldu!=>", tempDirPath)

	// Geçici dosya oluşturuldu!
	tempFile, err := ioutil.TempFile(tempDirPath, "geciciDosya.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Geçici dosya oluşturuldu!=>", tempFile.Name())

	// Dosya Kapatma İşlemleri
	err = tempFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	// Geçici dosyayı Silme

	err = os.Remove(tempFile.Name())
	if err != nil {
		log.Fatal(err)
	}

	// Geçici dizini silme
	err = os.Remove(tempDirPath)
	if err != nil {
		log.Fatal(err)
	}

}
