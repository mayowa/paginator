package paginator

import (
	"math"
	"strconv"
	"strings"
)

type PaginatorPage struct {
	Label      string
	Page       int
	IsActive   bool
	IsDisabled bool
}

// page is 1-based
// interval shows how many pages around current to show
// 1 ... 6 [7] 8 ... 20     -- interval = 1
// 1 ... 5 6 [7] 8 9 ... 20 -- interval = 2
func BootstrapPaginator(page, perPage, count, interval int, hrefTemplate string) string {
	res := []string{}
	res = append(res, `<ul class="pagination">`)
	for _, p := range Paginator(page, perPage, count, interval) {
		url := strings.Replace(hrefTemplate, "#", strconv.Itoa(p.Page), 1)

		if p.Label == "previous" {
			res = append(res, strings.Replace(`<li class="page-item"><a class="page-link" href="#" aria-label="Previous"><span aria-hidden="true">&laquo;</span></a></li>`, "#", url, 1))
		} else if p.Label == "next" {
			res = append(res, strings.Replace(`<li class="page-item"><a class="page-link" href="#" aria-label="Next"><span aria-hidden="true">&raquo;</span></a></li>`, "#", url, 1))
		} else {
			h := strings.Replace(strings.Replace(`<li class="page-item"><a class="page-link" href="#">{PAGENUM}</a></li>`, "#", url, 1), "{PAGENUM}", p.Label, 1)
			if p.IsActive {
				h = strings.Replace(h, `class="page-item"`, `class="page-item active"`, 1)
			}
			res = append(res, h)
		}
	}
	res = append(res, `</ul>`)
	return strings.Join(res, "")
}

func Paginator(page, perPage, count, interval int) []PaginatorPage {
	var pages []PaginatorPage

	// count=45
	// lastPage=5
	// page items: 0-9, 10-19, 20-29, 30-39, 40-45

	firstPage := 1
	lastPage := int(math.Ceil(float64(count) / float64(perPage)))

	if page < firstPage {
		page = firstPage
	}

	if page > lastPage {
		page = lastPage
	}

	if page-1 >= firstPage {
		pages = append(pages, PaginatorPage{
			Label: "previous",
			Page:  page - 1,
		})
	}

	pages = append(pages, PaginatorPage{
		Label:    strconv.Itoa(firstPage),
		Page:     firstPage,
		IsActive: firstPage == page,
	})

	if page-interval > firstPage+1 {
		if page-interval-1 == firstPage+1 {
			pages = append(pages, PaginatorPage{
				Label:      strconv.Itoa(firstPage + 1),
				Page:       firstPage + 1,
				IsActive:   firstPage+1 == page,
				IsDisabled: true,
			})
		} else {
			pages = append(pages, PaginatorPage{
				Label:      "...",
				Page:       0,
				IsDisabled: true,
			})
		}
	}

	for p := page - interval; p <= page+interval; p++ {
		if p > firstPage && p < lastPage {
			pages = append(pages, PaginatorPage{
				Label:    strconv.Itoa(p),
				Page:     p,
				IsActive: p == page,
			})
		}
	}

	if page+interval < lastPage-1 {
		if page+interval+1 == lastPage-1 {
			pages = append(pages, PaginatorPage{
				Label:      strconv.Itoa(lastPage - 1),
				Page:       lastPage - 1,
				IsDisabled: true,
			})
		} else {
			pages = append(pages, PaginatorPage{
				Label:      "...",
				Page:       0,
				IsDisabled: true,
			})
		}
	}

	if lastPage > firstPage {
		pages = append(pages, PaginatorPage{
			Label:    strconv.Itoa(lastPage),
			Page:     lastPage,
			IsActive: lastPage == page,
		})
	}

	if page+1 <= lastPage {
		pages = append(pages, PaginatorPage{
			Label: "next",
			Page:  page + 1,
		})
	}

	return pages
}
