package main

import (
	"fmt"
	"os"
	"time"

	"errors"
	"github.com/saucelabs/selenium"
)

func main() {
	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	const (
		// These paths will be different on your system.
		seleniumPath    = ""
		chromeDriverPath = ""
	)
	var (
		port            = 443
		region		= os.Getenv("SAUCE_REGION")
		protocol	= "https"
		ondemand	= fmt.Sprintf("ondemand.%s.saucelabs.com", region)
		username	= os.Getenv("SAUCE_USERNAME")
		accessKey	= os.Getenv("SAUCE_ACCESS_KEY")
	)
	if region == "" {
		panic("Please set up SAUCE_REGION")
	}
	if username == "" {
		panic("Please set up SAUCE_USERNAME")
	}
	if accessKey == "" {
		panic("Please set up SAUCE_ACCESS_KEY")
	}

	selenium.SetDebug(true)
	caps := selenium.Capabilities{"browserName": "chrome", "platformName": "windows 10"}
	hub := fmt.Sprintf("%s://%s:%s@%s:%d/wd/hub", protocol, username, accessKey, ondemand, port)
	wd, err := selenium.NewRemote(caps, hub)
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to the simple playground interface.
	if err := wd.Get("https://www.saucedemo.com"); err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second)
	el, err := wd.FindElement(selenium.ByID, "user-name")
	if err != nil {
		panic(err)
	}
	err = el.Click()
	if err != nil {
		panic(err)
	}
	err = el.SendKeys("standard_user")
	if err != nil {
		panic(err)
	}
	el, err = wd.FindElement(selenium.ByID, "password")
	if err != nil {
		panic(err)
	}
	err = el.Click()
	if err != nil {
		panic(err)
	}
	err = el.SendKeys("secret_sauce")
	if err != nil {
		panic(err)
	}
	el, err = wd.FindElement(selenium.ByID, "login-button")
	if err != nil {
		panic(err)
	}
	err = el.Click()
	if err != nil {
		panic(err)
	}
	url, err := wd.CurrentURL()
	if err != nil {
		panic(err)
	}
	time.Sleep(5 * time.Second)
	if url != "https://www.saucedemo.com/inventory.html" {
		panic(errors.New("Wrong url"))
	}
}

