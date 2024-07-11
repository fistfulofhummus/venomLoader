#!/bin/zsh

# Check if exactly one argument is provided
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <new_url>"
    exit 1
fi

file_name="main.go"

# Checks if file exists
if [ ! -f "$file_name" ]; then
    echo "File $file_name not found!"
    exit 1
fi

sed -i "s|url:=.*|url:=\"$new_url\"|g" "$file_name"

export GOOS=windows
export GOARCH=amd64
go build
echo "[+] The loader is ready"
