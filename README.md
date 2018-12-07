# lumbarcheck

This is a tool that check if you can connect some databases, URLs, or another datasources.

## How to use this

Most easy way to use lumbarcheck, is downloading from [releases](https://github.com/deadcheat/lumbarcheck/releases) directly.

lumbarcheck command needs an argument at least.

```
$ lumbarcheck --help
usage: main [<flags>] <Data Source>

Flags:
      --help          Show context-sensitive help (also try --help-long and --help-man).
  -t, --type="mysql"  Kind of Data Source, listed below

                        mysql connect to mysql database
                        http connect to given URL with GET Method

Args:
  <Data Source>  Data Source Name you'll try to connect to
```

