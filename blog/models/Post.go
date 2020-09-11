package models


// Object or Structure
type Post struct {

   // Properties
   Id string
   Title string
   ContentHtml string
   ContentMarkdown string
}


// Constructor
func NewPost(id, title, contentHtml, contentMarkdown string) *Post{
    return &Post{id, title, contentHtml, contentMarkdown}
}