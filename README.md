# Utility for e2guardian setup

## Build
```
go build utils.go
```

This can be used for the following:
./utils -generate_X25519_key="true" -to_encode="some text"

The above command generate 'X25519' key pair and prints on the console.
Also the base64 encoded version of 'some text' is printed on the screen
