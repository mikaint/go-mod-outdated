[![Build Status](https://travis-ci.com/psampaz/go-mod-outdated.svg?branch=master)](https://travis-ci.com/psampaz/go-mod-outdated)
[![codecov](https://codecov.io/gh/psampaz/go-mod-outdated/branch/master/graph/badge.svg)](https://codecov.io/gh/psampaz/go-mod-outdated)
[![Go Report Card](https://goreportcard.com/badge/github.com/psampaz/go-mod-outdated)](https://goreportcard.com/report/github.com/psampaz/go-mod-outdated)
[![GoDoc](https://godoc.org/github.com/psampaz/go-mod-outdated?status.svg)](https://godoc.org/github.com/psampaz/go-mod-outdated)


# go-mod-outdated

An easy way to find outdated dependencies of your Go projects.

go-mod-outdated provides a table view of the **go list -u -m -json all** command which lists all dependencies of a Go project and their available minor and patch updates. It also provides a way to filter indirect dependencies and dependencies without updates.

In short it turns this:

```
{
	"Path": "github.com/BurntSushi/locker",
	"Version": "v0.0.0-20171006230638-a6e239ea1c69",
	"Time": "2017-10-06T23:06:38Z",
	"GoMod": "/home/mojo/go/pkg/mod/cache/download/github.com/!burnt!sushi/locker/@v/v0.0.0-20171006230638-a6e239ea1c69.mod"
}
{
	"Path": "github.com/BurntSushi/toml",
	"Version": "v0.0.0-20170626110600-a368813c5e64",
	"Time": "2017-06-26T11:06:00Z",
	"Update": {
		"Path": "github.com/BurntSushi/toml",
		"Version": "v0.3.1",
		"Time": "2018-08-15T10:47:33Z"
	},
	"GoMod": "/home/mojo/go/pkg/mod/cache/download/github.com/!burnt!sushi/toml/@v/v0.0.0-20170626110600-a368813c5e64.mod"
}
```

into this

```
+-------------------------------------------+--------------------------------------+------------------------------------+--------+
|                  MODULE                   |               VERSION                |            NEW VERSION             | DIRECT |
+-------------------------------------------+--------------------------------------+------------------------------------+--------+
| github.com/BurntSushi/locker              | v0.0.0-20171006230638-a6e239ea1c69   |                                    | true   |
| github.com/BurntSushi/toml                | v0.0.0-20170626110600-a368813c5e64   | v0.3.1                             | true   |
+-------------------------------------------+--------------------------------------+------------------------------------+--------+
```
## Installation

```
go get -u github.com/psampaz/go-mod-outdated
```
## Usage
In the folder where your go.mod lives run
 
```
go list -u -m -json all | go-mod-outdated
```

to see all modules in table view.

If you want to see only the modules with updates run 

```
go list -u -m -json all | go-mod-outdated -update
```

If you want to see only the direct depedencies run

```
go list -u -m -json all | go-mod-outdated -direct
```

If you want to see only the direct depedencies that have updates run

```
go list -u -m -json all | go-mod-outdated -update -direct 
```

### Help
  
In order to see details about the usage of the command use the **-h** or **-help** flag

```
$ go-mod-outdated -help

Usage of go-mod-outdated:
  -direct
        List only direct modules
  -update
        List only modules with updates
```

### Shortcut

If **go list -u -m -json all | go-mod-outdated -update -direct** seems too difficult to use or remember you can create 
a shortcut using an alias. In linux try one of the following: 

```
alias gmo="go list -u -m -json all | go-mod-outdated"

alias gmod="go list -u -m -json all | go-mod-outdated -direct"

alias gmou="go list -u -m -json all | go-mod-outdated -update"

alias gmodu="go list -u -m -json all | go-mod-outdated -direct -update"
```  

## Real Example

The following example is based on Hugo's go.mod (v0.53) (https://raw.githubusercontent.com/gohugoio/hugo/v0.53/go.mod)

### Json output of go list -u -m json all command    

```
$ go list -u -m -json all
{
	"Path": "github.com/gohugoio/hugo",
	"Main": true,
	"Dir": "/home/mojo/Code/go/hugo",
	"GoMod": "/home/mojo/Code/go/hugo/go.mod"
}
{
	"Path": "github.com/BurntSushi/locker",
	"Version": "v0.0.0-20171006230638-a6e239ea1c69",
	"Time": "2017-10-06T23:06:38Z",
	"GoMod": "/home/mojo/go/pkg/mod/cache/download/github.com/!burnt!sushi/locker/@v/v0.0.0-20171006230638-a6e239ea1c69.mod"
}
{
	"Path": "github.com/BurntSushi/toml",
	"Version": "v0.0.0-20170626110600-a368813c5e64",
	"Time": "2017-06-26T11:06:00Z",
	"Update": {
		"Path": "github.com/BurntSushi/toml",
		"Version": "v0.3.1",
		"Time": "2018-08-15T10:47:33Z"
	},
	"GoMod": "/home/mojo/go/pkg/mod/cache/download/github.com/!burnt!sushi/toml/@v/v0.0.0-20170626110600-a368813c5e64.mod"
}
{
	"Path": "github.com/PuerkitoBio/purell",
	"Version": "v1.1.0",
	"Time": "2016-11-15T02:49:42Z",
	"Update": {
		"Path": "github.com/PuerkitoBio/purell",
		"Version": "v1.1.1",
		"Time": "2019-02-16T21:19:35Z"
	},
	"GoMod": "/home/mojo/go/pkg/mod/cache/download/github.com/!puerkito!bio/purell/@v/v1.1.0.mod"
}
.....
.....
.....
.....
{
	"Path": "golang.org/x/sys",
	"Version": "v0.0.0-20181206074257-70b957f3b65e",
	"Time": "2018-12-06T07:42:57Z",
	"Update": {
		"Path": "golang.org/x/sys",
		"Version": "v0.0.0-20190418153312-f0ce4c0180be",
		"Time": "2019-04-18T15:33:12Z"
	},
	"Indirect": true,
	"GoMod": "/home/mojo/go/pkg/mod/cache/download/golang.org/x/sys/@v/v0.0.0-20181206074257-70b957f3b65e.mod"
}
{
	"Path": "golang.org/x/text",
	"Version": "v0.3.0",
	"Time": "2017-12-14T13:08:43Z",
	"GoMod": "/home/mojo/go/pkg/mod/cache/download/golang.org/x/text/@v/v0.3.0.mod"
}
{
	"Path": "gopkg.in/check.v1",
	"Version": "v1.0.0-20180628173108-788fd7840127",
	"Time": "2018-06-28T17:31:08Z",
	"Indirect": true,
	"GoMod": "/home/mojo/go/pkg/mod/cache/download/gopkg.in/check.v1/@v/v1.0.0-20180628173108-788fd7840127.mod"
}
{
	"Path": "gopkg.in/yaml.v2",
	"Version": "v2.2.2",
	"Time": "2018-11-15T11:05:04Z",
	"GoMod": "/home/mojo/go/pkg/mod/cache/download/gopkg.in/yaml.v2/@v/v2.2.2.mod"
}
```

### Table view of go list -u -m -json all command using go-mod-outdated
```
$ go list -u -m -json all | go-mod-outdated
+-------------------------------------------+--------------------------------------+------------------------------------+--------+
|                  MODULE                   |               VERSION                |            NEW VERSION             | DIRECT |
+-------------------------------------------+--------------------------------------+------------------------------------+--------+
| github.com/BurntSushi/locker              | v0.0.0-20171006230638-a6e239ea1c69   |                                    | true   |
| github.com/BurntSushi/toml                | v0.0.0-20170626110600-a368813c5e64   | v0.3.1                             | true   |
| github.com/PuerkitoBio/purell             | v1.1.0                               | v1.1.1                             | true   |
| github.com/PuerkitoBio/urlesc             | v0.0.0-20170810143723-de5bf2ad4578   |                                    | false  |
| github.com/alecthomas/assert              | v0.0.0-20170929043011-405dbfeb8e38   |                                    | true   |
| github.com/alecthomas/chroma              | v0.6.0                               | v0.6.3                             | true   |
| github.com/alecthomas/colour              | v0.0.0-20160524082231-60882d9e2721   |                                    | false  |
| github.com/alecthomas/repr                | v0.0.0-20180818092828-117648cd9897   | v0.0.0-20181024024818-d37bc2a10ba1 | false  |
| github.com/armon/consul-api               | v0.0.0-20180202201655-eb2c6b5be1b6   |                                    | false  |
| github.com/bep/debounce                   | v1.1.0                               | v1.2.0                             | true   |
| github.com/bep/gitmap                     | v1.0.0                               |                                    | true   |
| github.com/bep/go-tocss                   | v0.6.0                               |                                    | true   |
| github.com/chaseadamsio/goorgeous         | v1.1.0                               |                                    | true   |
| github.com/cheekybits/is                  | v0.0.0-20150225183255-68e9c0620927   |                                    | false  |
| github.com/coreos/etcd                    | v3.3.10+incompatible                 | v3.3.12+incompatible               | false  |
| github.com/coreos/go-etcd                 | v2.0.0+incompatible                  |                                    | false  |
| github.com/coreos/go-semver               | v0.2.0                               | v0.3.0                             | false  |
| github.com/cpuguy83/go-md2man             | v1.0.8                               | v1.0.10                            | false  |
| github.com/danwakefield/fnmatch           | v0.0.0-20160403171240-cbb64ac3d964   |                                    | false  |
| github.com/davecgh/go-spew                | v1.1.1                               |                                    | false  |
| github.com/disintegration/imaging         | v1.5.0                               | v1.6.0                             | true   |
| github.com/dlclark/regexp2                | v1.1.6                               |                                    | false  |
| github.com/dustin/go-humanize             | v1.0.0                               |                                    | false  |
| github.com/eknkc/amber                    | v0.0.0-20171010120322-cdade1c07385   |                                    | true   |
| github.com/fortytw2/leaktest              | v1.2.0                               | v1.3.0                             | true   |
| github.com/fsnotify/fsnotify              | v1.4.7                               |                                    | true   |
| github.com/gobwas/glob                    | v0.2.3                               |                                    | true   |
| github.com/gorilla/websocket              | v1.4.0                               |                                    | true   |
| github.com/hashicorp/go-immutable-radix   | v1.0.0                               |                                    | true   |
| github.com/hashicorp/go-uuid              | v1.0.0                               | v1.0.1                             | false  |
| github.com/hashicorp/golang-lru           | v0.5.0                               | v0.5.1                             | false  |
| github.com/hashicorp/hcl                  | v1.0.0                               |                                    | false  |
| github.com/inconshreveable/mousetrap      | v1.0.0                               |                                    | false  |
| github.com/jdkato/prose                   | v1.1.0                               |                                    | true   |
| github.com/kr/pretty                      | v0.1.0                               |                                    | false  |
| github.com/kr/pty                         | v1.1.1                               | v1.1.4                             | false  |
| github.com/kr/text                        | v0.1.0                               |                                    | false  |
| github.com/kyokomi/emoji                  | v1.5.1                               | v2.1.0+incompatible                | true   |
| github.com/magefile/mage                  | v1.4.0                               | v1.8.0                             | true   |
| github.com/magiconair/properties          | v1.8.0                               |                                    | false  |
| github.com/markbates/inflect              | v1.0.0                               | v1.0.4                             | true   |
| github.com/matryer/try                    | v0.0.0-20161228173917-9ac251b645a2   |                                    | false  |
| github.com/mattn/go-isatty                | v0.0.4                               | v0.0.7                             | true   |
| github.com/mattn/go-runewidth             | v0.0.3                               | v0.0.4                             | false  |
| github.com/miekg/mmark                    | v1.3.6                               |                                    | true   |
| github.com/mitchellh/hashstructure        | v1.0.0                               |                                    | true   |
| github.com/mitchellh/mapstructure         | v1.1.2                               |                                    | true   |
| github.com/muesli/smartcrop               | v0.0.0-20180228075044-f6ebaa786a12   | v0.2.0                             | true   |
| github.com/nfnt/resize                    | v0.0.0-20180221191011-83c6a9932646   |                                    | false  |
| github.com/nicksnyder/go-i18n             | v1.10.0                              |                                    | true   |
| github.com/olekukonko/tablewriter         | v0.0.0-20180506121414-d4647c9c7a84   | v0.0.1                             | true   |
| github.com/pelletier/go-toml              | v1.2.0                               | v1.3.0                             | false  |
| github.com/pkg/errors                     | v0.8.0                               | v0.8.1                             | true   |
| github.com/pmezard/go-difflib             | v1.0.0                               |                                    | false  |
| github.com/russross/blackfriday           | v0.0.0-20180804101149-46c73eb196ba   | v2.0.0+incompatible                | true   |
| github.com/sanity-io/litter               | v1.1.0                               |                                    | true   |
| github.com/sergi/go-diff                  | v1.0.0                               |                                    | false  |
| github.com/shurcooL/sanitized_anchor_name | v0.0.0-20170918181015-86672fcb3f95   | v1.0.0                             | false  |
| github.com/spf13/afero                    | v1.2.0                               | v1.2.2                             | true   |
| github.com/spf13/cast                     | v1.3.0                               |                                    | true   |
| github.com/spf13/cobra                    | v0.0.3                               |                                    | true   |
| github.com/spf13/fsync                    | v0.0.0-20170320142552-12a01e648f05   |                                    | true   |
| github.com/spf13/jwalterweatherman        | v1.0.1-0.20181028145347-94f6ae3ed3bc | v1.1.0                             | true   |
| github.com/spf13/nitro                    | v0.0.0-20131003134307-24d7ef30a12d   |                                    | true   |
| github.com/spf13/pflag                    | v1.0.3                               |                                    | true   |
| github.com/spf13/viper                    | v1.3.1                               | v1.3.2                             | true   |
| github.com/stretchr/testify               | v1.2.3-0.20181014000028-04af85275a5c | v1.3.0                             | true   |
| github.com/tdewolff/minify/v2             | v2.3.7                               | v2.4.0                             | true   |
| github.com/tdewolff/parse/v2              | v2.3.5                               | v2.3.6                             | false  |
| github.com/tdewolff/test                  | v1.0.0                               |                                    | false  |
| github.com/ugorji/go/codec                | v0.0.0-20181204163529-d75b2dcb6bc8   |                                    | false  |
| github.com/wellington/go-libsass          | v0.9.3-0.20181113175235-c63644206701 | v0.9.2                             | false  |
| github.com/xordataexchange/crypt          | v0.0.3-0.20170626215501-b2862e3d0a77 | v0.0.2                             | false  |
| github.com/yosssi/ace                     | v0.0.5                               |                                    | true   |
| golang.org/x/crypto                       | v0.0.0-20181203042331-505ab145d0a9   | v0.0.0-20190418165655-df01cb2cc480 | false  |
| golang.org/x/image                        | v0.0.0-20180708004352-c73c2afc3b81   | v0.0.0-20190417020941-4e30a6eb7d9a | true   |
| golang.org/x/net                          | v0.0.0-20180906233101-161cd47e91fd   | v0.0.0-20190415214537-1da14a5a36f2 | false  |
| golang.org/x/sync                         | v0.0.0-20180314180146-1d60e4601c6f   | v0.0.0-20190412183630-56d357773e84 | true   |
| golang.org/x/sys                          | v0.0.0-20181206074257-70b957f3b65e   | v0.0.0-20190418153312-f0ce4c0180be | false  |
| golang.org/x/text                         | v0.3.0                               |                                    | true   |
| gopkg.in/check.v1                         | v1.0.0-20180628173108-788fd7840127   |                                    | false  |
| gopkg.in/yaml.v2                          | v2.2.2                               |                                    | true   |
+-------------------------------------------+--------------------------------------+------------------------------------+--------+

```

### Table view of go list -u -m -json all command using go-mod-outdated (only dependencies with updates)

```
$ go list -u -m -json all | go-mod-outdated -update
+-------------------------------------------+--------------------------------------+------------------------------------+--------+
|                  MODULE                   |               VERSION                |            NEW VERSION             | DIRECT |
+-------------------------------------------+--------------------------------------+------------------------------------+--------+
| github.com/BurntSushi/toml                | v0.0.0-20170626110600-a368813c5e64   | v0.3.1                             | true   |
| github.com/PuerkitoBio/purell             | v1.1.0                               | v1.1.1                             | true   |
| github.com/alecthomas/chroma              | v0.6.0                               | v0.6.3                             | true   |
| github.com/alecthomas/repr                | v0.0.0-20180818092828-117648cd9897   | v0.0.0-20181024024818-d37bc2a10ba1 | false  |
| github.com/bep/debounce                   | v1.1.0                               | v1.2.0                             | true   |
| github.com/coreos/etcd                    | v3.3.10+incompatible                 | v3.3.12+incompatible               | false  |
| github.com/coreos/go-semver               | v0.2.0                               | v0.3.0                             | false  |
| github.com/cpuguy83/go-md2man             | v1.0.8                               | v1.0.10                            | false  |
| github.com/disintegration/imaging         | v1.5.0                               | v1.6.0                             | true   |
| github.com/fortytw2/leaktest              | v1.2.0                               | v1.3.0                             | true   |
| github.com/hashicorp/go-uuid              | v1.0.0                               | v1.0.1                             | false  |
| github.com/hashicorp/golang-lru           | v0.5.0                               | v0.5.1                             | false  |
| github.com/kr/pty                         | v1.1.1                               | v1.1.4                             | false  |
| github.com/kyokomi/emoji                  | v1.5.1                               | v2.1.0+incompatible                | true   |
| github.com/magefile/mage                  | v1.4.0                               | v1.8.0                             | true   |
| github.com/markbates/inflect              | v1.0.0                               | v1.0.4                             | true   |
| github.com/mattn/go-isatty                | v0.0.4                               | v0.0.7                             | true   |
| github.com/mattn/go-runewidth             | v0.0.3                               | v0.0.4                             | false  |
| github.com/muesli/smartcrop               | v0.0.0-20180228075044-f6ebaa786a12   | v0.2.0                             | true   |
| github.com/olekukonko/tablewriter         | v0.0.0-20180506121414-d4647c9c7a84   | v0.0.1                             | true   |
| github.com/pelletier/go-toml              | v1.2.0                               | v1.3.0                             | false  |
| github.com/pkg/errors                     | v0.8.0                               | v0.8.1                             | true   |
| github.com/russross/blackfriday           | v0.0.0-20180804101149-46c73eb196ba   | v2.0.0+incompatible                | true   |
| github.com/shurcooL/sanitized_anchor_name | v0.0.0-20170918181015-86672fcb3f95   | v1.0.0                             | false  |
| github.com/spf13/afero                    | v1.2.0                               | v1.2.2                             | true   |
| github.com/spf13/jwalterweatherman        | v1.0.1-0.20181028145347-94f6ae3ed3bc | v1.1.0                             | true   |
| github.com/spf13/viper                    | v1.3.1                               | v1.3.2                             | true   |
| github.com/stretchr/testify               | v1.2.3-0.20181014000028-04af85275a5c | v1.3.0                             | true   |
| github.com/tdewolff/minify/v2             | v2.3.7                               | v2.4.0                             | true   |
| github.com/tdewolff/parse/v2              | v2.3.5                               | v2.3.6                             | false  |
| github.com/wellington/go-libsass          | v0.9.3-0.20181113175235-c63644206701 | v0.9.2                             | false  |
| github.com/xordataexchange/crypt          | v0.0.3-0.20170626215501-b2862e3d0a77 | v0.0.2                             | false  |
| golang.org/x/crypto                       | v0.0.0-20181203042331-505ab145d0a9   | v0.0.0-20190418165655-df01cb2cc480 | false  |
| golang.org/x/image                        | v0.0.0-20180708004352-c73c2afc3b81   | v0.0.0-20190417020941-4e30a6eb7d9a | true   |
| golang.org/x/net                          | v0.0.0-20180906233101-161cd47e91fd   | v0.0.0-20190415214537-1da14a5a36f2 | false  |
| golang.org/x/sync                         | v0.0.0-20180314180146-1d60e4601c6f   | v0.0.0-20190412183630-56d357773e84 | true   |
| golang.org/x/sys                          | v0.0.0-20181206074257-70b957f3b65e   | v0.0.0-20190418153312-f0ce4c0180be | false  |
+-------------------------------------------+--------------------------------------+------------------------------------+--------+
```
### Table view of go list -u -m -json all command using go-mod-outdated (only direct dependencies with updates)

```
$ go list -u -m -json all | go-mod-outdated -update -direct
+------------------------------------+--------------------------------------+------------------------------------+--------+
|               MODULE               |               VERSION                |            NEW VERSION             | DIRECT |
+------------------------------------+--------------------------------------+------------------------------------+--------+
| github.com/BurntSushi/toml         | v0.0.0-20170626110600-a368813c5e64   | v0.3.1                             | true   |
| github.com/PuerkitoBio/purell      | v1.1.0                               | v1.1.1                             | true   |
| github.com/alecthomas/chroma       | v0.6.0                               | v0.6.3                             | true   |
| github.com/bep/debounce            | v1.1.0                               | v1.2.0                             | true   |
| github.com/disintegration/imaging  | v1.5.0                               | v1.6.0                             | true   |
| github.com/fortytw2/leaktest       | v1.2.0                               | v1.3.0                             | true   |
| github.com/kyokomi/emoji           | v1.5.1                               | v2.1.0+incompatible                | true   |
| github.com/magefile/mage           | v1.4.0                               | v1.8.0                             | true   |
| github.com/markbates/inflect       | v1.0.0                               | v1.0.4                             | true   |
| github.com/mattn/go-isatty         | v0.0.4                               | v0.0.7                             | true   |
| github.com/muesli/smartcrop        | v0.0.0-20180228075044-f6ebaa786a12   | v0.2.0                             | true   |
| github.com/olekukonko/tablewriter  | v0.0.0-20180506121414-d4647c9c7a84   | v0.0.1                             | true   |
| github.com/pkg/errors              | v0.8.0                               | v0.8.1                             | true   |
| github.com/russross/blackfriday    | v0.0.0-20180804101149-46c73eb196ba   | v2.0.0+incompatible                | true   |
| github.com/spf13/afero             | v1.2.0                               | v1.2.2                             | true   |
| github.com/spf13/jwalterweatherman | v1.0.1-0.20181028145347-94f6ae3ed3bc | v1.1.0                             | true   |
| github.com/spf13/viper             | v1.3.1                               | v1.3.2                             | true   |
| github.com/stretchr/testify        | v1.2.3-0.20181014000028-04af85275a5c | v1.3.0                             | true   |
| github.com/tdewolff/minify/v2      | v2.3.7                               | v2.4.0                             | true   |
| golang.org/x/image                 | v0.0.0-20180708004352-c73c2afc3b81   | v0.0.0-20190417020941-4e30a6eb7d9a | true   |
| golang.org/x/sync                  | v0.0.0-20180314180146-1d60e4601c6f   | v0.0.0-20190412183630-56d357773e84 | true   |
+------------------------------------+--------------------------------------+------------------------------------+--------+
```
