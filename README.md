# licenser
A tool for looking up the licenses for any dependencies you are shipping with your (Go) apps

# Installation & usage

```
go get -u github.com/njern/licenser
licenser -d /folder/you/want/to/scan/for/licenses
```

# Example 

```
$ licenser -d ~/go_workspace/src/github.com/njern/licenser
/Users/njern/go_workspace/src/github.com/njern/licenser => MIT
/Users/njern/go_workspace/src/github.com/njern/licenser/vendor/github.com/ryanuber/go-license => MIT
/Users/njern/go_workspace/src/github.com/njern/licenser/vendor/github.com/ryanuber/go-license/fixtures/licenses => Unlicense
```
