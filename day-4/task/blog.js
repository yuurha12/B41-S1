let dataBlog = []

function addBlog(event) {
    event.preventDefault()

    let tittle = document.getElementById("input-tittle").value
    let content = document.getElementById("input-content").value
    let image = document.getElementById("input-blog-image").file[0]


// buat url gambar tampil

image = URL.createObjectURL(image[0])
console.log(image)

let blog= {
    tittle,
    content,
    image,
    postAt: new Date(),
    author: "Hoki Wahyu"

    }

dataBlog.push(blog)
console.log(dataBlog)

renderBlog()
}    

function renderBlog() {
    document.getElementById("contents").innerHTML = ''

    for (let index = 0; index < dataBlog.length; index++) {
        console.log(dataBlog[index])
    }

        document.getElementById("contents").innerHTML +=``
        
}
