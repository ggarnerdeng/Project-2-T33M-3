package main

import (
	"fmt"
	"net/http"

	//"text/html"

	"github.com/sfreiberg/simplessh"
)

//scp USER@HOSTNAME:/home/USER/test13.txt /home/USER2/Documents/
//^^ Code to copy one file from remote computer to local computer.

//How to list all files in the catchbox
//ssh USERNAME@HOSTNAME ls /home/USER/servercatchbox

func main() {
	var hostname, userName, password string
	fmt.Printf("Enter a hostname(IP):  ")
	fmt.Scanln(&hostname)
	fmt.Printf("Enter a username:  ")
	fmt.Scanln(&userName)
	fmt.Printf("Enter a password:  ")
	fmt.Scanln(&password)

	//location := userName + "@" + hostname + ":/home"

	//TREE: view all files as a tree
	//a = response
	//b = request

	//AUTHLOG
	http.HandleFunc("/authlog", func(a http.ResponseWriter, b *http.Request) {
		var authlog = terminalCommand(hostname, userName, password, "ls /home/"+userName+"/servercatchbox")
		fmt.Fprintf(a, "%s\n", "Viewing servercatchbox")
		fmt.Fprintln(a, " ")
		fmt.Fprintln(a, " ")
		fmt.Fprintf(a, "%s", authlog)
	})

	http.HandleFunc("/", func(a http.ResponseWriter, b *http.Request) {
		var output = terminalCommand(hostname, userName, password, "ps")
		var html = `<html>
		<head>
		
		<style>
		ul {
		list-style-type: none;
		margin: 0;
		padding: 0;
		overflow: hidden;
		background-color: #333333;
	  }
	  
	  li {
		float: center;
	  }
	  
	  li a {
		display: block;
		color: red;
		text-align: left;
		padding: 16px;
		text-decoration: none;
	  }
	  
	  li a:hover {
		background-color: #111111;
	  }
		* {
		 box-sizing: border-box; 
		}
		
		body {
		  margin: 0;
		}
		#main {
		  display: flex;
		  min-height: calc(100vh - 40vh);
		}
		#main > article {
		  flex: 1;
		}
		
		#main > nav, 
		#main > aside {
		  flex: 0 0 20vw;
		  background: beige;
		}
		#main > nav {
		  order: -1;
		}
		header, footer, article, nav, aside {
		  padding: 1em;
		}
		header, footer {
		  background: yellowgreen;
		  height: 20vh;
		}
	  </style>
			</head>
	  <body>
		<header>Logged in as garner@192.168.1.33</header>
		<div id="main">
		<title> catchbox</title>
		  <article>Command Line Options
		  <ol>
			  <li><a href="/authlog">View files in servercatchbox</a></li>
			  <li><a href="/authlog">CLONE OF 1Upload files to servercatchbox</a></li>
			  <li><a href="/authlog">CLONE OF 1Download files from servercatchbox</a></li>
		  </ol></article>
		  <nav></nav>
		  <aside></aside>
		</div>
		<footer></footer>
	  </body>
					</html>
			`
		fmt.Fprintf(a, html, output)
	})
	fmt.Println()
	fmt.Println("Successfully connected;")
	fmt.Println("Open Localhost:12345")
	http.ListenAndServe(":12345", nil)
}

func terminalCommand(hostname string, userName string, password string, command string) []byte {
	var client *simplessh.Client
	var err error
	if client, err = simplessh.ConnectWithPassword(hostname, userName, password); err != nil {
		fmt.Print(err)
	}
	if err != nil {
		panic(err)
	}
	defer client.Close()
	output, err := client.Exec(command)
	if err != nil {
		panic(err)
	}
	return output
}
