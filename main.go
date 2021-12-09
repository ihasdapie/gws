package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"reflect"
	"strings"
)

type Page struct {
	Path    string
	Content []byte
}

type Config struct {
	BaseUrl         string   `yaml:"BaseUrl"`
	Serve_Directory string   `yaml:"Serve_Directory"`
	Allowed_Paths   []string `yaml:"Allowed_Paths"`
    Error_404       string   `yaml:"Error_404"`
    Error_403       string   `yaml:"Error_403"`
    Port            string   `yaml:"Port"`
}

func (c Config) printConfig() {
	// fmt.Println(fmt.Sprintf("%+v", c))  // <-- Alternatively just use that -_-
	log.Println("---------- gws config ----------")
	val := reflect.ValueOf(c)
	for i := 0; i < val.NumField(); i++ {
		log.Println(val.Type().Field(i).Name, " : ", val.Field(i))
	}
	log.Println("---------- gws config ----------")

}

func (c *Config) fillConfig(cpath string) *Config {
	yaml_file, err := ioutil.ReadFile(cpath)
	if err != nil {
		log.Printf("Error in reading configuration #%v", err)
	}
	err = yaml.Unmarshal(yaml_file, c)
	if err != nil {
		log.Printf("Error in unmarshalling yaml: %v", err)
	}
	return c
}

func loadPage(pagePath string) (*Page, error) {
	p := path.Clean(pagePath)
	content, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	return &Page{Path: p, Content: content}, nil
}

func contains(s []string, str string) bool {
	for _, i := range s {
		if i == str {
			return true
		}
	}
	return false
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	page_path := GWS_CONFIG.Serve_Directory
	if status == http.StatusNotFound {
		page_path = page_path + GWS_CONFIG.Error_404
	} else if status == http.StatusForbidden {
		page_path = page_path + GWS_CONFIG.Error_403
	}
	pg, _ := loadPage(page_path)
	fmt.Fprintf(w, "%s", pg.Content)
	http.Error(w, http.StatusText(status), status)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("%s Request to %s", r.Method, r.URL.Path))
	path_split := strings.SplitAfter(r.URL.Path, "/")
	page_path := GWS_CONFIG.Serve_Directory + r.URL.Path + "index.html"
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if (path_split[1] != "/") && !contains(GWS_CONFIG.Allowed_Paths, "/"+path_split[1]) {
		errorHandler(w, r, http.StatusForbidden)
		return
	}
	if _, err := os.Stat(page_path); os.IsNotExist(err) {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	pg, _ := loadPage(page_path)
	fmt.Fprintf(w, "%s", pg.Content)
}

var GWS_CONFIG Config

func main() {
	GWS_CONFIG.fillConfig("CONFIG.yml")
	GWS_CONFIG.printConfig()
	http.HandleFunc("/", baseHandler)
	log.Println(fmt.Sprintf("Starting server at %s, serving %s",
		GWS_CONFIG.BaseUrl, path.Base(GWS_CONFIG.Serve_Directory)))
	log.Fatal(http.ListenAndServe(GWS_CONFIG.Port, nil))
}
