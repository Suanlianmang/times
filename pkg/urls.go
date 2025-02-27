
package pkg

// Define the struct with exported fields
type URLToTemplate struct {
    URL      string
    Template string
}

// Define the slice of URLToTemplate structs
var IndexUrl = URLToTemplate {URL: "/", Template: "index.html"}
var PostUrl = URLToTemplate {URL: "/post", Template: "post.html"}
var AuthUrl = URLToTemplate {URL: "/auth", Template: "auth.html"}

