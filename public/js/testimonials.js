// step 1
// - making class
// - making object
// - called it so it appears in browser using loop and innerHTML
// - refactor the code so it implements inheritance AuthorTestimonials
// - refactor the code so it implements inheritance polymorphism by making class CompanyTestimonial and override get author(the key point is polymorphism implement and override)
// - refactor the code so it implement abstraction by making get testimonialHTML and get author throw error when the chilren didn't use it
// refactor the code so it implements encapsulation, use setter and getter (key point is encapsulation try to make app mor secure by implementing restriction)


// ini super class
// class Testimonial {
//     #quote = "";
//     #image = "";

//     constructor (quote, image) {
//         this.#quote = quote;
//         this.#image = image;
//     }

//     get quote() {
//         return this.quote;
//     }

//     get image() {
//         return this.image;
//     }

//     get author() {
//         throw new Error("getAuthor() method must be implemented");
        // ini adalah cnstructor error
    // }
// this is a polymorphic method that can take any subclasses of Testimonial
    // get testiHTML() {
    //     return `<div class="testi-card">
    //     <img src="${this.#image}"
    //     <p class="quote-testi">${this.#quote}</p>
    //     <p class="author-testi">${this.author}</p>
    // </div>`
    // };
    // fungsi publik dan private juga berhubungan dengan inheritance, sehingga author diatas tidak dipakai tagar
// }

// ini subclass
// class AuthorTestimonials extends Testimonial {
//     #author = ""

//     constructor(author, quote, image) {
//         super(quote, image);
//         this.#author = author;
//     };
// this is an abstract method that subclasses will implement
    // get author() {
    //     return this.#author
    // }
    // kalau get author gk dipanggil atau dipake disini maka yang dikembalikan error. diatas cuma bikin template atau default value

    // cara seperti ini membuat class menjadi global atau terbuka sehingga mudah bagi orang lain untuk mengeditnya, artinya sudah menyentuh prvasi
    // cara menutupnya kasih tanda hash atau pagar diatas, shingga properties quote dan image menjadi private
    // cari di stackoverflow "hash" what does this symbol mean in Javascript
// }
// ini subclass
// class CompanyTestimonials extends Testimonial {
//     #company = ""

//     constructor(company, quote, image) {
//         super(quote, image);
//         this.#company = company;
//     }
//     get author() {
//         return this.#company + " Company"
//     }
// }

// const testimonial1 = new AuthorTestimonials("-Thomas Shelby", "Mantep Banget weh!", "https://assets.kompasiana.com/items/album/2023/02/03/1675436655926-63dd22708c39af24284b5532.png?t=o&v=400");
// const testimonial2 = new AuthorTestimonials("-Johnny Depp", "Keren Banget!", "https://img.thedailybeast.com/image/upload/c_crop,d_placeholder_euli9k,h_2250,w_4000,x_0,y_0/dpr_1.5/c_limit,w_1044/fl_lossy,q_auto/v1672176770/GettyImages-1240637933_nuomc8");
// const testimonial3 = new AuthorTestimonials("-Mamang Mikir", "Jenius Banget!", "https://pyxis.nymag.com/v1/imgs/d6a/dc7/4a5001b7beea096457f480c8808572428b-09-roll-safe.2x.h473.w710.jpg");

// let testimonialData = [testimonial1, testimonial2, testimonial3]
// let testiHTML = ""

// for (let i = 0; i < testimonialData.length; i++) {
//     testiHTML += testimonialData[i].testiHTML;
// }

// document.getElementById("testi").innerHTML = testiHTML;

// console.log(testimonial1.quote)
// console.log(testimonial1.quote + testimonial1.image)
// console.log(testimonial1.author) ini untuk bikin cache error fungsinya biar ketahuan errornya dimana sama kayak feedback error

// get author diatas menandakan kita sudah menggunakan polymorphism


