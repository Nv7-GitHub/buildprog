# buildprog
Get a progress bar for your Go builds! 

> :warning: This uses `go build`'s `-a` flag. This means it re-compiles everything every time, so this will most likely not build nearly as fast as just plain `go build`. I would only advise using this for understanding how fast your builds are, not for actual use.

This works by using `go build`'s `-x` flag. Read more [here](https://maori.geek.nz/how-go-build-works-750bb2ba6d8e)! 

You can use `-cleancache` to have it refresh the cache. This is useful if you have added a dependency, since sometimes the automatic updating doesn't work.

You can use `-h` to get help. It is just:
```
buildprog -cleancache <Arguments to go build>
```
