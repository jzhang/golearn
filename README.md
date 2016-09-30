
# learning golang by coding

## using gdb debugger with go 

[using gdb debugger with go] (https://blog.codeship.com/using-gdb-debugger-with-go/)
[godco/gdb] (https://golang.org/doc/gdb)

* when debugging, pass the flags -gcflags "-N -l" to the go command used to build the code being debugged
* add this line	" add-auto-load-safe-path /root/go/src/runtime/runtime-gdb.py " to your configuration file "/root/.gdbinit"

## benchmark 
