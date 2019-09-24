package common

import (
	"fmt"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/gavv/httpexpect"
	"net/http"
)

func PermissionTest(e *httpexpect.Expect, sesId *http.Cookie) {

	fmt.Println()
	printlnWithColor("Permission", "blue")
	fmt.Println("============================")

	// show

	printlnWithColor("show", "green")
	e.GET(config.Get().Url("/info/permission")).
		WithCookie(sesId.Name, sesId.Value).
		Expect().
		Status(200).
		Body().Contains("Dashboard").Contains("All permission")

	// show new form

	printlnWithColor("show new form", "green")
	formBody := e.GET(config.Get().Url("/info/permission/new")).
		WithCookie(sesId.Name, sesId.Value).
		Expect().Status(200).Body()

	token := reg.FindStringSubmatch(formBody.Raw())

	// new permission tester

	printlnWithColor("new permission tester", "green")
	res := e.POST(config.Get().Url("/new/permission")).
		WithCookie(sesId.Name, sesId.Value).
		WithMultipart().
		WithFormField("http_method[]", "GET").
		WithForm(map[string]interface{}{
			"name": "tester",
			"slug": "tester",
			"http_path": `/
/admin/info/op`,
			"_previous_": config.Get().Url("/info/permission?page=1&pageSize=10&sort=id&sort_type=desc"),
			"_t":         token[1],
		}).Expect().Status(200)
	res.Header("X-Pjax-Url").Contains(config.Get().Url("/info/"))
	res.Body().Contains("tester").Contains("GET")

	// show form: without id

	printlnWithColor("show form: without id", "green")
	e.GET(config.Get().Url("/info/permission/edit")).
		WithCookie(sesId.Name, sesId.Value).
		Expect().Status(200).Body().Contains("wrong id")

	// show form

	printlnWithColor("show form", "green")
	formBody = e.GET(config.Get().Url("/info/permission/edit")).
		WithQuery("id", "1").
		WithCookie(sesId.Name, sesId.Value).
		Expect().Status(200).Body()

	token = reg.FindStringSubmatch(formBody.Raw())

	// edit form

	printlnWithColor("edit form", "green")
	res = e.POST(config.Get().Url("/edit/permission")).
		WithCookie(sesId.Name, sesId.Value).
		WithMultipart().
		WithFormField("http_method[]", "GET").
		WithFormField("http_method[]", "POST").
		WithForm(map[string]interface{}{
			"name": "tester",
			"slug": "tester",
			"http_path": `/
/admin/info/op`,
			"_previous_": config.Get().Url("/info/permission?page=1&pageSize=10&sort=id&sort_type=desc"),
			"_t":         token[1],
			"id":         "3",
		}).Expect().Status(200)
	res.Header("X-Pjax-Url").Contains(config.Get().Url("/info/"))
	res.Body().Contains("tester").Contains("GET,POST")

	printlnWithColor("delete", "green")
	e.GET("/pong").Expect().Status(404)
}
