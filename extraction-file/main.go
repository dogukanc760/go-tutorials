package main

import (
	"archive/tar"
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var fileFolderPath = "./files/"
var files = []string{fileFolderPath + "demo.go", fileFolderPath + "note1.txt"}

// TAR İŞLEMLERİ BAŞLANGIÇ
func addFile(fileName string, tw *tar.Writer) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Dosya açılırken bir hata meydana geldi %s: %s", fileName, err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Bilgileri alınırken bir hata meydana geldi %s: %s", fileName, err)
	}

	hdr := &tar.Header{
		ModTime: stat.ModTime(),
		Name:    fileName,
		Size:    stat.Size(),
		Mode:    int64(stat.Mode().Perm()),
	}

	err2 := tw.WriteHeader(hdr)
	if err2 != nil {
		msg := "Tar header yazılırken bir hata meydana geldi %s: %s"
		return fmt.Errorf(msg, fileName, err2)
	}

	copied, err3 := io.Copy(tw, file)
	if err3 != nil {
		msg := "Tar kopyalanırken bir hata meydana geldi %s: %s"
		return fmt.Errorf(msg, fileName, err2)
	}
	if copied < stat.Size() {
		msg := "%s dosyasına %d kadar yazıldı ancak beklenen veri %d kadardı."
		return fmt.Errorf(msg)
	}

	return nil

}

func createArchiveTARFile(archiveFileName string) int {
	if len(archiveFileName) == 0 {
		return -1
	}

	// Dosya üzerinde ki izinler, ne yapabileceğini söylüyoruz.
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(archiveFileName+".tar", flags, 0644)
	if err != nil {
		return -1
	}
	defer file.Close()

	tw := tar.NewWriter(file)
	defer tw.Close()

	for _, fileName := range files {
		if err := addFile(fileName, tw); err != nil {
			log.Fatalf("Hata Oluştu=>%s: %d", fileName, err)
		}
	}
	return 1
}

// TAR İŞLEMLERİ BİTİŞ

// ZIP işlemleri başlangıç
func addFileZIP(fileName string, zw *zip.Writer) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Hata Meydana Geldi=>%s", err)
	}
	defer file.Close()

	wr, err := zw.Create(fileName)
	if err != nil {
		return fmt.Errorf("Hata Meydana Geldi=>%s", err)
	}

	if _, err := io.Copy(wr, file); err != nil {
		return fmt.Errorf("Hata Meydana Geldi=>%s", err)
	}
	return nil
}

func createArchiveZIPFile(archiveFileName string) int {
	if len(archiveFileName) < 1 {
		log.Fatal("Dosya adı 1 haneden küçük olamaz!")
		return -1
	}

	flags := os.O_RDONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(archiveFileName+".zip", flags, 0644)
	if err != nil {
		log.Fatal("Bir hata meydana geldi=>%s", err)
		return -1
	}
	defer file.Close()
	zw := zip.NewWriter(file)
	defer zw.Close()
	for _, fileName := range files {
		if err := addFileZIP(fileName, zw); err != nil {
			log.Fatal("Hata meydana geldi=>%s", err)
			return -1
		}
	}
	return 1

}

func CreateDirIfNotExists(dirName string) {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err = os.MkdirAll(dirName, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func extractionZIPFile() {
	zr, err := zip.OpenReader("./files/virusdosyası.zip")
	if err != nil {
		log.Fatalf("Bir hata meydana geldi=> %s", err)
	}
	defer zr.Close()

	for _, file := range zr.Reader.File {
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatalf("Bir hata meydana geldi=> %s", err)
		}
		defer zippedFile.Close()

		targetDir := "./"
		extractedFilePath := filepath.Join(targetDir, file.Name)

		dirName := strings.Split(file.Name, "\\")
		CreateDirIfNotExists(dirName[0])

		if file.FileInfo().IsDir() {
			log.Println("Klasör Oluşturuluyor: ", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			log.Println("Dosya çıkartılıyor.....", file.Name)
			outFile, err := os.OpenFile(
				extractedFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode(),
			)
			if err != nil {
				log.Fatalf("Bir hata meydana geldi=> %s", err)
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, zippedFile)
			if err != nil {
				log.Fatalf("Bir hata meydana geldi=> %s", err)
			}

		}
	}
}

func main() {
	// Tar Dosyaları Üretmek
	result := createArchiveTARFile("ersankuneri")
	if result > 0 {
		fmt.Println("İşlem Başarılı: ", result)
	} else {
		log.Fatal("İşlem başarısız")
	}

	// ZIP dosyaları üretmek
	resultZip := createArchiveZIPFile("virusdosyası")
	if resultZip > 0 {
		fmt.Println("İşlem Başarılı: ", resultZip)
	} else {
		log.Fatalf("Bir hata meydana geldi=> %d", resultZip)
	}

	// ZİP dosyası çıkartmak
	// Fonksiyona giderek görebilirsiniz.
	extractionZIPFile()

}
