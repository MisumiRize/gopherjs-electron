package browser

import "github.com/gopherjs/gopherjs/js"

type App struct {
	*js.Object
}

func NewApp() App {
	return App{js.Global.Get("require").Invoke("app")}
}

func (a App) On(event string, fn func(...*js.Object)) App {
	return App{a.Call("on", event, fn)}
}

type BrowserWindow struct {
	*js.Object
	Id int `js:"id"`
}

func NewBrowserWindow(options map[string]interface{}) *BrowserWindow {
	return &BrowserWindow{Object: js.Global.Get("require").Invoke("browser-window").New(js.Global.Get("Object").New(options))}
}

func (b BrowserWindow) LoadURL(url string) {
	b.Call("loadUrl", url)
}

func (b BrowserWindow) AddListener(event string, fn func(...*js.Object)) WebContents {
	return WebContents{b.Call("addListener", event, fn)}
}

func (b BrowserWindow) On(event string, fn func(...*js.Object)) WebContents {
	return WebContents{b.Call("on", event, fn)}
}

func (b BrowserWindow) Once(event string, fn func(...*js.Object)) WebContents {
	return WebContents{b.Call("once", event, fn)}
}

func (b BrowserWindow) Center() {
	b.Call("center")
}

type WebContents struct {
	*js.Object
}

func (b BrowserWindow) WebContents() WebContents {
	return WebContents{b.Get("webContents")}
}

func (w WebContents) GetURL() string {
	return w.Call("getUrl").String()
}

func (w WebContents) On(event string, fn func(...*js.Object)) WebContents {
	return WebContents{w.Call("on", event, fn)}
}
