package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type user struct {
	// mapping with api response keys
	Name  string `json:"name"`
	Email string `json:"email"`
	Dob   string `json:"dob"`
	Image string `json:"image,omitempty"`
}

func main() {
	fmt.Println("Welcome to MOD tut")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/users", serveUsers).Methods("GET")

	// log if error happens
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Listening at : http://localost:4000/")

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(webPage))
}

func serveUsers(w http.ResponseWriter, r *http.Request) {
	myUsers := []user{
		{"Hitesh Choudhary", "hitesh@lco.dev", "1975-01-01", "https://avatars.githubusercontent.com/u/11613311?v=4"},
		{"Bryan Bolt", "bolt121@gmail.com", "1855-01-01", "https://wl-brightside.cf.tsp.li/resize/728x/jpg/978/58b/cecf2d560ca3833915b1fad736.jpg"},
		{"LCO", "lco123@gmail.com", "2015-12-12", "https://lco.dev/images/logo-white.png"},
	}
	// json formatted data
	// jsonData, _ := json.Marshal(myUsers)

	json.NewEncoder(w).Encode(myUsers)
}

/*
go mode build
go mod run // server will start at port 4000


// will removes "//indirect" from mod file those packages were used
go mod tidy

// verifying the SHA/SUM and security of imported third party modules
go mod verify


// on what our application is dependent on (third party) only imported
go list

// DEPENDENT on these modules
go list -m all

// Check all available third paty module
go list -m -versions github.com/gorilla/mux


// why i'm dependent on a mentioned module
go mod why github.com/gorilla/mux

// Which kg are dependent on which
go mod graph

// Edit mod file info with cmd
// here go version will be changed to 17.16
go mod edit -go 17.16


// Bring all required modules like (node_modules folder)
go mod vendor
*/

const webPage = `
<!DOCTYPE html>
<html>
    <head>
        <title>HOME | Kishan's GO API</title>
        <link rel="preconnect" href="https://fonts.googleapis.com" />
        <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
        <link href="https://fonts.googleapis.com/css2?family=Oleo+Script&display=swap" rel="stylesheet" />

        <style>
            .header {
                font-family: "Oleo Script", cursive;
                padding: 10px;
                color: white;
                background-color: #596275;
                text-align: center;
                font-size: 50px;
                border-radius: 5px;
                text-shadow: 2px 2px 4px #000000;
                box-shadow: 1px 2px 10px #000;
            }
            .button {
                padding: 15px;
                background-color: #242b2e;
                text-align: center;
                font-size: 20px;
                font-weight: bold;
                box-shadow: 2px 2px 10px #000;
                border-radius: 5px;
            }
            a {
                text-decoration: none;
                width: 160px;
                color: white;display:block;
            }
            .button:hover {
                background-color: purple;
                transition: 0.5s;
                transform: scale(1.05);
                background-color: #5a20cb;
            }
			.avatar {
  				vertical-align: middle;
  				width: 200px;
  				height: 200px;
  				border-radius: 50%;
                box-shadow: 2px 2px 10px #000;
			}
			img:hover{
				transition: 0.4s;
                transform: scale(1.05);
                box-shadow: 2px 2px 10px #596275;
			}
			img{
				margin-top: 60px;
			}
        </style>
    </head>
    <body style="background-color: #303952; color: white;">
        <div class="header">Kishan Kumar Sharma</div>
        <center>
			<img src="https://avatars.githubusercontent.com/u/36340195?v=4" alt="Kishan Image" class="avatar" />
                
            <h1 style="align: center;text-shadow: 2px 2px 10px #000;">Welcome to Kishan API with GO.</h1>
			
            <h2>GET | POST | PUT | DELETE</h2>
            <h3>Get your api key at skishan781[at]gmail.com</h3>
            <a href="https://github.com/iamkishansharma" target="_blank"><div class="button">GitHub</div></a>
        </center>
    </body>
</html>

`
