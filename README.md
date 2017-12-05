# goproxy-get
Tool to run "go get" with proxy address:port.

# Usage
goproxy-get is just a small and simple tool like "go get ...". Truely, it just wraped "go get" and the only purpose of it is to use "go get" through http proxy. 

**Of course, I will not provide the proxy tools and server.**

# Why I developed this stuff?
For some reason, I can't "go get" the packages from "google.golang.org", "golang.org" and so on because of network issues. And that bothers me a lot.

# How to use it
Simple.

```bash
goproxy-get -p proxy_address:port package_url
```

> Default value of "-p" is "127.0.0.1:1087". Why? Because it's convenient for me. Want to change the default value, you can edit the go files and build yourself.

for example:
```bash
goproxy-get -p 127.0.0.1:1087 google.golang.org/grpc
```

Also, you can still use the arguments of "go get". Including "-d", "-v", "u" and so on.

> Be aware of that "build flags" is not supported yet.