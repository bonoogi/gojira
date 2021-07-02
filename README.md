# gojira

## 사용법

```
JIRA 이슈를 입력해서 해당하는 웹페이지를 열어줍니다. Ex) $ gojira BAR-123

Usage:
   gojira <issue-num> {flags}
   gojira <command> {flags}

Commands: 
   get-base                      현재 저장된 base url을 알려줍니다.
   help                          displays usage informationn
   set-base                      JIRA 이슈를 열기 위해 필요한 Base URL을 지정합니다. EX) gojira set-base https://foo.atlassian.net
   version                       displays version number

Arguments: 
   issue-num                     열고자 하는 JIRA 이슈 넘버. Ex) BAR-123

Flags: 
   -h, --help                    displays usage information of the application or a command (default: false)
   -v, --version                 displays version number (default: false)
```
