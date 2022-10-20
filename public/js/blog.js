let dataBlog = []

function addBlog(event) {
    event.preventDefault()

    let title = document.getElementById("input-title").value
    let start = document.getElementById("start").value
    let end = document.getElementById("end").value
    let content = document.getElementById("input-content").value
    let tech = document.getElementById("js").value
    let image = document.getElementById("input-blog-image").files[0]

    
    // buat url gambar nantinya tampil
    image = URL.createObjectURL(image)
    console.log(tech)
    
    let distance = getDistanceTime(start, end);
    
    let blog = {
        title,
        duration: distance,
        content,
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

        const htmlRender = `
        <a href="blog-detail/1">
            <div  class="blog-list-item">
                <div class="blog-image">
                    <img src="${dataBlog[index].image}">
                </div>
                <div class="blog-content">
                    <div class="btn-group">
                        <button class="btn-edit">Edit</button>
                        <button class="btn-post">Delete</button>
                    </div>
                    <div>
                    <h2 style="color:grey;">
                    Durasi: ${getDistanceTime(dataBlog[index].postAt)}
                    </div>
                    </div>
                    <h1>
                        <a href="blog-detail.html" target="_blank">
                            ${dataBlog[index].title}
                        </a>
                    </h1>
                    <div class="detail-blog-content">
                        ${getFullTime(dataBlog[index].postAt)} | ${dataBlog[index].author}
                    </div>
                    <p>
                        ${dataBlog[index].content}
                    </p> 
                </div>
            </div>
        <a/>
        `
        document.getElementById("contents").innerHTML += htmlRender
    }

}


function getFullTime(time) {
    time = new Date()
    console.log(time)

    let monthName = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']
    console.log(monthName[9])

    // 14
    let date = time.getDate()
    console.log(date)

    // 9
    let monthIndex = time.getMonth()
    console.log(monthIndex)

    // 2022
    let year = time.getFullYear()
    console.log(year)

    let hours = time.getHours()
    let minutes = time.getMinutes()

    console.log(hours)

    if (hours <= 9) {
        hours = "0" + hours
    } 
    
    if (minutes <= 9) {
        minutes = "0" + minutes
    }

    // 14 Oct 2022 09:07 WIB
    return `${date} ${monthName[monthIndex]} ${year} ${hours}:${minutes} WIB`
}

function getDistanceTime(time) {
    let timeNow = new Date()
    let timePost = time

    let distance = timeNow - timePost //milisecond
    console.log(distance)

    let milisecond = 1000 // milisecond
    let secondInHours = 3600 // 1 jam = 3600 detik
    let hoursInDay = 24 // 1 hari = 24 jam

    let distanceDay = Math.floor(distance / (milisecond * secondInHours * hoursInDay))
    let distanceHours = Math.floor(distance / (milisecond * 60 * 60))
    let distanceMinutes = Math.floor(distance / (milisecond * 60))
    let distanceSecond = Math.floor(distance / milisecond)

    if (distanceDay > 0) {
        return `${distanceDay} day(s) ago`
    } else if (distanceHours > 0) {
        return `${distanceHours} hour(s) ago`
    } else if (distanceMinutes > 0) {
        return `${distanceMinutes} minute(s) ago`
    } else {
        return `${distanceSecond} second(s) ago`
    }
}


/*setInterval(function() {
    renderBlog()
    }, 3000)
*/
/* setInterval(intervalFunction, 6000)

 function intervalFunction() {
     renderBlog()
 }*/