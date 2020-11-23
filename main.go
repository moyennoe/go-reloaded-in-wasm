package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	piscine "./exercies"
)

// Load the index.html template.
var tmpl = template.Must(template.New("tmpl").ParseFiles("index.html"))

var templatesDir = os.Getenv("TEMPLATES_DIR")

func main() {

	fmt.Println("\n\033[2;31m _______  __   __    __   __  __    _  _______  _______ ")
	fmt.Println("|  _    ||  | |  |  |  |_|  ||  |  | ||       ||       |")
	fmt.Println("| |_|   ||  |_|  |  |       ||   |_| ||   _   ||    ___|")
	fmt.Println("|       ||       |  |       ||       ||  | |  ||   |___ ")
	fmt.Println("|  _   | |_     _|  |       ||  _    ||  |_|  ||    ___|")
	fmt.Println("| |_|   |  |   |    | ||_|| || | |   ||       ||   |___ ")
	fmt.Println("|_______|  |___|    |_|   |_||_|  |__||_______||_______|\033[0m")
	fmt.Println("==>  * localhost:9000 *  <==")

	// Serve / with the index.html file.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	//func Atoi
	http.HandleFunc("/Atoi", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		atoi := r.Form["machaine"][0]
		fmt.Fprintln(w, piscine.Atoi(atoi))
	})

	//func RecursivePower
	http.HandleFunc("/RecursivePower", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		nb, _ := strconv.Atoi(r.Form["machaine1"][0])
		power, _ := strconv.Atoi(r.Form["machaine2"][0])
		fmt.Fprintln(w, nb, "^", power, "=", piscine.RecursivePower(nb, power))
	})

	//func PrintCombN
	http.HandleFunc("/PrintCombN", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		printcombn, _ := strconv.Atoi(r.Form["machaine"][0])
		fmt.Fprint(w, piscine.PrintCombN(printcombn))
	})

	//func ActiveBits
	http.HandleFunc("/ActiveBits", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		activebits, _ := strconv.Atoi(r.Form["machaine"][0])
		fmt.Fprintln(w, "La représentation binaire du nombre entier", activebits, "est:", piscine.ActiveBits(activebits), "bits")
	})

	//func Doop
	http.HandleFunc("/Doop", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		arg0, _ := strconv.Atoi(r.Form["machaine1"][0])
		arg1 := r.Form["machaine2"][0]
		arg2, _ := strconv.Atoi(r.Form["machaine3"][0])

		fmt.Fprintln(w, piscine.Doop(arg0, arg1, arg2))
	})

	//func PrintNbrBase
	http.HandleFunc("/PrintNbrBase", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		base, _ := strconv.Atoi(r.Form["machaine1"][0])
		nbr := r.Form["machaine2"][0]

		fmt.Fprintln(w, piscine.PrintNbrBase(base, nbr))
	})

	//func AtoiBase
	http.HandleFunc("/AtoiBase", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		base := r.Form["machaine1"][0]
		nbr := r.Form["machaine2"][0]

		fmt.Fprintln(w, piscine.AtoiBase(base, nbr))
	})
	//func SplitWhiteSpaces
	http.HandleFunc("/SplitWhiteSpaces", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		split := r.Form["machaine"][0]
		fmt.Fprintln(w, piscine.SplitWhiteSpaces(split))
	})
	//func ConvertBase
	http.HandleFunc("/ConvertBase", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		str1 := r.Form["machaine1"][0]
		str2 := r.Form["machaine2"][0]
		str3 := r.Form["machaine3"][0]
		fmt.Fprintln(w, piscine.ConvertBase(str1, str2, str3))
	})

	//func Split
	http.HandleFunc("/Split", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		str1 := r.Form["machaine1"][0]
		str2 := r.Form["machaine2"][0]
		fmt.Fprintln(w, piscine.Split(str1, str2))
	})
	// Start the server at http://localhost:9000
	log.Fatal(http.ListenAndServe(":9000", nil))
}
