package test_tool

import (
	"fmt"
	"os"
	"text/template"
)

// template 將文本嵌入其他文本中
// 由 parse 定義文本之模板 {{.}}為嵌入當前對象值 {{.$name}}為嵌入變數名為$name之值
// execute 將data丟入解析好的文本執行，將輸出寫入 wr
func Texttemplate() {
	fmt.Println("text template test start:")

	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	t1 = template.Must(t1.Parse("Value is {{.}}\n"))

	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	t2 := Create("t2", "Name : {{.Name}}\n")
	t2.Execute(os.Stdout, struct {
		Name string
	}{"Brian Chiang"})

	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")

	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	t4 := Create("t4",
		"Range : {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})

	fmt.Println("text template test end")
}
