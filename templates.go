package main

import (
	"html/template"
	"log"
	"net/http"
)

var indexTmpl = template.Must(template.New("index.html").Parse(`<html>
<head>
<style>
html, body {
    height: 100%;
}
body {
    margin: 0;
}
.flex-container {
    height: 100%;
    padding: 0;
    margin: 0;
    display: flex;
    align-items: center;
    justify-content: center;
}
.row {
    width: auto;
}
.flex-item {
    padding: 5px;
    margin: 10px;
    color: #757575;
    font-weight: bold;
    font-size: 2em;
    text-align: center;
}
input {
height: 50;
width: 250;
cursor: pointer;
font-size: 20;
color: #444
}
</style>
</head>
<body>
<div class="flex-container">
    <div class="row"> 
        <div class="flex-item"> Token Retrieval </div>
		<form action="/login" method="post" class="flex-item">
			<input type="submit" value="Login">
		</form>
    </div>
</div>
</body>
</html>`))


func renderIndex(w http.ResponseWriter) {
	renderTemplate(w, indexTmpl, nil)
}

type tokenTmplData struct {
	IDToken      string
	AccessToken  string
	RefreshToken string
	RedirectURL  string
	Claims       string
}

var tokenTmpl = template.Must(template.New("token.html").Parse(`<html>
  <head>
    <style>
/* make pre wrap */
pre {
 white-space: pre-wrap;       /* css-3 */
 white-space: -moz-pre-wrap;  /* Mozilla, since 1999 */
 white-space: -pre-wrap;      /* Opera 4-6 */
 white-space: -o-pre-wrap;    /* Opera 7 */
 word-wrap: break-word;       /* Internet Explorer 5.5+ */
}
    </style>
  </head>
  <body>
    <p> Access Token: <pre><code>{{ .AccessToken }}</code></pre></p>
  </body>
</html>
`))

func renderToken(w http.ResponseWriter, accessToken string) {
	renderTemplate(w, tokenTmpl, tokenTmplData{
		AccessToken:  accessToken,
	})
}

func renderTemplate(w http.ResponseWriter, tmpl *template.Template, data interface{}) {
	err := tmpl.Execute(w, data)
	if err == nil {
		return
	}

	switch err := err.(type) {
	case *template.Error:
		// An ExecError guarantees that Execute has not written to the underlying reader.
		log.Printf("Error rendering template %s: %s", tmpl.Name(), err)

		// TODO(ericchiang): replace with better internal server error.
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	default:
		// An error with the underlying write, such as the connection being
		// dropped. Ignore for now.
	}
}