// const testimonialData = [
// {
//     author: "Thomas Shelby",
//     quote: "Mantep banget weh!",
//     image:
//         "https://assets.kompasiana.com/items/album/2023/02/03/1675436655926-63dd22708c39af24284b5532.png?t=o&v=400",
//     rating: 5,
// },
// {
//     author: "Johnny Depp",
//     quote: "Keren banget!",
//     image:
//         "https://img.thedailybeast.com/image/upload/c_crop,d_placeholder_euli9k,h_2250,w_4000,x_0,y_0/dpr_1.5/c_limit,w_1044/fl_lossy,q_auto/v1672176770/GettyImages-1240637933_nuomc8",
//     rating: 4,
// },
// {
//     author: "Mamang Mikir",
//     quote: "Jenius banget!",
//     image:
//         "https://pyxis.nymag.com/v1/imgs/d6a/dc7/4a5001b7beea096457f480c8808572428b-09-roll-safe.2x.h473.w710.jpg",
//     rating: 4,
// },
// {
//     author: "Amin Richman",
//     quote: "Website Murahan! Gua beli 1000 hp buat bukain nih web, kasian lu!",
//     image:
//         "https://cdn.kibrispdr.org/data/4/amin-richman-meme-50.png",
//     rating: 3,
// },
// {
//     author: "Mad Dog",
//     quote: "Website sampah! Gua bengkokin tulang ekorlu ya!",
//     image:
//         "https://i.kym-cdn.com/entries/icons/original/000/016/068/mad-dog.jpg",
//     rating: 1,
// },
// ]

// function allTestimonials() {
//     let testiHTML = "";
//     testimonialData.forEach(function(item) {
//         testiHTML += `<div class="testi-card">
//         <img src="${item.image}"
//         <p class="quote-testi">${item.quote}</p>
//         <p class="author-testi">${item.author}</p>
//         <p class="author-testi">${item.rating} <i class="fa-solid fa-star"></i></p>
//         </div>`
//     })
//     document.getElementById("testi").innerHTML = testiHTML;
// }

// allTestimonials()

// function filterTestimonials(rating) {
//     let testiHTML = "";

//     const testimonialFiltered = testimonialData.filter(function (item) {
//         return item.rating === rating;
//     })

//     // console.log(testimonialFiltered)
//     if (testimonialFiltered.length === 0) {
//         testiHTML += `<h1>Data not found!</h1>`
//     } else {
//         testimonialFiltered.forEach(function(item) {
//             testiHTML += `<div class="testi-card">
//             <img src="${item.image}"
//             <p class="quote-testi">${item.quote}</p>
//             <p class="author-testi">${item.author}</p>
//             <p class="author-testi">${item.rating} <i class="fa-solid fa-star"></i></p>
//             </div>`
//         })
//     }

//     document.getElementById("testi").innerHTML = testiHTML;
// }


const promise = new Promise((resolve, reject) => {
    const testimonialWeb = new XMLHttpRequest();
    testimonialWeb.open('GET', "https://api.npoint.io/7ec9b4e57de9e5dda66f", true);
    console.log(testimonialWeb);
    testimonialWeb.onload = () => {
        if (testimonialWeb.status === 200) {
            // we parsing it so it is easier to read in console
            // response vs responseText, the difference are, responseText is an older version, when response is more newer, but the output is still same/similiar.
            // 200 ini maksudnya yaitu status code di network
            resolve(JSON.parse(testimonialWeb.response));
        } else {
            reject("Data not found.");
        }
    }
    testimonialWeb.onerror = () => {
        reject("Network error.")
    }
    testimonialWeb.send();
})

promise
    .then(function(value) {
    console.log(value)
}).catch(function(reason) {
    return alert("Network error, can't load the testimonials.")
})

async function getAllTestimonials() {
    const response = await promise;
    // console.log(response)

    let testimonialHTML = "";
    response.forEach(function(item) {
    testimonialHTML += `<div class="card p-3 bg-body-tertiary w-lg-25">
                            <img src="${item.image}" class="img-fluid object-fit-cover w-100" />
                            <p class="quote-testi fs-6 fst-italic mt-1">${item.quote}</p>
                            <p class="author-testi fw-bold text-end">${item.rating} <i class="fa-solid fa-star"></i></p>
                        </div>`;
    });

    document.getElementById("testi").innerHTML = testimonialHTML;
}

getAllTestimonials();

async function getFilteredTestimonials(rating) {
    const response = await promise; //dipakai disini artinya kita menunggu promisenya selesai, kalau tidak pake await itu akan ditampilkan status pending

    const testimonialFiltered = response.filter((item) => {
        return item.rating === rating;
    })
    // console.log(testimonialFiltered)

    let testimonialHTML = "";

    if (testimonialFiltered.length === 0) {
        testimonialHTML = "<h1>Data not found!</h1>";
    } else {
        testimonialFiltered.forEach((item) => {
            testimonialHTML += `<div class="card p-3 bg-body-tertiary w-lg-25">
                                    <img src="${item.image}" class="img-fluid object-fit-cover w-100" />
                                    <p class="quote-testi fs-6 fst-italic mt-1">${item.quote}</p>
                                    <p class="author-testi fw-bold text-end">${item.rating} <i class="fa-solid fa-star"></i></p>
                                </div>`;
        });
    };
    document.getElementById("testi").innerHTML = testimonialHTML;
}




