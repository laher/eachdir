# eachdir

run a command on each directory specified by env variable EACHDIR

EACHDIR can be a series of globs or space-delimited directory names. It skips files

```
$ EACHDIR="*" eachdir ls
$ EACHDIR="a b c" eachdir tree -d
$ EACHDIR="a b c" eachdir ./build.sh
```

## Why?

Well, sometimes `make` is hard. Also sometimes shell scripting is hard. This is just a simple and dumb utility
