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
	fmt.Printf("Enter a hostname(IP) of target remote computer:  ")
	fmt.Scanln(&hostname)
	fmt.Printf("Enter a username of this target remote computer:  ")
	fmt.Scanln(&userName)
	fmt.Printf("Enter a password for this target remote computer:  ")
	fmt.Scanln(&password)
	enter = userName + "@" + hostname

	http.HandleFunc("/", index)
	http.HandleFunc("/remotefiles.html", remotefiles)
	http.HandleFunc("/localfiles.html", localfiles)
	http.HandleFunc("/downloader", downloader)
	http.HandleFunc("/uploader", uploader)
	fmt.Println(" Open browser to localhost:7004")
	http.ListenAndServe(":7004", nil)
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
	enter = userName + "@" + hostname
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
		fmt.Print("some error")
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

/*prototype
func formsubmit(response http.ResponseWriter, request *http.Request) {
	fmt.Println("WIP", request.FormValue("transfer1"))
}*/

//prototype UPLOAD
func uploader(response http.ResponseWriter, request *http.Request) {
	var upload1 = request.FormValue("upload1")
	fmt.Println("The file to upload is: " + upload1)
	//enter = userName + "@" + hostname
	temp, _ := template.ParseFiles("html/localfiles.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	//scp /home/garner/servercatchbox/localtest1 garner@192.168.1.33:/home/garner/servercatchbox ;
	// ls /home/garner/servercatchbox > file1 ; tail file1
	//remote, err := exec.Command("scp", "/home/garner/servercatchbox/"+upload, enter+":/home/"+userName+"/servercatchbox", ";", "ls", "/home/garner/servercatchbox", ">", "file1", ";", "tail", "file1").Output()
	remote, err := exec.Command("ls", "/home/garner/servercatchbox", ">", "file1", ";", "tail", "file1").Output()
	remote2 := exec.Command("scp", "/home/garner/servercatchbox/"+upload1, "garner@192.168.1.33:/home/garner/servercatchbox")

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
	remote2.Run()
	temp.Execute(response, g)
}

//downloader
func downloader(response http.ResponseWriter, request *http.Request) {
	var download1 = request.FormValue("download1")
	fmt.Println("The file to download is: " + download1)
	//enter = userName + "@" + hostname
	temp, _ := template.ParseFiles("html/remotefiles.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	//scp /home/garner/servercatchbox/localtest1 garner@192.168.1.33:/home/garner/servercatchbox ;
	// ls /home/garner/servercatchbox > file1 ; tail file1
	//remote, err := exec.Command("scp", "/home/garner/servercatchbox/"+upload, enter+":/home/"+userName+"/servercatchbox", ";", "ls", "/home/garner/servercatchbox", ">", "file1", ";", "tail", "file1").Output()
	remote, err := exec.Command("ssh", enter, "ls", "/home/"+userName+"/servercatchbox", ">", "file1", ";", "tail", "file1").Output()
	remote2 := exec.Command("scp", "garner@192.168.1.33:/home/garner/servercatchbox/"+download1, "/home/garner/servercatchbox")

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
	remote2.Run()
	temp.Execute(response, g)
}
