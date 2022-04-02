# funcs

`funcs` displays the exported Golang functions with their number of occurrences.

## Installation

```
go install github.com/xpetit/funcs@latest
```

## Usage

```shell
funcs "$(go env GOROOT)/src" | sort -n | tail -n 20
```

> ```
>   39 Swap
>   39 Truncate
>   40 Add
>   40 Reset
>   42 GoString
>   42 Size
>   44 Stat
>   45 Set
>   46 Seek
>   55 Open
>   55 SetLen
>   61 End
>   65 Pos
>   67 Len
>  113 Write
>  124 Read
>  141 Error
>  144 Close
>  260 String
> 9209 exported functions
> ```
