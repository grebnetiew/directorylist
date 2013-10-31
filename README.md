Directorylist
=============

Small program to print an indented recursive directory listing

Usage
-----
Invoke the program (after building, for example with `go build filesys.go`), using

    filesys.exe "C:\Somedir"

or

    ./filesys "/etc"

How is this different from path/filepath/Walk?
----------------------------------------------

1. Fancy printing with brackets and stuff
2. Directories are sorted before files, whereas Walk sorts only lexicographically
