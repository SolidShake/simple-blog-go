package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/SolidShake/simple-blog-go/internal/config"
	"github.com/SolidShake/simple-blog-go/internal/db"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	posts, _ := db.GetAllPosts()

	data := struct {
		Posts []db.Post
	}{
		posts,
	}
	tmpl, _ := template.ParseFiles(
		"templates/base.html",
		"templates/elements/header.html",
		"templates/elements/blog.html",
		"templates/elements/footer.html",
	)
	tmpl.Execute(w, data)
}

func BlogPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// validate postid
	post, _ := db.GetPostByID(ps.ByName("postID"))
	//

	tmpl, _ := template.ParseFiles(
		"templates/base.html",
		"templates/elements/header.html",
		"templates/elements/post.html",
		"templates/elements/comments.html",
		"templates/elements/footer.html",
	)

	tmpl.Execute(w, post)
}

func Tags(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	posts, _ := db.GetAllPostsByTag(ps.ByName("tag"))

	log.Println(posts)

	data := struct {
		Posts []db.Post
	}{
		posts,
	}
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, data)
}

// func Admin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	username, password, ok := r.BasicAuth()
// 	log.Println(username, password, ok)
// 	tmpl, _ := template.ParseFiles("templates/index.html")
// 	tmpl.Execute(w, nil)
// }

func BasicAuth(h httprouter.Handle, requiredUser, requiredPassword string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && user == requiredUser && password == requiredPassword {
			// Delegate request to the given handle
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

func Admin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tags, _ := db.GetAllTags()
	posts, _ := db.GetAllPosts()

	data := struct {
		Tags  []db.Tag
		Posts []db.Post
	}{
		tags,
		posts,
	}

	for _, v := range posts {
		fmt.Println(v.PostID)
		fmt.Println(v.Content)
	}

	fmt.Println("----")

	log.Println(data)
	tmpl, _ := template.ParseFiles("templates/admin/admin.html")
	tmpl.Execute(w, data)
}

// Post handlers

func CreatePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl, _ := template.ParseFiles("templates/admin/post.html")
	tmpl.Execute(w, nil)
}

func CreatePostForm(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	postTitle := r.FormValue("title")
	postContent := r.FormValue("content")
	db.CreatePost(postTitle, postContent)

	http.Redirect(w, r, "/admin", 302)
}

func EditPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl, _ := template.ParseFiles("templates/admin/post.html")
	tmpl.Execute(w, nil)
}

func EditPostForm(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl, _ := template.ParseFiles("templates/admin/post.html")
	tmpl.Execute(w, nil)
}

// Tag handlers

func CreateTag(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl, _ := template.ParseFiles("templates/admin/tag.html")
	tmpl.Execute(w, nil)
}

func CreateTagForm(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tagName := r.FormValue("tagName")
	log.Println(tagName)
	db.CreateTag(tagName)

	http.Redirect(w, r, "/admin", 302)
}

func EditTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tagID, _ := strconv.Atoi(ps.ByName("id"))
	tag, _ := db.GetTagById(tagID)

	tmpl, _ := template.ParseFiles("templates/admin/tag.html")
	tmpl.Execute(w, tag)
}

func EditTagForm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tagID, _ := strconv.Atoi(ps.ByName("id"))
	tagName := r.FormValue("tagName")
	db.EditTag(tagID, tagName)

	http.Redirect(w, r, "/admin", 302)
}

func main() {
	// user := "user"
	// pass := "pass"

	config.InitConfig()
	db.InitDB(config.Cnf)

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/blog/:postID/", BlogPost)
	router.GET("/tags/:tag/", Tags)
	// /tags/:tag
	// /blog/:blogname

	// Admin
	router.GET("/admin", Admin)

	// Posts
	router.GET("/admin/post/create", CreatePost)
	router.POST("/admin/post/create", CreatePostForm)
	router.GET("/admin/post/edit/:id", EditPost)
	router.POST("/admin/post/edit/:id", EditPostForm)

	// Tags
	router.GET("/admin/tag/create", CreateTag)
	router.POST("/admin/tag/create", CreateTagForm)
	router.GET("/admin/tag/edit/:id", EditTag)
	router.POST("/admin/tag/edit/:id", EditTagForm)

	// router.GET("/protected/", BasicAuth(Protected, user, pass))
	// router.GET("/protected/page", BasicAuth(Protected, user, pass))
	router.ServeFiles("/static/*filepath", http.Dir("static"))

	log.Println("Server started at port: " + config.Cnf.Server.Port)
	log.Fatal(http.ListenAndServe(":"+config.Cnf.Server.Port, router))
}
