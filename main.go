package main

import (
	"context"
	"day-11/connection"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

type MetaData struct {
	Title     string
	IsLogin   bool
	UserName  string
	Password  string
	FlashData string
}

var Data = MetaData{
	Title: "web",
}

type Blog struct {
	Id           int
	Name         string
	Start_date   time.Time
	End_date     time.Time
	Format_start string
	Format_end   string
	Duration     string
	Description  string
	Technologies string
	Image        string
	IsLogin      bool
}

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

var Blogs = []Blog{
	/*{
		Id:           0,
		Name:         "Dumbways mobile app-2021",
		Start_date:   "2022-10-17",
		End_date:     "2022-10-24",
		Duration:     "1 Weeks",
		Description:  "Test",
		Technologies: "Node Js",
	},
	{
		Id:           1,
		Name:         "Dumbways mobile app-2021",
		Start_date:   "2022-10-17",
		End_date:     "2022-10-24",
		Duration:     "1 Weeks",
		Description:  "Test",
		Technologies: "Node Js",
	},*/
}

func main() {
	route := mux.NewRouter()

	connection.DatabaseConnect()

	// route path folder untuk public
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	//routing
	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/blog", blog).Methods("GET")
	route.HandleFunc("/blog-detail/{id}", blogDetail).Methods("GET")
	route.HandleFunc("/form-blog", formAddBlog).Methods("GET")
	route.HandleFunc("/add-blog", addBlog).Methods("POST")
	route.HandleFunc("/delete-blog/{id}", deleteBlog).Methods("GET")
	route.HandleFunc("/edit-form-blog/{id}", editForm).Methods("GET")
	route.HandleFunc("/edit-blog/{id}", editBlog).Methods("POST")

	route.HandleFunc("/form-register", formRegister).Methods("GET")
	route.HandleFunc("/register", register).Methods("POST")
	route.HandleFunc("/form-login", formLogin).Methods("GET")
	route.HandleFunc("/login", login).Methods("POST")
	route.HandleFunc("/logout", logout).Methods("GET")

	fmt.Println("Server running on port 5000")
	http.ListenAndServe("localhost:5000", route)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	var store = sessions.NewCookieStore([]byte("SESSION_KEY"))
	session, _ := store.Get(r, "SESSION_KEY")

	if session.Values["IsLogin"] != true {
		Data.IsLogin = false
	} else {
		Data.IsLogin = session.Values["IsLogin"].(bool)
		Data.UserName = session.Values["Name"].(string)
	}

	rows, _ := connection.Conn.Query(context.Background(), "SELECT id, name, start_date, end_date, description, technologies, image FROM tb_projects")

	var result []Blog // array data

	for rows.Next() {
		var each = Blog{} //call struct
		err := rows.Scan(&each.Id, &each.Name, &each.Start_date, &each.End_date, &each.Description, &each.Technologies, &each.Image)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		//each.Author = "Hoki"

		each.Format_start = each.Start_date.Format("02-01-2006")
		each.Format_end = each.End_date.Format("02-01-2006")

		layoutDate := "2006-01-02"
		startParse, _ := time.Parse(layoutDate, each.Start_date.Format("2006-01-02"))
		endParse, _ := time.Parse(layoutDate, each.End_date.Format("2006-01-02"))
		fmt.Println(startParse)

		hour := 1
		day := hour * 24
		week := hour * 24 * 7
		month := hour * 24 * 30
		year := hour * 24 * 365

		differHour := endParse.Sub(startParse).Hours()
		var differHours int = int(differHour)
		// fmt.Println(differHours)
		days := differHours / day
		weeks := differHours / week
		months := differHours / month
		years := differHours / year

		if differHours < week {
			each.Duration = strconv.Itoa(int(days)) + " Days"
		} else if differHours < month {
			each.Duration = strconv.Itoa(int(weeks)) + " Weeks"
		} else if differHours < year {
			each.Duration = strconv.Itoa(int(months)) + " Months"
		} else if differHours > year {
			each.Duration = strconv.Itoa(int(years)) + " Years"
		}

		if session.Values["IsLogin"] != true {
			each.IsLogin = false
		} else {
			each.IsLogin = session.Values["IsLogin"].(bool)
		}

		result = append(result, each)
	}

	fmt.Println(result)

	respData := map[string]interface{}{
		"Data":  Data,
		"Blogs": result,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, respData)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/contact.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func blog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// fmt.Println(Blogs)

	var tmpl, err = template.ParseFiles("views/blog.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	var store = sessions.NewCookieStore([]byte("SESSION_KEY"))
	session, _ := store.Get(r, "SESSION_KEY")

	if session.Values["IsLogin"] != true {
		Data.IsLogin = false
	} else {
		Data.IsLogin = session.Values["IsLogin"].(bool)
		Data.UserName = session.Values["Name"].(string)
	}

	rows, _ := connection.Conn.Query(context.Background(), "SELECT id, title, content, post_date, author FROM tb_blog")

	var result []Blog // array data

	for rows.Next() {
		var each = Blog{} // manggil struct

		err := rows.Scan(&each.Id, &each.Name, &each.Start_date, &each.End_date, &each.Description, &each.Technologies, &each.Image)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		each.Format_start = each.Start_date.Format("02-01-2006")
		each.Format_end = each.End_date.Format("02-01-2006")

		layoutDate := "2006-01-02"
		startParse, _ := time.Parse(layoutDate, each.Start_date.Format("2006-01-02"))
		endParse, _ := time.Parse(layoutDate, each.End_date.Format("2006-01-02"))
		fmt.Println(startParse)

		hour := 1
		day := hour * 24
		week := hour * 24 * 7
		month := hour * 24 * 30
		year := hour * 24 * 365

		differHour := endParse.Sub(startParse).Hours()
		var differHours int = int(differHour)
		// fmt.Println(differHours)
		days := differHours / day
		weeks := differHours / week
		months := differHours / month
		years := differHours / year

		if differHours < week {
			each.Duration = strconv.Itoa(int(days)) + " Days"
		} else if differHours < month {
			each.Duration = strconv.Itoa(int(weeks)) + " Weeks"
		} else if differHours < year {
			each.Duration = strconv.Itoa(int(months)) + " Months"
		} else if differHours > year {
			each.Duration = strconv.Itoa(int(years)) + " Years"
		}

		if session.Values["IsLogin"] != true {
			each.IsLogin = false
		} else {
			each.IsLogin = session.Values["IsLogin"].(bool)
		}

		result = append(result, each)
	}

	fmt.Println(result)

	respData := map[string]interface{}{
		"Data":  Data,
		"Blogs": result,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, respData)
}

func blogDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var tmpl, err = template.ParseFiles("views/blog-detail.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	var BlogDetail = Blog{}

	var store = sessions.NewCookieStore([]byte("SESSION_KEY"))
	session, _ := store.Get(r, "SESSION_KEY")

	if session.Values["IsLogin"] != true {
		Data.IsLogin = false
	} else {
		Data.IsLogin = session.Values["IsLogin"].(bool)
		Data.UserName = session.Values["Name"].(string)
	}

	err = connection.Conn.QueryRow(context.Background(), "SELECT id, name, start_date, end_date, description, technologies, image FROM tb_projects WHERE id=$1", id).Scan(
		&BlogDetail.Id, &BlogDetail.Name, &BlogDetail.Start_date, &BlogDetail.End_date, &BlogDetail.Description, &BlogDetail.Technologies, &BlogDetail.Image)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message: " + err.Error()))
		return
	}
	BlogDetail.Format_start = BlogDetail.Start_date.Format("2 January 2006")
	BlogDetail.Format_end = BlogDetail.End_date.Format("2 January 2006")

	layoutDate := "2006-01-02"
	startParse, _ := time.Parse(layoutDate, BlogDetail.Start_date.Format("2006-01-02"))
	endParse, _ := time.Parse(layoutDate, BlogDetail.End_date.Format("2006-01-02"))
	fmt.Println(startParse)

	hour := 1
	day := hour * 24
	week := hour * 24 * 7
	month := hour * 24 * 30
	year := hour * 24 * 365

	differHour := endParse.Sub(startParse).Hours()
	var differHours int = int(differHour)
	// fmt.Println(differHours)
	days := differHours / day
	weeks := differHours / week
	months := differHours / month
	years := differHours / year

	if differHours < week {
		BlogDetail.Duration = strconv.Itoa(int(days)) + " Days"
	} else if differHours < month {
		BlogDetail.Duration = strconv.Itoa(int(weeks)) + " Weeks"
	} else if differHours < year {
		BlogDetail.Duration = strconv.Itoa(int(months)) + " Months"
	} else if differHours > year {
		BlogDetail.Duration = strconv.Itoa(int(years)) + " Years"
	}

	if session.Values["IsLogin"] != true {
		BlogDetail.IsLogin = false
	} else {
		BlogDetail.IsLogin = session.Values["IsLogin"].(bool)
	}

	data := map[string]interface{}{
		"Blog": BlogDetail,
		"Data": Data,
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, data)

}

func formAddBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/add-blog.html")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func addBlog(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	var name = r.PostForm.Get("inputTitle")
	var start = r.PostForm.Get("inputStart")
	var end = r.PostForm.Get("inputEnd")
	//var each.Duration string
	var description = r.PostForm.Get("inputContent")
	var technologies = r.PostForm.Get("js")
	var image = r.PostForm.Get("inputImage")

	fmt.Println("Name : " + r.PostForm.Get("inputTitle")) // value berdasarkan dari tag input name
	fmt.Println("Start : " + r.PostForm.Get("inputStart"))
	fmt.Println("End : " + r.PostForm.Get("inputEnd"))
	fmt.Println("Description : " + r.PostForm.Get("inputContent"))
	fmt.Println("Technologies : " + r.PostForm.Get("js"))
	fmt.Println("Image : " + r.PostForm.Get("inputImage"))

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_projects(name, start_date, end_date, description, technologies, image) VALUES ($1, $2, $3, $4, $5, $6)", name, start, end, description, technologies, image)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func deleteBlog(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	fmt.Println(id)

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_projects WHERE id=$1", id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func editForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("views/edit-blog.html")
	if tmpl == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Message : " + err.Error()))
		return
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	editSelectedData := Blog{}

	err = connection.Conn.QueryRow(context.Background(), "SELECT id, name, start_date, end_date, description, technologies, image FROM tb_projects WHERE id=$1", id).Scan(
		&editSelectedData.Id, &editSelectedData.Name, &editSelectedData.Start_date, &editSelectedData.End_date, &editSelectedData.Description, &editSelectedData.Technologies, &editSelectedData.Image)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	editSelectedData.Format_start = editSelectedData.Start_date.Format("2006-01-02")
	editSelectedData.Format_end = editSelectedData.End_date.Format("2006-01-02")
	fmt.Println(editSelectedData.Id, editSelectedData.Name, editSelectedData.Format_start, editSelectedData.Format_end, editSelectedData.Description, editSelectedData, editSelectedData.Technologies, editSelectedData.Image)

	response := map[string]interface{}{
		"editSelected": editSelectedData,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, response)
}

