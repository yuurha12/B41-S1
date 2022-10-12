//console.log("Hello World");

//variable





//var bisa di panggil ulang berkali kali
var name = "Hoki Wahyuono Pranto"

//let tidak dapat deklarasi ulang untuk menggantikan tinggal ketik let lalu value
let food = "seblak"
food     = "martabak"

//const constant tidak bisa di deklarasikan ulang dan valuenya tidak bisa diubah
const drink = "es teh manis"
//drink = es teh manis

// Data Type

//string huruf kalimat
//number 100/100.00 angka
//boolean true/false no quotes

var school = "SMP NEGERI 17 JAKARTA" //string

var age = 26 //number no quotes

var isStudying = true // no quotes
//camelcase penggunaan dua kata awal huruf kecil lalu kata berikutnya diawali huruf besar

var veggies = "kangkung"

const foodGradeQuality =  "A"

//"" '' 
//$ operator



//console.log("Hi My Name is Kiki", name)
//console.log("Hi my name is", name,"I like cat", veggies)
//console.log(`Hi my name is ${name}, I like cat ${veggies}`)

let x = 5.5
let y = 4.6

let result = y + x


//console.log(result);

// condition



//if - else if else

let score = 75

/*if (score <= 60) {
    console.log("Maaf, Kamu tidak lulus")
} else if (score < 70) {console.log("Maaf kamu tidak lulus, Namun berkesempatan mengulang")
} else {
    console.log("Selamat, kamu lulus!")
}*/

// function
 
function Aritmatika() {
    let bilanganPertama = 50
    let bilanganKedua = 70

    let result = bilanganPertama + bilanganKedua

    console.log(result)
    
}

Aritmatika()

function Aritmatika2(bilanganPertama, bilanganKedua) {
    let result = bilanganPertama + bilanganKedua
    console.log(result)
}

Aritmatika2(40, 80)

function MyName(name) {
    console.log(name)    
}

MyName("Hoki Wahyu")