package main

import (
	"fmt"
	"log"

	"github.com/stanlee321/assets_service/pkg"

)

func main(){
	fileName := "file.pdf"

	URL := "http://201.222.81.2/AnalisisNormativo/normasPDF/9495.pdf"

	err := pkg.DownloadFile(fileName, URL)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File %s downlaod in current working directory", fileName)
}