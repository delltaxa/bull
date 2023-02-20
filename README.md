
# Bull

Bull is an simple reverse shell manager with a async http server.

## Installation

Install Golang from https://go.dev/

Bull Installation script
```bash
git clone https://github.com/delltaxa/bull
cd bull/
```
    
## Build

Build bull by simply running 
```bash
go build .
```
Note that you have to be in the source directory.
## Usage/Examples

Running with default settings
```bash
./bull
```

Configuring Server
```bash
./bull <local_addr>:<local_port>
```

or (this will run the server on all devices)

```bash
./bull :<local_port>
```
## Showcase

Generating a payload
![image](https://user-images.githubusercontent.com/114283067/220091299-421feaa2-4f41-4acb-a84e-602b164d039a.png)<br />

Display sessions
![image](https://user-images.githubusercontent.com/114283067/220092636-7349d470-85d1-463e-9585-fe416446ba09.png)<br />

Getting a shell
![image](https://user-images.githubusercontent.com/114283067/220093738-605d2222-aedb-437d-9019-daec351067a4.png)
