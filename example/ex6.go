// package main

// import (
// 	"html/template"
// 	"log"
// 	"os"
// )

// func main() {
// 	const tmpl = "{{.Greeting}} {{.Name}}"

// 	data := struct {
// 		Greeting string
// 		Name     string
// 	}{"Hello", "Joe"}

// 	t := template.Must(template.New("").Parse(tmpl))

// 	err := t.Execute(os.Stdout, data)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// package main

// import (
// 	"html/template"
// 	"log"
// 	"os"
// )

// type Todo struct {
// 	Title string
// 	Done  bool
// }

// type TodoPageData struct {
// 	PageTitle string
// 	Todos     []Todo
// }

// func main() {
// 	tmpl, err := template.ParseFiles("template.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	data := TodoPageData{
// 		PageTitle: "Список дел",
// 		Todos: []Todo{
// 			{Title: "Task 1", Done: false},
// 			{Title: "Task 2", Done: true},
// 			{Title: "Task 3", Done: true},
// 		},
// 	}
// 	tmpl.Execute(os.Stdout, data)
// }

// package main

// import (
// 	"net/http"
// )

// func main() {
// 	fs := http.FileServer(http.Dir("static"))
// 	http.Handle("/", fs)
// 	http.ListenAndServe(":80", nil)
// }

package main

import (
	"io"
	"log"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	log.Println(req.FormValue("param1"))
	io.WriteString(res,
		`<doctype html>
<html>
	<head>
    	<title>Hello world!</title>
	</head>
	<body>
    	My lovely sisi!!
	</body>
</html>`)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":80", nil)
}
