package models


// Object or Structure
type Post struct {

   // Properties
   Id string
   Title string
   Content string
}


// Constructor
func NewPost(id, title, content string) *Post{
    return &Post{id, title, content}
}