func editBlog(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	} else {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])

		var name = r.PostForm.Get("inputTitle")
		var start = r.PostForm.Get("inputStart")
		var end = r.PostForm.Get("inputEnd")
		//var each.Duration string
		var description = r.PostForm.Get("inputContent")
		var technologies = r.PostForm.Get("js")
		var image = r.PostForm.Get("inputImage")
		fmt.Println(technologies)

		_, err = connection.Conn.Exec(context.Background(), "UPDATE tb_projects SET name=$2, start_date=$3, end_date=$4, description=$5, technologies=$6, image=$7 WHERE id=$1", id, name, start, end, description, technologies, image)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("message : " + err.Error()))
			return
		}

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}

}

func formRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/form-register.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var name = r.PostForm.Get("inputName")
	var email = r.PostForm.Get("inputEmail")
	var password = r.PostForm.Get("inputPass")

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	fmt.Println(passwordHash)
	fmt.Println(name, email, password)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_user(name, email, password) VALUES ($1, $2, $3)", name, email, passwordHash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	http.Redirect(w, r, "/form-login", http.StatusMovedPermanently)
}

func formLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/form-login.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	var store = sessions.NewCookieStore([]byte("SESSION_KEY"))
	session, _ := store.Get(r, "SESSION_KEY")

	fm := session.Flashes("message")

	var flashes []string
	if len(fm) > 0 {
		session.Save(r, w)
		for _, fl := range fm {
			flashes = append(flashes, fl.(string))
		}
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, Data)
}

func login(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("SESSION_KEY"))
	session, _ := store.Get(r, "SESSION_KEY")

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	email := r.PostForm.Get("inputEmail")
	password := r.PostForm.Get("inputPass")

	user := User{}

	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_user WHERE email=$1", email).Scan(&user.Id, &user.Name, &user.Email, &user.Password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("message : " + err.Error()))
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("message : " + err.Error()))
		return
	}
	session.Values["IsLogin"] = true
	session.Values["Name"] = user.Name
	session.Options.MaxAge = 100000 // 3 jam

	session.AddFlash("Successfully Login", "message")
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logout")

	var store = sessions.NewCookieStore([]byte("SESSION_KEY"))
	session, _ := store.Get(r, "SESSION_KEY")
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
