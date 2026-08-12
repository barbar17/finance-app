#!/usr/bin/env bash

show_spinner() {
    local pid=$1
    local message=$2
    local frames="/-\|"

    while kill -0 "$pid" 2>/dev/null; do
        for (( i=0; i<${#frames}; i++ )); do
            if ! kill -0 "$pid" 2>/dev/null; then break; fi
            
            printf "\r%s... [%c]" "$message" "${frames:$i:1}"
            sleep 0.1
        done
    done
    
    tput cnorm
    printf "\r%s... Done!\n" "$message"
}

go build -o ./bin/main ./cmd/server &
BUILD_PID=$!

show_spinner "$BUILD_PID" "Building app"

echo "Finished building app. Launching server..."
./bin/main