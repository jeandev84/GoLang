package main


import (
   "fmt" // для вывода
   "html/template" // работа с шаблоном
   "net/http" // работа с http
   "github.com/gavruk/go-blog-example/models"
)


// хранить наши посты в памяти
var posts map[string]*models.Post


// Route Handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
    /* fmt.Fprintf(w, "<h1>Hello world</h1>") */

    // ParseFiles возвращает наш template и ошибку
    t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

    // Выводим ошибку если она есть (nil ?)
    if err != nil {
        fmt.Fprintf(w, err.Error())
    }


    // Вывести наши посты
    fmt.Println(posts)

    // Выполняем наш template ( как render() в php )
    // index - это название нашего вида для этого указали {{ define "index" }}
    t.ExecuteTemplate(w, "index", nil)
}


func writeHandler(w http.ResponseWriter, r *http.Request) {
    /* fmt.Fprintf(w, "<h1>Hello world</h1>") */

    // ParseFiles возвращает наш template и ошибку
    t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")

    // Выводим ошибку если она есть (nil ?)
    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    // Выполняем наш template ( как render() в php )
    // write - это название нашего вида для этого указали {{ define "write" }}
    t.ExecuteTemplate(w, "write", nil)
}


func savePostHandler(w http.ResponseWriter, r *http.Request) {

    /*
      если мы не хотим объявлять переменную тогда пишем нижнее почеркивание _
      _ := r.FormValue("id")
      _ := r.FormValue("title")
      _ := r.FormValue("content")
    */

    // r - Request
    // получаем данные из формы
    id := r.FormValue("id")
    title := r.FormValue("title")
    content := r.FormValue("content")

    // create a new post model
    post := models.NewPost(id, title, content)

    // redirect permanently
    http.Redirect(w, r, "/", 302)
}


// Entry Point of application
func main() {

   // Выводим в консоле что мы слушаем в порте 3000
   fmt.Println("Listening on port :3000")

   // объявим наши посты (posts)
   posts = make(map[string]*models.Post, 0) // нуль элементов
   posts[post.Id] = post


   // Подключение стили и скрипты resources
   // http://example.com/assets/css/app.css
   // убрезаем папку assets с помощью http.StripPrefix(), и мы ищем в папке "assets"
   // мы имем стили в папке "assets"  /css/app.css
   http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))


   // Регистрируем наши маршруты
   /*
   http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
       fmt.Fprintf(w, "<h1>Hello world</h1>")
   })
   */

   http.HandleFunc("/", indexHandler)
   http.HandleFunc("/write", writeHandler)
   http.HandleFunc("/SavePost", savePostHandler)

   // Слушаем в порте 3000
   http.ListenAndServe(":3000", nil)
}