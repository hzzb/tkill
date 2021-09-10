# tkill

Similar to kill(1),  but send signal to a thread instead of a process.


# install

go get github.com/hzzb/tkill


# usage

```
tkill  [-s signal_name]  pid  tid
tkill  -signal_name      pid  tid
tkill  -signal_number    pid  tid
```
