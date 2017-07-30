# nano
Nanosecond command operating on cross platform  
#making it as a hobby!!

## Download and installation
### Download
https://github.com/wablerfam/nano/releases  
### installation 
No need. There is a binary that can be executed immediately in the decompressed compressed file

## Example
### Get current time in nanoseconds
    /go # nano
    2017-07-29T15:44:11.617736606Z

## Options
### Measure the command execution time in nanoseconds
    /go # nano ls
    README.md
    main.go
    
    time: 9034835 ns

    /go # nano ls -l
    -rw-r--r--    1 foo     bar           515 Jul 30 15:23 README.md
    drwxr-xr-x    4 foo     bar           136 Jul 30 09:48 nano

    time: 8412035 ns