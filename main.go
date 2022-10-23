package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var Data = map[string]interface{}{
	"Title": "Personal Web",
}

type Blog struct {
	Id           int
	Name         string
	Start_date   string
	End_date     string
	Duration     string
	Description  string
	Technologies string
	Image        string
}

var Blogs = []Blog{
	{
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
	},
}

func main() {
	route := mux.NewRouter()

	// route path folder untuk public
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	//routing
	route.HandleFunc("/hello", helloWorld).Methods("GET")
	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/blog", blog).Methods("GET")
	route.HandleFunc("/blog-detail/{index}", blogDetail).Methods("GET")
	route.HandleFunc("/form-blog", formAddBlog).Methods("GET")
	route.HandleFunc("/add-blog", addBlog).Methods("POST")
	route.HandleFunc("/delete-blog/{index}", deleteBlog).Methods("GET")
	route.HandleFunc("/edit-form-blog/{index}", editForm).Methods("GET")
	route.HandleFunc("/edit-blog/{index}", editBlog).Methods("GET")

	fmt.Println("Server running on port 5000")
	http.ListenAndServe("localhost:5000", route)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	respData := map[string]interface{}{
		"Blogs": Blogs,
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

	respData := map[string]interface{}{
		"Blogs": Blogs,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, respData)
}

func blogDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/blog-detail.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	var BlogDetail = Blog{}

	index, _ := strconv.Atoi(mux.Vars(r)["index"])

	for i, data := range Blogs {
		if index == i {
			BlogDetail = Blog{
				Name:         data.Name,
				Description:  data.Description,
				Start_date:   data.Start_date,
				End_date:     data.End_date,
				Duration:     data.Duration,
				Technologies: data.Technologies,
				Image:        data.Image,
			}
		}
	}

	data := map[string]interface{}{
		"Blog": BlogDetail,
	}

	// fmt.Println(data)

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
	var duration string
	var description = r.PostForm.Get("inputContent")
	var technologies = r.PostForm.Get("js")
	var image = r.PostForm.Get("inputImage")

	fmt.Println("Name : " + r.PostForm.Get("inputTitle")) // value berdasarkan dari tag input name
	fmt.Println("Start : " + r.PostForm.Get("inputStart"))
	fmt.Println("End : " + r.PostForm.Get("inputEnd"))
	fmt.Println("Description : " + r.PostForm.Get("inputContent"))
	fmt.Println("Technologies : " + r.PostForm.Get("js"))
	fmt.Println("Image : " + r.PostForm.Get("inputImage"))

	layoutDate := "2006-01-02"
	startParse, _ := time.Parse(layoutDate, start)
	endParse, _ := time.Parse(layoutDate, end)

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
		duration = strconv.Itoa(int(days)) + " Days"
	} else if differHours < month {
		duration = strconv.Itoa(int(weeks)) + " Weeks"
	} else if differHours < year {
		duration = strconv.Itoa(int(months)) + " Months"
	} else if differHours > year {
		duration = strconv.Itoa(int(years)) + " Years"
	}

	newBlog := Blog{
		Name:         name,
		Start_date:   start,
		End_date:     end,
		Duration:     duration,
		Description:  description,
		Technologies: technologies,
		Image:        image,
	}

	//Blogs.push(newBlog)
	Blogs = append(Blogs, newBlog)

	fmt.Println(Blogs)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func deleteBlog(w http.ResponseWriter, r *http.Request) {
	index, _ := strconv.Atoi(mux.Vars(r)["index"])
	fmt.Println(index)

	Blogs = append(Blogs[:index], Blogs[index+1:]...)
	fmt.Println(Blogs)

	http.Redirect(w, r, "/", http.StatusFound)
}

func editForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("views/edit-blog.html")

	if tmpl == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Message : " + err.Error()))
	} else {
		index, _ := strconv.Atoi(mux.Vars(r)["index"])

		BlogDetail := Blog{}

		for id, data := range Blogs {
			if id == index {
				BlogDetail = Blog{
					Id:          id,
					Name:        data.Name,
					Start_date:  data.Start_date,
					End_date:    data.End_date,
					Description: data.Description,
				}
				fmt.Println(BlogDetail.Description)
			}
		}

		response := map[string]interface{}{
			"BlogDetail": BlogDetail,
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
		index, _ := strconv.Atoi(mux.Vars(r)["index"])

		var name = r.PostForm.Get("inputTitle")
		var start = r.PostForm.Get("inputStart")
		var end = r.PostForm.Get("inputEnd")
		var duration string
		var description = r.PostForm.Get("inputContent")
		var technologies = r.PostForm.Get("js")
		var image = r.PostForm.Get("inputImage")

		fmt.Println("Name : " + r.PostForm.Get("inputTitle")) // value berdasarkan dari tag input name
		fmt.Println("Start : " + r.PostForm.Get("inputStart"))
		fmt.Println("End : " + r.PostForm.Get("inputEnd"))
		fmt.Println("Description : " + r.PostForm.Get("inputContent"))
		fmt.Println("Technologies : " + r.PostForm.Get("js"))
		fmt.Println("Image : " + r.PostForm.Get("inputImage"))

		layoutDate := "2006-01-02"
		startParse, _ := time.Parse(layoutDate, start)
		endParse, _ := time.Parse(layoutDate, end)

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
			duration = strconv.Itoa(int(days)) + " Days"
		} else if differHours < month {
			duration = strconv.Itoa(int(weeks)) + " Weeks"
		} else if differHours < year {
			duration = strconv.Itoa(int(months)) + " Months"
		} else if differHours > year {
			duration = strconv.Itoa(int(years)) + " Years"
		}

		editBlog := Blog{
			Name:         name,
			Start_date:   start,
			End_date:     end,
			Duration:     duration,
			Description:  description,
			Technologies: technologies,
			Image:        image,
		}

		Blogs[index] = editBlog

		fmt.Println(Blogs)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}
