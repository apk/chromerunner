A tool to open URLs remotely in chrome, like from unix machines on your local windows desktop.

=== Build

    GOOS=windows GOARCH=386 go build chromerunner.go

=== Run daemon

    C:\Users\akrey\chromerunner.exe --path "C:\Program Files (x86)\Google\Chrome\Application\chrome.exe"

=== Open $URL

    curl -d "$URL" http://localhost:3045/start
