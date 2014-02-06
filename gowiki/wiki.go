// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//To DO
//Store templates in tmpl/ and page data in data/.
//Add a handler to make the web root redirect to /view/FrontPage.
//Spruce up the page templates by making them valid HTML and adding some CSS rules.
//Implement inter-page linking by converting instances of [PageName] to
//<a href="/view/PageName">PageName</a>. (hint: you could use regexp.ReplaceAllFunc to do this)

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Page struct {
	Title string
	Body  []byte
}

//The type []byte means "a byte slice". (See Slices: usage and internals for
// more on slices.)
//The Body element is a []byte rather than string because that is the type
// expected by the io libraries we will use, as you'll see below.

func (p *Page) save(newTitle string) error {
	filename := p.Title + ".txt"
	if p.Title != newTitle {
		os.Remove(filename)
		filename = newTitle + ".txt"
	}
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//Functions can return multiple values. The standard library function io.ReadFile returns []byte and error.
//In loadPage, error isn't being handled yet; the "blank identifier" represented by the underscore (_) symbol is used to throw away the error return value (in essence, assigning the value to nothing).

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

//An http.Request is a data structure that represents the client HTTP request.
// r.URL.Path is the path component of the request URL.
//The trailing [1:] means "create a sub-slice of Path from the
//1st character to the end." This drops the leading "/" from the path name.
//First, this function extracts the page title from r.URL.Path,
// the path component of the request URL.
// The Path is re-sliced with [len("/view/"):] to drop the leading "/view/"
//component of the request path.
// This is because the path will invariably begin with "/view/", which is
// not part of the page's title.
//The http.Redirect function adds an HTTP status code of http.StatusFound (302)
// and a Location header to the HTTP response.
func listHandler(w http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadDir(".")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//files := make([]string, 0)
	files := []string{}
	for _, f := range d {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".txt") {
			files = append(files, strings.TrimSuffix(f.Name(), ".txt"))
		}
	}
	err = templates.ExecuteTemplate(w, "list.html", files)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	newTitle := r.FormValue("newTitle")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save(newTitle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+newTitle, http.StatusFound)
}

//The function template.Must is a convenience wrapper that panics when passed a
//non-nil error value, and
//otherwise returns the *Template unaltered. A panic is appropriate here;
//if the templates can't be loaded the only sensible
//thing to do is exit the program.

//The ParseFiles function takes any number of string arguments that identify
//our template files, and parses those files into
//templates that are named after the base file name.
//If we were to add more templates to our program,
//we would add their names to the ParseFiles call's arguments.

var templates = template.Must(template.ParseFiles("edit.html", "view.html", "list.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/?([a-zA-Z0-9]*)$")

//The returned function is called a closure because it encloses values defined outside of it.
//In this case, the variable fn (the single argument to makeHandler) is enclosed by the closure.
//The variable fn will be one of our save, edit, or view handlers.

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			fmt.Printf("Can't find %v\n", r.URL.Path)
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		listHandler(w, r)
	})
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.ListenAndServe(":8080", nil)
}
