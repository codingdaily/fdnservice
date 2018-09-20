# fire-starter golang http boiler plate code.
how to use:

on project root you can do : 
```make``` to build binary
```make clean``` to clean binary
```make test``` to do test
```make docker``` to build docker image

used libraries:
- testing:
    - gomega 
    - ginkgo
- routing:
    - gorilla/mux
- command line & config:
    - spf13/cobra
    - spf13/viper
- Logging:
    - zap
- Error Trace:
    - raven-go (sentry)
- RPC:
    - protocol buffer

## What to Configure.
when you using the make command, it will read the Makefile
change following lines so you'll have correct binary name.
APP_NAME : the binary name
USER : user on cvs (git)
DOMAIN : dsn of the source.

all three info will be used when the binary is generated.