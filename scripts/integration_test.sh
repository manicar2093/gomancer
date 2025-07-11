#!/bin/bash
task install
cd $HOME
gomancer new github.com/manicar2093/gomancer_test
cd $HOME/gomancer_test
go mod tidy
gomancer gen Entity arg1:string arg2:string:optional arg3:time:optional arg4:string arg5:time
go tool templ generate --proxy="http://localhost:8090" --open-browser=false -v
tailwindcss -i ./cmd/service/sources/css/input.css -o ./cmd/service/assets/css/styles.css
go mod tidy
task run &
sleep 2
if http --check-status --ignore-stdin --timeout=2.5 GET http://localhost:3000/api/v1/initial &> /dev/null; then
    echo
    echo
    echo 'SUCCESS!'
    lsof -t -i:3000 | xargs kill -9
    rm -r $HOME/gomancer_test
    exit 0
else
    case $? in
        2) echo 'Request timed out!' ;;
        3) echo 'Unexpected HTTP 3xx Redirection!' ;;
        4) echo 'HTTP 4xx Client Error!' ;;
        5) echo 'HTTP 5xx Server Error!' ;;
        6) echo 'Exceeded --max-redirects=<n> redirects!' ;;
        *) echo 'Other Error!' ;;
    esac
    lsof -t -i:3000 | xargs kill -9
    rm -r $HOME/gomancer_test
    exit 1
fi
