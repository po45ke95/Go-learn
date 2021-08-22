package main
import (
	"log"
	"net/http"
	"html/template"
)

type IndexData struct {
	Title   string
	Content string
}

func test(w http.ResponseWriter, r *http.Request)  {
	//w.WriteHeader(http.StatusOK)
	//w.Write([]byte(`my first website`))
	//str := `<!DOCTYPE html>

	//<html>
	//<head><title>首頁</title></head>
	//<body><h1>首頁</h1><p>我的第一個首頁</p></body>
	//</html>
	//`
	//w.Write([]byte(str))

	tmpl := template.Must(template.ParseFiles("./index.html"))
	data := new(IndexData)
	data.Title = "首頁"
	data.Content= "我的第一個首頁"
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", test)
	http.HandleFunc("/index", test)
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Fatal("ListenAndServe: ", err)
	}
}
