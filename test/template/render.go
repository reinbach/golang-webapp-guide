package template

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/zenazn/goji/web"
)

func GetAbsDir(a ...string) string {
	p, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// this sucks need better way to get abs path to base package
	if p[len(p)-len(PARENT_PACKAGE):] != PARENT_PACKAGE {
		p = path.Dir(p)
	}
	for _, v := range a {
		p = path.Join(p, v)
	}
	return p
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	static_file := r.URL.Path[len(STATIC_URL):]
	static_dir := GetAbsDir("template", STATIC_ROOT)
	if len(static_file) != 0 {
		f, err := http.Dir(static_dir).Open(static_file)
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, r, static_file, time.Now(), content)
			return
		}
	}
	http.NotFound(w, r)
}

func UpdateTemplateList(tmpls []string) []string {
	d := GetAbsDir("template", TEMPLATE_DIR)
	for i, v := range tmpls {
		tmpls[i] = filepath.Join(d, v)
	}
	return tmpls
}

func Render(c web.C, w http.ResponseWriter, r *http.Request, tmpls []string, ctx *Context) {
	ctx.Add("Static", STATIC_URL)

	tmpl_list := UpdateTemplateList(tmpls)
	t, err := template.ParseFiles(tmpl_list...)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, ctx.Values)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
