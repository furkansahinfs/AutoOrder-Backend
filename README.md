# AutoOrder
CENG415 - Thesis Project

```
├─ bin                           //The folder where the binary file was created
├─ cmd                           //The code that started it all
├─ config.yml                    //Config showing which information the backend server will run
├─ go.mod                        //3rd party libraries
├─ go.sum                        //Sums and versions of 3rd party libraries
├─ makefile                      //MakeFile for version control and creation of binary file
└─ pkg                           //Server codes 
   ├─ api                        //Api Layer for all aplication
   ├─ errors                     //For generic errors which served to mobile
   ├─ images                     //For saving images to this folder
   ├─ model                      //Models for every type of object
   ├─ repository                 //DB Layer
   │  ├─ configuration  
   │  ├─ image
   │  ├─ user
   │  └─ user_information
   ├─ server                     //Server Layer for all aplication. Main part of http server
   ├─ service                    //Service Layer
   │  ├─ configuration
   │  ├─ image
   │  ├─ user
   │  └─ user_information
   ├─ static                     //Auto generated static files
   └─ version                    //Version control&save for git