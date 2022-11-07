package main

import (
	"context"
	"day-final/connection"
	"day-final/middleware"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

// Struct for session data
type MetaData struct {
	Title     string
	UserId    int
	IsLogin   bool
	UserName  string
	Password  string
	FlashData string
}

var Data = MetaData{
	Title: "web",
}

// struct for User data
type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

// struct for Blog
type Blog struct {
	Id           int
	Name         string
	Start_date   time.Time
	End_date     time.Time
	Format_start string
	Format_end   string
	Author       string
	Duration     string
	Description  string
	Technologies []string
	Image        string
	IsLogin      bool
}

var Blogs = []Blog{
	/*{
		Id:           0,
		Name:         "Dumbways mobile app-2021",
		Start_date:   "2022-10-17",
		End_date:     "2022-10-24",
		Duration:     "1 Minggu",
		Description:  "Test",
		Technologies: "Node Js",
	},
	{
		Id:           1,
		Name:         "Dumbways mobile app-2021",
		Start_date:   "2022-10-17",
		End_date:     "2022-10-24",
		Duration:     "1 Minggu",
		Description:  "Test",
		Technologies: "Node Js",
	},*/
}

// route function for all page
func main() {
	route := mux.NewRouter()

	connection.DatabaseConnect()

	// route path folder for public
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	//route for upload
	route.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	//routing page and post data
	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/blog", blog).Methods("GET")
	route.HandleFunc("/blog-detail/{id}", blogDetail).Methods("GET")
	route.HandleFunc("/form-blog", formAddBlog).Methods("GET")
	route.HandleFunc("/add-blog", middleware.UploadFile(addBlog)).Methods("POST")
	route.HandleFunc("/delete-blog/{id}", deleteBlog).Methods("GET")
	route.HandleFunc("/edit-form-blog/{id}", editForm).Methods("GET")
	route.HandleFunc("/edit-blog/{id}", middleware.UploadFile(editBlog)).Methods("POST")

	route.HandleFunc("/form-register", formRegister).Methods("GET")
	route.HandleFunc("/register", register).Methods("POST")
	route.HandleFunc("/form-login", formLogin).Methods("GET")
	route.HandleFunc("/login", login).Methods("POST")
	route.HandleFunc("/logout", logout).Methods("GET")

	fmt.Println("Server running on port 5000")
	http.ListenAndServe("localhost:5000", route)
}

// home page
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, err := template.ParseFiles("views/index.html")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return

	} else {
		Data = MetaData{}

		var store = sessions.NewCookieStore([]byte("SESSION_KEY"))
		session, _ := store.Get(r, "SESSION_KEY")

		if session.Values["IsLogin"] != true {
			Data.IsLogin = false
		} else {
			fm := session.Flashes("message")
			var flashes []string
			if len(fm) > 0 {
				session.Save(r, w)
				for _, fl := range fm {
					flashes = append(flashes, fl.(string))
				}
			}

			Data.IsLogin = session.Values["IsLogin"].(bool)
			Data.UserId = session.Values["Id"].(int)
			Data.UserName = session.Values["Name"].(string)
			Data.FlashData = strings.Join(flashes, "")
		}

		if !Data.IsLogin {
			var blogData []Blog
			dataDb, _ := connection.Conn.Query(context.Background(), "SELECT id, name, start_date, end_date, duration, description, technologies, image FROM tb_projects")
			for dataDb.Next() {
				each := Blog{}
				err := dataDb.Scan(&each.Id, &each.Name, &each.Start_date, &each.End_date, &each.Duration, &each.Description, &each.Technologies, &each.Image)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				blogData = append(blogData, each)
			}

			response := map[string]interface{}{
				"Data":     Data,
				"blogData": blogData,
			}

			w.WriteHeader(http.StatusOK)
			tmpl.Execute(w, response)

		} else {
			var blogData []Blog
			dataDb, _ := connection.Conn.Query(context.Background(), "SELECT tb_projects.id, tb_projects.name, start_date, end_date, duration, description, technologies, image FROM tb_projects LEFT JOIN tb_user ON tb_projects.author_id = tb_user.id WHERE tb_user.id = $1 ORDER BY id DESC", Data.UserId)
			for dataDb.Next() {
				each := Blog{}
				err := dataDb.Scan(&each.Id, &each.Name, &each.Start_date, &each.End_date, &each.Duration, &each.Description, &each.Technologies, &each.Image)
				if err != nil {
					fmt.Println(err.Error())
					return
				}

				each.IsLogin = Data.IsLogin

				blogData = append(blogData, each)
			}

			response := map[string]interface{}{
				"Data":     Data,
				"blogData": blogData,
			}

			w.WriteHeader(http.StatusOK)
			tmpl.Execute(w, response)
		}
	}
}

