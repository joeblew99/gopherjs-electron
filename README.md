# gopherjs-electron
Atom/Electron (https://github.com/atom/electron/) desktop apps with Go. Package use https://github.com/gopherjs/gopherjs.

I am developing on OSX, but testing on Windows and Ubuntu. 

Work in Progress for Windows & Ubuntu still...

# Install

Package require electron-prebuilt npm-module

	npm install -g electron-prebuilt


# Usage

Look in cmd for basic example. 

````
# Kick off backend and front end together.
go run all.go
````


# Roadmap
- DONE: When Electron starts, start / spawn the Golang backend web server
- DONE: Get basic communication between Backend and Frontend working.
- Packaging for all OS's
- More gopherjs modules: Polymer, etc, etc



# Test

For tests used Jasmine (http://jasmine.github.io/) and adapter fo go https://github.com/arvitaly/gopherjs-jasmine

	go get github.com/arvitaly/gopherjs-jasmine
	npm install electron-prebuilt -g
	npm install
	npm test

# Docker

Also you can use docker-image for electron (includes NodeJS, Golang, gopherjs, electron-prebuilt). Example for run app in docker-container in file .drone.yml (config for drone.io CI)

	docker pull arvitaly/electron-go
