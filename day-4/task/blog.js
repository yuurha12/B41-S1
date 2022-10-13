let dataBlog = []

function addBlog(event) {
    event.preventDefault()

    let tittle = document.getElementById("tittle").value
    let start = document.getElementById("start").value
    let end = document.getElementById("end").value
    let description = document.getElementById("input-content").value
    let tech = document.getElementById("js").checked = true;
    let image = document.getElementById("input-image").files[0]

    // buat url gambar nantinya tampil
    image = URL.createObjectURL(image)
    console.log(image)

    let blog = {
        tittle,
        start,
        end,
        description,
        tech,
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
        console.log("test",dataBlog[index])

        document.getElementById("contents").innerHTML += `
        <div class="blog-list-item">
            <div class="blog-image">
                <img src="${dataBlog[index].image}">
            </div>
            <div class="blog-content">
            <div class="btn-group">
            <button class="btn-edit">Edit</button>
            <button class="btn-post">Delete</button>
                </div>
                <h1>
                    <a href="blog-detail.html" target="_blank">
                        ${dataBlog[index].title}
                    </a>
                </h1>
                <div class="detail-blog-content">
                </div>
                <p>
                    ${dataBlog[index].content}
                </p>
            </div>
        </div>
        `
    }
}