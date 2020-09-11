package main


import (
   "fmt" // для вывода
   "net/http" // работа с http
   "html/template"
   "models/models" // modules
   "github.com/codegangsta/martini" // framework
   "github.com/martini-contrib/render" // middleware martini
   //"github.com/russross/blackfriday" // Convert markdown
   "github.com/gomarkdown/markdown" // go-markdown
   //"github.com/gomarkdown/markdown/parser"
   "github.com/gomarkdown/markdown/html"
   "crypto/rand" // random
)


/* Обявление константы */
var posts map[string]*models.Post
var counter int


//----------- Функция Handlers -------//
func indexHandler(rnd render.Render) {

    fmt.Println(counter)

    // render
    rnd.HTML(200, "index", posts)
}


func writeHandler(rnd render.Render) {

   // render
   rnd.HTML(200, "write", nil)
}


func editHandler(rnd render.Render, r *http.Request, params martini.Params) {

    // читаем id из URL
    // id := r.FormValue("id")
    id := params["id"] // получить параметра id

    // ищем пост (может возвращать post или статус что не нашел)
    post, found := posts[id]

    // если не нашел пост то мы сделаем редирект страница не найдена
    if !found {
       rnd.Redirect("/") // redirect на главную
       return
    }

    rnd.HTML(200, "write", post)
}


func savePostHandler(rnd render.Render, r *http.Request) {

    // Получаем данные из формы
    id := r.FormValue("id")
    title := r.FormValue("title")
    contentMarkdown := r.FormValue("content")
    // contentHtml := string(blackfriday.MarkdownBasic([]byte(contentMarkdown)))
    contentHtml := ""


    // обявляем post
    var post *models.Post

    // если id не пустой, значить мы редактируем
    if id != "" {
       post = posts[id]
       post.Title = title
       post.ContentHtml = contentHtml
       post.ContentMarkdown = contentMarkdown
    } else {
       id  = GenerateId() // сгенируем новый id
       post := models.NewPost(id, title, contentHtml, contentMarkdown)
       posts[post.Id] = post
    }

    rnd.Redirect("/")
}


func deleteHandler(rnd render.Render, r *http.Request, params martini.Params) {

    // прочитать id из URL
    id := params["id"]

    if id == "" {
       rnd.Redirect("/")
       return
    }

    // удаление из постов
    delete(posts, id)

    // редирект на главную страницу
    rnd.Redirect("/")
}



// handler via request ajax
func getHtmlHandler(rnd render.Render, r *http.Request) {

   htmlFlags := html.CommonFlags | html.HrefTargetBlank
   opts := html.RendererOptions{Flags: htmlFlags}
   renderer := html.NewRenderer(opts)

   // Получаем наш markdown
   md_form := r.FormValue("md")

   // Возвращает массив байтов
   // htmlBytes := blackfriday.MarkdownBasic([]byte(md))
   md := []byte(md_form)
   htmlBytes := markdown.ToHTML(md, nil, renderer)


   // Возвращает массив байтов
   // htmlBytes := blackfriday.MarkdownBasic([]byte(md))


   // Отдаем в JSON Format
   rnd.JSON(200, map[string]interface{} {"html": string(htmlBytes)})
}




//------------- Начало Helpers ---------------//
// Метод для генерации массив bytes
func GenerateId() string {

   // сгенируем наш random числа 16 bytes
   b := make([]byte, 16)

   // считаем наш random
   rand.Read(b)

   // выводим
   return fmt.Sprintf("%x", b)
}


func unescape(x string) interface{} {
    return template.HTML(x)
}


//----------- Конец Helpers --------------//




// Entry Point of application
func main() {

   // Выводим в консоле что мы слушаем в порте 3000
   fmt.Println("Listening on port :3000")


   // объявим наши посты (posts)
   posts = make(map[string]*models.Post, 0) // нуль элементов
   counter = 0

   // Исползуем фреймворк "martini"
   m := martini.Classic()

   unescapeFuncMap := template.FuncMap{"unescape": unescape}


   // Middleware ( https://github.com/martini-contrib/render )
   // Он позволяет работать с template (html) и json
   // m.Use - это как наше контейнер (Container Dependency Injection)
   m.Use(render.Renderer(render.Options{
     Directory: "templates", // Specify what path to load the templates from.
     Layout: "layout", // Specify a layout template. Layouts can call {{ yield }} to render the current template.
     Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
     Funcs: []template.FuncMap{unescapeFuncMap}, // Specify helper function maps for templates to access.
     //Delims: render.Delims{"{[{", "}]}"}, // Sets delimiters to the specified strings.
     Charset: "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
     IndentJSON: true, // Output human readable JSON
     //IndentXML: true, // Output human readable XML
     //HTMLContentType: "application/xhtml+xml", // Output XHTML content type instead of default "text/html"
   }))



   // Увеличиваем counter для теста
   /*
   m.Use(func(r *http.Request) {
       if r.URL.Path == "/write" {
          counter++
       }
   })
   */

   // Подключение обычные статичные файлы (m.Use(martini.Static("assets", ...)) без опций)
   staticOptions := martini.StaticOptions{Prefix:"assets"}
   m.Use(martini.Static("assets", staticOptions))

   // Регистрируем наши маршруты
   m.Get("/", indexHandler)
   m.Get("/write", writeHandler)
   m.Get("/edit/:id", editHandler)
   m.Get("/delete/:id", deleteHandler)
   m.Post("/save", savePostHandler)
   m.Post("/gethtml", getHtmlHandler)

   /*
   m.Get("/test", func() string {
       return "test"
   })
   */


   // Запускаем приложение в порте :3000
   m.Run()
}
