package main


import(
	"net/http"
	"log"
)

func main () {

	const port = "8080"
	const filepathRoot = "."

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))

	srv := &http.Server{
			Addr: ":" + port,
			Handler: mux,
		}
	log.Println("Serving files from ",filepathRoot," on port: ", port)
	log.Fatal(srv.ListenAndServe())
}
