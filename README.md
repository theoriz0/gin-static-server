# gin-static-server
A very simple static file server built with github.com/gin-gonic/gin

## Usage
Run the executable to serve the dir where the executable locates.

For example in Windows:
```bash
.\gin-static-server.exe
```

Then visit ip:port according to the output
```
>> Serving {the folder served}
>> Port 8080
```

If the port is used by other program, it will auto-retry display the port tried to use and the final port used.
```
>> Serving {the folder served}
>> Port 8080 Failed
>> Port 8081
```