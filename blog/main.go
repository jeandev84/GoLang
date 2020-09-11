package main


import (
   "fmt" // для вывода
   "html/template" // работа с шаблоном
   "net/http" // работа с http
   "models/models" // modules
   "crypto/rand" // random
)


// Хранить наши посты в памяти
var posts map[string]*models.Post


// Вывести список постов
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
    // t.ExecuteTemplate(w, "index", nil)
    t.ExecuteTemplate(w, "index", posts) // posts [ текущий контекст ]
}


// Добавление поста
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


// Редактирование поста
func editHandler(w http.ResponseWriter, r *http.Request) {
    /* fmt.Fprintf(w, "<h1>Hello world</h1>") */

    // ParseFiles возвращает наш template и ошибку
    t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")

    // Выводим ошибку если она есть (nil ?)
    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    // прочитать id из URL
    id := r.FormValue("id")

    // ищем пост (может возвращать post или статус что не нашел)
    post, found := posts[id]

    // если не нашел пост то мы сделаем редирект страница не найдена
    if !found {
       http.NotFound(w, r)
    }

    // Выполняем наш template ( как render() в php )
    // а если нашел "post" мы передаем в template
    t.ExecuteTemplate(w, "write", post)
}



// Удаление поста
func deleteHandler(w http.ResponseWriter, r *http.Request) {

    // прочитать id из URL
    id := r.FormValue("id")

    if id == "" {
       http.NotFound(w, r)
    }

    // удаление из постов
    delete(posts, id)

    // редирект на главную страницу
    http.Redirect(w, r, "/", 302)
}


// Cохранение поста
func savePostHandler(w http.ResponseWriter, r *http.Request) {

    /*
      если мы не хотим объявлять переменную тогда пишем нижнее почеркивание _
      _ := r.FormValue("id")
      _ := r.FormValue("title")
      _ := r.FormValue("content")
    */

    /*
    id := GenerateId()
    title := r.FormValue("title")
    content := r.FormValue("content")

    post := models.NewPost(id, title, content)

    // Запишем наш пост в мап(map) постов
    posts[post.Id] = post


    // Редирект ( redirect permanently )
    http.Redirect(w, r, "/", 302)
    */

    // Получаем данные из формы
    id := r.FormValue("id")
    title := r.FormValue("title")
    content := r.FormValue("content")

    // обявляем post
    var post *models.Post

    // если id не пустой, значить мы редактируем
    if id != "" {
       post = posts[id]
       post.Title = title
       post.Content = content
    } else {
       id  = GenerateId() // сгенируем новый id
       post := models.NewPost(id, title, content)
       posts[post.Id] = post
    }

    // redirect permanently
    http.Redirect(w, r, "/", 302)
}


//------------- Helpers ---------------//
// Метод для генерации массив bytes
func GenerateId() string {

   // сгенируем наш random числа 16 bytes
   b := make([]byte, 16)

   // считаем наш random
   rand.Read(b)

   // выводим
   return fmt.Sprintf("%x", b)
}

//----------- End Helpers --------------//


// Entry Point of application
func main() {

   // Выводим в консоле что мы слушаем в порте 3000
   fmt.Println("Listening on port :3000")

   // объявим наши посты (posts)
   posts = make(map[string]*models.Post, 0) // нуль элементов


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
   http.HandleFunc("/edit", editHandler)
   http.HandleFunc("/delete", deleteHandler)
   http.HandleFunc("/save", savePostHandler)

   // Слушаем в порте 3000
   http.ListenAndServe(":3000", nil)
}