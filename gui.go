package main

import (
	"github.com/andlabs/ui"
	"time"
	"github.com/yosssi/gohtml"
)

// todo: add post requests

// fixme: add color to response

func openGui() {
	err := ui.Main(func() {
		url := ui.NewEntry()
		send := ui.NewButton("Send request")
		response := ui.NewLabel("")

		box := ui.NewVerticalBox()

		box.Append(ui.NewLabel("Url:"), false)
		box.Append(url, false)
		box.Append(send, false)
		box.Append(response, false)

		box.SetPadded(true)

		send.OnClicked(func(*ui.Button) {
			requestUrl := url.Text()

			if requestUrl != "" {
				startTime := time.Now()

				httpResp, err := getRequest(requestUrl)

				output, json := parseResponse(httpResp, err, startTime)

				beautified := "Parse error: "

				if json {
					json, err := prettyJson(output[4])

					if err == nil {
						beautified = json
					} else {
						// fixme: err to string
					}

				} else {
					beautified = gohtml.Format(output[4])
				}


				response.SetText(output[0]+"\n"+output[1]+"\n"+output[2]+"\n"+output[3]+"\n"+beautified)

			} else {
				response.SetText("You have to set a url")
			}

		})

		window := ui.NewWindow("RESTfulTest", 500, 300, false)
		window.SetChild(box)
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		window.Show()
	})

	if err != nil {
		sendError(err)
	}
}
