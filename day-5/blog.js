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
                        ${dataBlog[index].tittle}
                    </a>
                </h1>
                <div class="detail-blog-content">
                </div>
                <p>
                    ${dataBlog[index].content}
                </p>
                <div>
                <p style="font-size: 15px; color: grey">${distanceTime}</p>
                </div>
            </div>
        </div>
        `
    }
}

function getFullTime(time) {
    //time = new Date ()
    //cosnsole.log(time)

    let monthName = ['Jan','Feb','Mar','Apr','May','Jun','Jul','Agt','Sep','Oct','Nov','Dec']
    //console.log(monthName[9]);

    let date = time.getDate()
    console.log(date);

    let monthIndex = time.getMonth()
    console.log(monthIndex);

    let year = time.getFullYear
    console.log(year);

    let hours = time.getHours()
    let minutes = time.getMinutes()

    if (hours <= 9) {
        hours = "0" + hours
        } else if (minutes <= 9) {
    minutes = "0" + minutes
    
    }

    // 
    return `$(date) $(monthName[monthIndex]) $(year) $(hours):$(minutes)`

}

function getDistance(time) {
    let timeNow = new Date()
    let timePost = time

    let distance = timeNow - timePost //milisecond
    console.log(distance);
    
    let milisecond = 1000 // milisecond
    let secondINHours = 3600 // 1 jam 3600
    let hoursInDay = 24 // 1 hari = 24 jam

    let distanceDay = Math.floor(distance/ (milisecond * secondINHours * hoursInDay))
    let distanceHours = Math.floor(distance / (milisecond * 60 * 60))
    let distanceMinutes = Math.floor(distance / (milisecond * 60))
    let distanceSecond = Math.floor(distance / milisecond)

    if (distanceDay > 0) {
        return `$(distanceDay) day(s) ago`
    }
    else if (distanceHours > 0) {
        return `$(distanceHours) hour(s)`
    } else if (distanceMinutes > 0) {
        return `$(distanceMinutes) minutes(s)`
    } else {
        return `${distanceSecond} second(s) ago`
    }
}

/*setInterval (function() {
    renderBlog()    
}, 1000)*/