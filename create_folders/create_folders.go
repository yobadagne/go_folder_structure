package createfolders

import (
	"errors"
	"go_folder_structure/constant"
	"os"
)

func CreateFolders() error {

	if len(os.Args) <= 1 {
		return errors.New("please provide a root folder name")
	}
	folderName := os.Args[1]
	if err := os.Mkdir(folderName, 0755); err != nil {
		return errors.New(err.Error())
		
	}
	//fmt.Println("folder created")
	err := os.Chdir(folderName)
	if err != nil {
		return errors.New(err.Error())
		
	}
	for _, v := range constant.Root_folders {
		if err := os.Mkdir(v, 0755); err != nil {
			return errors.New(err.Error())
		
		}
	}
	// create folder in internal
	err = os.Chdir("internal")
	if err != nil {
		return errors.New(err.Error())
		
	}
	for _, v := range constant.Internal_folders {
		if err := os.Mkdir(v, 0755); err != nil {
			return errors.New(err.Error())
		
		}
	}
	// create folders in constant
	err = os.Chdir("constant")
	if err != nil {
		return errors.New(err.Error())
		
	}
	for _, v := range constant.Constant_folders {
		if err := os.Mkdir(v, 0755); err != nil {
			return errors.New(err.Error())
		
		}
	}
	// create folders in glue
	err = os.Chdir("../glue")
	if err != nil {
		return errors.New(err.Error())
		
	}
	for _, v := range constant.Glue_folders {
		if err := os.Mkdir(v, 0755); err != nil {
			return errors.New(err.Error())
		
		}
	}
	// create folders in handler
	
	err = os.Chdir("../handler")
	if err != nil {
		return errors.New(err.Error())
		
	}
	for _, v := range constant.Handler_folders{
		if err := os.Mkdir(v, 0755); err != nil {
			return errors.New(err.Error())
		
		}
	}
	// create folders in query
	
	err = os.Chdir("../constant/query")
	if err != nil {
		return errors.New(err.Error())
		
	}
	for _, v := range constant.Query_folders{
		if err := os.Mkdir(v, 0755); err != nil {
			return errors.New(err.Error())
		
		}
	}

	// create folders in model
	
	err = os.Chdir("../model")
	if err != nil {
		return errors.New(err.Error())
		
	}
	for _, v := range constant.Model_folders{
		if err := os.Mkdir(v, 0755); err != nil {
			return errors.New(err.Error())
		
		}
	}
	// create folders in test
	
	err = os.Chdir("../../../test")
	if err != nil {
		return errors.New(err.Error())
		
	}
	for _, v := range constant.Test_folders{
		if err := os.Mkdir(v, 0755); err != nil {
			return errors.New(err.Error())
		
		}
	}
	// go back to root
	err = os.Chdir("../")
	if err != nil {
		return errors.New(err.Error())
		
	}
	return nil
}
