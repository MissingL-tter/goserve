# goserve
Minimalist local server for quick web testing without domain errors.
This does nothing but loading the directory for the browser to interpret. I have found it is sufficient for most static testing purposes.


### Opens the working directory on port 8080
```bash
$ goserve
```

### Opens the working directory on some port
```bash
$ goserve :port
```

### Opens the listed directory on port 8080
```bash
$ goserve /some/absolute/path
```

### Opens the listed directory on appended port
```bash
$ goserve /some/absolute/path:port
```
