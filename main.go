package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
)

//localstruct
type localstruct struct {
	FILES []string
}

var hostname, userName, password, enter string

func main() {

	//These are login variables for the REMOTE/target computer.
	fmt.Printf("Enter a hostname(IP):  ")
	fmt.Scanln(&hostname)
	fmt.Printf("Enter a username:  ")
	fmt.Scanln(&userName)
	fmt.Printf("Enter a password:  ")
	fmt.Scanln(&password)
	enter = userName + "@" + hostname

	http.HandleFunc("/", index)
	http.HandleFunc("/remotefiles.html", remotefiles)
	http.HandleFunc("/localfiles.html", localfiles)
	http.HandleFunc("/formsubmit", formsubmit)
	fmt.Println(" Open browser to localhost:7003")
	http.ListenAndServe(":7003", nil)
}

//index runs the index page
func index(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/index.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}

//remotefiles displays remote computer files to html page.
func remotefiles(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/remotefiles.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	remote, err := exec.Command("ssh", enter, "ls", "/home/"+userName+"/servercatchbox", ">", "file1", ";", "tail", "file1").Output()
	g := localstruct{FILES: make([]string, 1)}
	length := 0
	if err != nil {
		fmt.Println(err)
	}
	for l := 0; l < len(remote); l = l + 1 {
		if remote[l] != 10 {
			g.FILES[length] = g.FILES[length] + string(remote[l])
		} else {
			g.FILES = append(g.FILES, "\n")
			length = length + 1
		}
	}
	temp.Execute(response, g)
}

//localfiles displays host computer files to html page.
func localfiles(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/localfiles.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	//fmt.Println("Local TEST PASS")
	remote, err := exec.Command("ls", "/home/garner/servercatchbox", ">", "file1", ";", "tail", "file1").Output()
	g := localstruct{FILES: make([]string, 1)}
	length := 0

	if err != nil {
		fmt.Println(err)
	}

	for l := 0; l < len(remote); l = l + 1 {
		if remote[l] != 10 {
			g.FILES[length] = g.FILES[length] + string(remote[l])
		} else {
			g.FILES = append(g.FILES, "\n")
			length = length + 1
		}
	}
	temp.Execute(response, g)
}

//prototype
func formsubmit(response http.ResponseWriter, request *http.Request) {
	fmt.Println("WIP", request.FormValue("transfer1"))
}
