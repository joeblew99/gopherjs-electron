package main

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"

	"github.com/joeblew99/gopherjs-electron"
)

func main() {

	fmt.Println("starting frontend... ")
	var app = electron.GetApp()
	fmt.Println("app path: " + app.GetAppPath())

	app.OnWillQuit(func(event *js.Object) {
		fmt.Println("app quited: ")
	})

	app.OnReady(func() {
		var w = electron.NewBrowserWindow(&map[string]interface{}{"index": "index"})

		fmt.Println("window title: " + w.GetTitle())

		//w.LoadURL("file:///"+js.Global.Get("process").Call("cwd").String()+"index.html", nil)
		w.LoadURL("file:///"+app.GetAppPath()+"/index.html", nil)

		w.GetWebContents().OnDidStopLoading(func() {
			fmt.Println("w Loaded. Title: " + w.GetTitle())
		})

		w.OnClosed(func() {
			fmt.Println("w Closed.")

		})

		w.GetWebContents().OnWillNavigate(func(event *js.Object, url string) {
			fmt.Println("w OnWillNavigate.")
			//w.Close()
		})

		var w1 = electron.NewBrowserWindow(&map[string]interface{}{"index1": "index1"})
		fmt.Println("window1 title: " + w1.GetTitle())
		w1.LoadURL("file:///"+app.GetAppPath()+"/index.1.html", nil)

		w1.GetWebContents().OnDidStopLoading(func() {
			fmt.Println("w1 Loaded. Title: " + w1.GetTitle())
			w1.Destroy()
		})

		w1.OnClosed(func() {
			fmt.Println("w1 Closed.")
			w1.Destroy()
		})

	})

}

/*
var _ = jasmine.Run(func() {
	jasmine.Describe("App", func() {
		var app = electron.GetApp()
		jasmine.ItAsync("OnReady", func(done func()) {
			app.OnReady(func() {
				done()
			})
		})
		jasmine.It("GetAppPath", func() {
			jasmine.Expect(app.GetAppPath() != "").ToBeTruthy()
		})
		jasmine.ItAsync("OnWillQuit", func(done func()) {
			var firstTime = true
			app.OnWillQuit(func(event *js.Object) {
				if firstTime {
					event.Call("preventDefault")
				}
				firstTime = false
				done()
			})
			var br = electron.NewBrowserWindow(nil)
			br.Close()
		})
	})
})
*/
