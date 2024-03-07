package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

type Page struct {
	Title string
	Body []byte   // Using bytes rather than string for os read/write libraries
}

type PageHTML struct {
	Title string
	Body template.HTML
}

var templates = template.Must(template.ParseFiles("tmpl/edit.html", "tmpl/view.html"))
var validPath = regexp.MustCompile("^/(view|edit|save)/([a-zA-Z0-9_]+)$")
var pageLink = regexp.MustCompile(`\[[a-zA-Z0-9_]+\]`)

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/", rootHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

	// Test page struct
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page.")}
	// p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))
}

// Save a page for persistent storage
func (p *Page) save() error {
	filename := "data/" + p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

// Load a page
func loadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
// 	m := validPath.FindStringSubmatch(r.URL.Path)
	
// 	if m == nil {
// 		http.NotFound(w, r)
// 		return "", errors.New("Invalid Page Title")
// 	}

// 	return m[2], nil
// }

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/FrontPage", http.StatusFound)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)

		if m == nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, m[2])
	}
}

func renderTemplate[PageType Page | PageHTML](w http.ResponseWriter, tmpl string, p *PageType) {
	err := templates.ExecuteTemplate(w, tmpl + ".html", p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	// load the page
	p, err := loadPage(title)

	if err != nil {
		http.Redirect(w, r, "/edit/" + title, http.StatusFound) // status code 302
		// fmt.Fprintf(w, "<h1>Error not found.</h1><div>Cannot find the request page %s.</div>", title)
		return
	}

	p.Body = pageLink.ReplaceAllFunc(p.Body, addInterpageLinks)
	pHTML := &PageHTML{Title: p.Title, Body: template.HTML(p.Body)}

	renderTemplate(w, "view", pHTML)

	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)

	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit", p)

	// fmt.Fprintf(w, "<h1>Editing %s</h1>"+
    //     "<form action=\"/save/%s\" method=\"POST\">"+
    //     "<textarea name=\"body\">%s</textarea><br>"+
    //     "<input type=\"submit\" value=\"Save\">"+
    //     "</form>",
    //     p.Title, p.Title, p.Body)
}

func addInterpageLinks(s []byte) []byte {
	s = s[1:len(s)-1]
	return []byte(fmt.Sprintf("<a href=\"/view/%s\">%s</a>", s, s))
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.Body = pageLink.ReplaceAllFunc(p.Body, addInterpageLinks)

	err := p.save()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}
