# gopherjs-electron
Atom/Electron (https://github.com/atom/electron/) desktop apps with Go. Package use https://github.com/gopherjs/gopherjs.



# Install

Package require electron-prebuilt npm-module

	npm install -g electron-prebuilt

	go get github.com/arvitaly/gopherjs-electron

# Usage

Look in cmd for basic example

# TODO
- When Electron starts, start / spawn the golang backend.
- Get communication between backend and frontend working.



# Test

For tests used Jasmine (http://jasmine.github.io/) and adapter fo go https://github.com/arvitaly/gopherjs-jasmine

	go get github.com/arvitaly/gopherjs-jasmine
	npm install electron-prebuilt -g
	npm install
	npm test

# Docker

Also you can use docker-image for electron (includes NodeJS, Golang, gopherjs, electron-prebuilt). Example for run app in docker-container in file .drone.yml (config for drone.io CI)

	docker pull arvitaly/electron-go
