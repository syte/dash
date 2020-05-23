# dash
SimpleHTTPServer that let's you mock requests quickly

## Run
from the dash folder, `run go build`. Once you have the binary you can run `./dash -h` to get the help test instructing you on how to use the application. If you specify both the cert and key it'll run on an https server. If either or none of these parameters are specified it'll run as insecure http.

## Security
I expect this to be used locally so security considerations are ignored. For example, you can easily exploit a path traversal issue but I didn't see this as a big deal for testing applications locally.
