package main

import (
	"html/template"
	"log"
	"os"
	"time"

	"rayjun.cn/godemo/the_go_programming_language/ch4/github"
)

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var html = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align:left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Items}}
<tr>
 <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
 <td>{{.State}}</td>
 <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
 <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {

	result, err := github.SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	if err := html.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