// contact page
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

// blog page
func blog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

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
		Data.UserId = session.Values["Id"].(int)
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

		each.Format_start = each.Start_date.Format("02-01-2006")
		each.Format_end = each.End_date.Format("02-01-2006")

		layoutDate := "2006-01-02"
		startParse, _ := time.Parse(layoutDate, each.Start_date.Format("2006-01-02"))
		endParse, _ := time.Parse(layoutDate, each.End_date.Format("2006-01-02"))

		hour := 1
		day := hour * 24
		week := hour * 24 * 7
		month := hour * 24 * 30
		year := hour * 24 * 365

		differHour := endParse.Sub(startParse).Hours()
		var differHours int = int(differHour)

		days := differHours / day
		weeks := differHours / week
		months := differHours / month
		years := differHours / year

		if differHours < week {
			each.Duration = strconv.Itoa(int(days)) + " Hari"
		} else if differHours < month {
			each.Duration = strconv.Itoa(int(weeks)) + " Minggu"
		} else if differHours < year {
			each.Duration = strconv.Itoa(int(months)) + " Bulan"
		} else if differHours > year {
			each.Duration = strconv.Itoa(int(years)) + " Tahun"
		}

		if session.Values["IsLogin"] != true {
			each.IsLogin = false
		} else {
			each.IsLogin = session.Values["IsLogin"].(bool)
		}

		result = append(result, each)
	}

	respData := map[string]interface{}{
		"Data":  Data,
		"Blogs": result,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, respData)
}

// blog detail page
func blogDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//for id project/blogs
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var tmpl, err = template.ParseFiles("views/blog-detail.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	//array struct blog
	var BlogDetail = Blog{}

	var store = sessions.NewCookieStore([]byte("SESSION_KEY"))
	session, _ := store.Get(r, "SESSION_KEY")

	if session.Values["IsLogin"] != true {
		Data.IsLogin = false
	} else {
		Data.IsLogin = session.Values["IsLogin"].(bool)
		Data.UserName = session.Values["Name"].(string)
		Data.UserId = session.Values["Id"].(int)
	}

	err = connection.Conn.QueryRow(context.Background(), "SELECT id, name, start_date, end_date, duration, description, technologies, image, author_id FROM tb_projects WHERE id=$1", id).Scan(
		&BlogDetail.Id, &BlogDetail.Name, &BlogDetail.Start_date, &BlogDetail.End_date, &BlogDetail.Duration, &BlogDetail.Description, &BlogDetail.Technologies, &BlogDetail.Image, &BlogDetail.Author)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message: " + err.Error()))
		return
	}
	BlogDetail.Format_start = BlogDetail.Start_date.Format("2 January 2006")
	BlogDetail.Format_end = BlogDetail.End_date.Format("2 January 2006")

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

// page form addblog
func formAddBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/add-blog.html")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Message : " + err.Error()))
		return
	} else {
		var store = sessions.NewCookieStore([]byte("SESSION_KEY"))
		session, _ := store.Get(r, "SESSION_KEY")

		if session.Values["IsLogin"] != true {
			Data.IsLogin = false
		} else {
			Data.IsLogin = session.Values["IsLogin"].(bool)
			Data.UserId = session.Values["Id"].(int)
			Data.UserName = session.Values["Name"].(string)
		}
		response := map[string]interface{}{
			"Data": Data,
		}

		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, response)
	}
}

