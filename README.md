# Paginator for Go (incl. Bootstrap)

```sh
go get github.com/slava-vishnyakov/paginator
```

## Raw data

```go
import "github.com/slava-vishnyakov/paginator"

paginator.Paginator(1, 10, 45, 2)
// returns []PaginatorPage raw data (see below)
```

- `page` (in example: `1`) is 1-based

- `perPage` (in example: `10`)

- `count` (total items to list: `45`)

- `interval` (in example: `2`) shows how many pages around current to show
  - `1 ... 6 [7] 8 ... 20` is interval = 1
  - `1 ... 5 6 [7] 8 9 ... 20` is interval = 2


```go
type PaginatorPage struct {
	Label      string
	Page       int
	IsActive   bool
	IsDisabled bool
}
```

## Bootstrap

```go
paginator.BootstrapPaginator(1, 10, 5, 2, "?page=#")
// returns HTML string
```

