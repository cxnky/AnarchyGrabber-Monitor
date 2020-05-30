package threads

import (
	"fmt"
	"github.com/cxnky/AnarchyGrabber-Monitor/util"
	"github.com/gen2brain/beeep"
	"io/ioutil"
	"time"
)

func StartFileMonitoring() {
	util.Info("Starting file monitor")
	go monitor()
}

func monitor() {
	for {
		util.Info("Checking for AnarchyGrabber")
		appDataPath := util.FetchAppDataPath()
		discordModulesDirectory := fmt.Sprintf("%s\\Discord\\0.0.306\\modules\\discord_desktop_core", appDataPath)
		indexJsLocation := fmt.Sprintf("%s\\index.js", discordModulesDirectory)

		if !util.FileExists(indexJsLocation) {
			util.Error("Could not find index.js in the discord_desktop_core directory!")
			return
		}

		indexJsContents, err := ioutil.ReadFile(indexJsLocation)
		util.CheckError(err)

		indexJsContentString := string(indexJsContents)

		if indexJsContentString != "module.exports = require('./core.asar');" {
			util.Fatal("WARNING: You may be potentially infected by AnarchyGrabber!!!")
			util.Fatal("index.js contents are as follows:")
			util.Fatal(indexJsContentString)

			// show a notification to the user, in case they have the window minimised
			err := beeep.Alert("Warning", "You may be potentially infected by AnarchyGrabber! Please re-install Discord as soon as possible", "assets/error.png")
			util.CheckError(err)

		} else {
			util.Info("File contents are as expected, whew.")
		}

		time.Sleep(30 * time.Second)
	}
}
