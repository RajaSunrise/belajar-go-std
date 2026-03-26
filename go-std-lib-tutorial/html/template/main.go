package main

import (
	"html/template"
	"log"
	"os"
)

type PageData struct {
	Title   string
	Heading string
	Items   []string
}

func main() {
	// Template HTML dengan sintaks aksi {{ }}
	tmplString := `
<!DOCTYPE html>
<html>
<head>
	<title>{{.Title}}</title>
</head>
<body>
	<h1>{{.Heading}}</h1>
	<ul>
	{{range .Items}}
		<li>{{.}}</li>
	{{else}}
		<li>Tidak ada item.</li>
	{{end}}
	</ul>
</body>
</html>
`

	// Parsing template
	tmpl, err := template.New("webpage").Parse(tmplString)
	if err != nil {
		log.Fatal(err)
	}

	// Data yang akan dimasukkan ke dalam template
	data := PageData{
		Title:   "Belajar HTML Template di Go",
		Heading: "Daftar Bahasa Pemrograman",
		Items:   []string{"Go", "Python", "JavaScript", "Rust"},
	}

	// Mengeksekusi template dan menulis hasilnya ke stdout (bisa juga ke http.ResponseWriter)
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
