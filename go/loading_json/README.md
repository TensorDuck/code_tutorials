Some lessons learned: the json Unmarshall and Marshall requires all the struct methods
to be public. Thus it uses "tags" enclosed in \`. Struct tags are meant to be things other
modules can look up, but they themselves will have no effect on the code.

some useful links
https://tutorialedge.net/golang/parsing-json-with-golang/

https://blog.golang.org/json

https://gobyexample.com/json

https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go