// send addblog data
func addBlog(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	//input data from form input
	var name = r.PostForm.Get("inputTitle")
	var start = r.PostForm.Get("inputStart")
	var end = r.PostForm.Get("inputEnd")
	var duration string
	var description = r.PostForm.Get("inputContent")
	var technologies = []string{r.PostForm.Get("node"), r.PostForm.Get("react"), r.PostForm.Get("nextjs"), r.PostForm.Get("typescript")}

	dataContex := r.Context().Value("dataFile")
	image := dataContex.(string)

	layoutDate := "2006-01-02"
	startParse, _ := time.Parse(layoutDate, start)
	endParse, _ := time.Parse(layoutDate, end)

	hour := 1
	day := hour * 24
	week := hour * 24 * 7
	month := hour * 24 * 30
	year := hour * 24 * 365

	differHour := endParse.Sub(startParse).Hours()
	var differHours int = int(differHour)

	days := differHours / day
	weeks := differHours / week
	months := differHours / month
	years := differHours / year

	if differHours < week {
		duration = strconv.Itoa(int(days)) + " Hari"
	} else if differHours < month {
		duration = strconv.Itoa(int(weeks)) + " Minggu"
	} else if differHours < year {
		duration = strconv.Itoa(int(months)) + " Bulan"
	} else if differHours > year {
		duration = strconv.Itoa(int(years)) + " Tahun"
	}

	//query call
	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_projects(name, start_date, end_date, duration, description, technologies, image, author_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", name, start, end, duration, description, technologies, image, Data.UserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	//redirect after submit to home
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func deleteBlog(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	//command to delete post from query/db pgadmin4
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
	} else {
		var store = sessions.NewCookieStore([]byte("SESSION_KEY"))
		session, _ := store.Get(r, "SESSION_KEY")

		if session.Values["IsLogin"] != true {
			Data.IsLogin = false
		} else {
			Data.IsLogin = session.Values["IsLogin"].(bool)
			Data.UserName = session.Values["Name"].(string)
		}
		//id which post want to delete
		id, _ := strconv.Atoi(mux.Vars(r)["id"])

		editSelectedData := Blog{}

		err = connection.Conn.QueryRow(context.Background(), "SELECT id, name, start_date, end_date, duration, description, technologies, image FROM tb_projects WHERE id=$1", id).Scan(
			&editSelectedData.Id, &editSelectedData.Name, &editSelectedData.Start_date, &editSelectedData.End_date, &editSelectedData.Duration, &editSelectedData.Description, &editSelectedData.Technologies, &editSelectedData.Image)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("message : " + err.Error()))
			return
		}

		editSelectedData.Format_start = editSelectedData.Start_date.Format("2006-01-02")
		editSelectedData.Format_end = editSelectedData.End_date.Format("2006-01-02")

		response := map[string]interface{}{
			"editSelected": editSelectedData,
			"Data":         Data,
		}

		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, response)
	}
}

func editBlog(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	} else {
		//which post want to edit
		id, _ := strconv.Atoi(mux.Vars(r)["id"])

		//to show what content are gonna be update
		var name = r.PostForm.Get("inputTitle")
		var start = r.PostForm.Get("inputStart")
		var end = r.PostForm.Get("inputEnd")
		var duration string
		var description = r.PostForm.Get("inputContent")
		var technologies = []string{r.PostForm.Get("node"), r.PostForm.Get("react"), r.PostForm.Get("nextjs"), r.PostForm.Get("typescript")}

		layoutDate := "2006-01-02"
		startParse, _ := time.Parse(layoutDate, start)
		endParse, _ := time.Parse(layoutDate, end)

		hour := 1
		day := hour * 24
		week := hour * 24 * 7
		month := hour * 24 * 30
		year := hour * 24 * 365

		differHour := endParse.Sub(startParse).Hours()
		var differHours int = int(differHour)

		days := differHours / day
		weeks := differHours / week
		months := differHours / month
		years := differHours / year

		if differHours < week {
			duration = strconv.Itoa(int(days)) + " Hari"
		} else if differHours < month {
			duration = strconv.Itoa(int(weeks)) + " Minggu"
		} else if differHours < year {
			duration = strconv.Itoa(int(months)) + " Bulan"
		} else if differHours > year {
			duration = strconv.Itoa(int(years)) + " Tahun"
		}

		dataContex := r.Context().Value("dataFile")
		image := dataContex.(string)

		_, err = connection.Conn.Exec(context.Background(), "UPDATE tb_projects SET name=$2, start_date=$3, end_date=$4, duration=$5, description=$6, technologies=$7, image=$8 WHERE id=$1", id, name, start, end, duration, description, technologies, image)
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

	//registration input form in form-register
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
	session.Values["Id"] = user.Id
	session.Options.MaxAge = 100000 // 3 jam

	//alert/notification login success
	session.AddFlash("Login Success", "message")
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func logout(w http.ResponseWriter, r *http.Request) {

	var store = sessions.NewCookieStore([]byte("SESSION_KEY"))
	session, _ := store.Get(r, "SESSION_KEY")
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
