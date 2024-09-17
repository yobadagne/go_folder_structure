package createfiles

import ( 
	// "errors"
 	"fmt"
 	"go_folder_structure/constant"
 	"os"
)

func CreateFiles() error {
	// create root files
	for _, v := range constant.Root_folder_files{
		_, err := os.Create(v)
		if err != nil{
			return fmt.Errorf("could not create file %s", v)
		}
		
	} 
	for _, v := range constant.Root_folders{
		switch  v {
			case "cmd":
				os.Chdir(v)
				_, err := os.Create("main.go")
					if err != nil{
						return fmt.Errorf("could not create file main.go", )
					}
					os.Chdir("../")
			case "config":
				os.Chdir(v)
				for _, v := range constant.Config_folder_files{
					_, err := os.Create(v)
					if err != nil{
						return fmt.Errorf("could not create file %s", v)
					}
					
				}
				os.Chdir("../") 
			case "initiator":
				os.Chdir(v)
				for _, v := range constant.Initiator_folder_files{
					_, err := os.Create(v)
					if err != nil{
						return fmt.Errorf("could not create file %s", v)
					}
					
				} 
				os.Chdir("../") 
			case "internal":
				os.Chdir("internal/constant")
				for _, v := range constant.Constant_folder_files{
					_, err := os.Create(v)
					if err != nil{
						return fmt.Errorf("could not create file %s", v)
					}
				} 
				//os.Chdir("../")
				os.Chdir("errors")
				for _, v := range constant.Constant_errors_folder_files{
					_, err := os.Create(v)
					if err != nil{
						return fmt.Errorf("could not create file %s", v)
					}
				}
				os.Chdir("../../") // back to internal

				os.Chdir("glue/routing")
				for _, v := range constant.Glue_folder_files{
					_, err := os.Create(v)
					if err != nil{
						return fmt.Errorf("could not create file %s", v)
					}
					
				} 
				os.Chdir("../../") // back to internal
				
				os.Chdir("handler/middleware")
				for _, v := range constant.Handler_middleware_folder_files{
					_, err := os.Create(v)
					if err != nil{
						return fmt.Errorf("could not create file %s", v)
					}
					
				} 
				os.Chdir("../../") // back to internal
			
				os.Chdir("handler/rest")
				for _, v := range constant.Handler_rest_folder_files{
					_, err := os.Create(v)
					if err != nil{
						return fmt.Errorf("could not create file %s", v)
					}
					
				} 
				os.Chdir("../../") // back to internal
				
				os.Chdir("module")
				for _, v := range constant.Module_folder_files{
					_, err := os.Create(v)
					if err != nil{
						return fmt.Errorf("could not create file %s", v)
					}
					
				}
				os.Chdir("../") // back to internal
				os.Chdir("persistence")
				for _, v := range constant.Persistence_folder_files{
					_, err := os.Create(v)
					if err != nil{
						return fmt.Errorf("could not create file %s", v)
					}
					
				}
			default:
		}
	}
	return nil
}