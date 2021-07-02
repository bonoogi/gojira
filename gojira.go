package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"

	"github.com/thatisuday/commando"
)

func main() {
	var baseurlFilename = "baseurl"
	var errorNoBaseURL = "지정된 Base URL이 없습니다. `gojira set-base` 로 지라 Base URL을 지정해주세요."
	// https://github.com/thatisuday/commando
	commando.
		SetExecutableName("gojira").
		SetVersion("v1.0.1").
		SetDescription("JIRA 이슈를 입력해서 해당하는 웹페이지를 열어줍니다. Ex) $ gojira BAR-123")

	commando.
		Register(nil).
		AddArgument("issue-num", "열고자 하는 JIRA 이슈 넘버. Ex) BAR-123", "").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {

			data, err := ioutil.ReadFile(baseurlFilename)
			if err != nil {
				fmt.Println(errorNoBaseURL)
				fmt.Println(err)
				return
			}

			// https://golang.org/pkg/os/exec/
			// https://gist.github.com/hyg/9c4afcd91fe24316cbf0
			var issue = args["issue-num"].Value

			matched, _ := regexp.MatchString("[A-Za-z]{2,}-\\d+", issue)
			if matched {
				var err error
				var url = string(data) + "/browse/" + issue
				switch runtime.GOOS {
				case "linux":
					err = exec.Command("xdg-open", url).Start()
				case "windows":
					err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
				case "darwin":
					err = exec.Command("open", url).Start()
				default:
					err = fmt.Errorf("unsupported platform")
				}
				if err != nil {
					// http://pyrasis.com/book/GoForTheReallyImpatient/Unit60
					log.Fatal(err)
				}
			} else {
				fmt.Printf("%v는 올바른 이슈 넘버가 아닙니다.\n", issue)
			}
		})

	commando.
		Register("set-base").
		AddArgument("base-url", "JIRA 이슈를 열기 위한 Base URL. Ex) https://foo.atlassian.net", "").
		SetDescription("JIRA 이슈를 열기 위해 필요한 Base URL을 지정합니다. EX) gojira set-base https://foo.atlassian.net").
		SetShortDescription("JIRA 이슈를 열기 위해 필요한 Base URL을 지정합니다. EX) gojira set-base https://foo.atlassian.net").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			var url = args["base-url"].Value
			err := ioutil.WriteFile(baseurlFilename, []byte(url), os.FileMode(0644))
			if err != nil {
				fmt.Println(err)
				return
			}

			data, err := ioutil.ReadFile(baseurlFilename)
			if err != nil {
				fmt.Println(err)
				return
			}
			var resultMessage = "JIRA Base URL " + string(data) + " 가 지정되었습니다."
			fmt.Println(resultMessage)
		})

	commando.
		Register("get-base").
		SetDescription("현재 저장된 base url을 알려줍니다.").
		SetShortDescription("현재 저장된 base url을 알려줍니다.").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			data, err := ioutil.ReadFile(baseurlFilename)
			if err != nil {
				fmt.Println(errorNoBaseURL)
				fmt.Println(err)
				return
			}
			fmt.Println("현재 지정된 Base URL은" + string(data) + " 입니다.")
		})
	commando.Parse(nil)
}
