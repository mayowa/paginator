package paginator

import "fmt"

func ExampleTest() {
	fmt.Println(`
		<link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap.min.css" rel="stylesheet">
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js"></script>
		<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/js/bootstrap.min.js"></script>
		<div class="container">
		`)
	for i := 0; i < 15; i++ {
		fmt.Println("<div class=row>", BootstrapPaginator(i, 10, 100, 2, "?page=#"), "</div>")
	}

	fmt.Println("<div class=row>", BootstrapPaginator(1, 10, 5, 2, "?page=#"), "</div>")
	fmt.Println("<div class=row>", BootstrapPaginator(1, 10, 19, 2, "?page=#"), "</div>")
	fmt.Println("<div class=row>", BootstrapPaginator(2, 10, 19, 2, "?page=#"), "</div>")

}
