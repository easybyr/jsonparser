```go
import (
    "github.com/easybyr/jsonparser"
)

var config = jsonconfig.New("./conf/app.json")
var run_mode = config.Get("run_mode")
var write_timeout = config.Get("server.write_timeout")
```

```json
// ./conf/app.json

{
    "run_mode": "debug",
    "server": {
        "http_port": 8000,
        "read_timeout": 60,
        "write_timeout": 60
    }  
}
```