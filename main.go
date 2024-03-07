<<<<<<< HEAD
package main

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/http"
	"math/rand"
	"strconv"
)

type Movie struct {
	ID string 'json:"id"'
	Isbn string 'json:"isbn"'
	Title string 'json:"title"'
	Director *Director 'json:"director"'
}

type Director struct {
	Firstname string 'json:"firstname"'
	Lastname string 'json:"lastname"'
}

var movies[] Movie

func main() {
	
	r :=mux.NewRouter()

	r.HandleFunc()
=======
// package main

// import "fmt"

// func main() {
//     fmt.Println("Hello, World!")
// }

package main

import "fmt"

func main() {
	i:= 2
	fmt.Println(i)
>>>>>>> origin/main
}
