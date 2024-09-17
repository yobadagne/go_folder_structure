package main

import (
	// "fmt"
	// "go_folder_structure/constant"
	createfiles "go_folder_structure/create_files"
	createfolders "go_folder_structure/create_folders"
	"log"
	// "os"
)
func main()  {
	// create folder
	err := createfolders.CreateFolders()
	if err!= nil{
		log.Fatal(err)
	}
	
	// create files
	err =  createfiles.CreateFiles()
	if err!= nil{
		log.Fatal(err)
	}
}