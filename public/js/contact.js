// console.log("Hello Student");
// alert("Good Morning madafaka");
// document.write("Batch 47");
// Variable digunakan untuk menampun data-data, ada var, let dan const
// variable bisa dipakai dengan 3 yaitu string atau kata huruf, number dan boolean (true false)
// variable itu bisa di deklarasikan ulang
// var gelas = "Air Putih";
// var gelas = "Kopi";
// console.log(gelas);

// // javascript mengerjakan perintah dari ataas sampe bawah berurutan jadi jika console lognya dibawah air putih maka yang akan muncul var air putih

// let pemerintah = "PDIP";
// pemerintah = "Nasdem"
// console.log(pemerintah)
// // let cuma bisa dipanggil valuenya, gk bisa dipanggil let lagi, ibarat pemerintah itu cuma satu yang berkuasa

// let Pemerintah = "Golkar"
// console.log(Pemerintah)

// // const valuenya konstan yidak bisa dideklaraasikan tidak bisa dirubah sama sekali valuenya
// const rog = "asus"
// // rog = "lenovo"
// console.log(rog)

// tipe data primitif
// let nama = "Rizki"; //valuenya string
// let umur = 24; //number dan tidak pakai tanda kutip kalau pakai tanda kutip malah string
// let isMarried = false; //boolean

// // nama saya Rizki umur saya 24 tahun
// console.log("nama saya Rizki umur saya 23 tahun");
// // ini kepanjangan dan manual
// console.log(`nama saya ${nama} umur saya ${umur}`);
// console.log("nama saya", nama, "umur saya", umur, "tahun");
// console.log("nama saya" + nama + "umur saya" + umur + "tahun"); //spasinya harus manual

// operator
// let x = 48;
// let y = 4;
// kalau 4 diubah jadi string dengan menambahkan kutip" maka resultnya menjadi 484"
// kalau operatornya dibagi sedangkan 4 sudah menjadi string js malah membaginya menjadi 12 sedangkan number tidak bisa dioperasikan dengan string, inilah kelemahan js
// let result = x + y;

// console.log(result);

// condition
// jika nilai sama dengan atau lebih dari 75 maka lulus, jadi ada kondisi disini
// let nilai = 75
// kalau if nilai 75 diubah jadi string tetep keluar kamu lulus
// if (nilai = 75) {
//     console.log("Kamu lulus!");
// } else {
//     console.log("Kamu tidak lulus!")
// }

// function
// function Hello() {
//     let x = 5;
//     let y = 10;

//     let result = x * y;
//     console.log(result)
// }
// // fungsinya harus dipanggil supaya ibarat suatu mobil bisa berfungsi karena dipakai
// Hello();

// function Hello2(x, y) {
//     console.log(x);
//     console.log(y);

//     let result = x * y;
//     console.log(result)
// }
// Hello2(5, 10)

// dalam memberi nama fungsi ada unwriter rules misal kalo kita menggunakan dua kata misal namasaya itu ditulis namaSaya
// camelcase = namaSayaAdalah

// function namaSaya(name) {
//     console.log(name;)
// }

// namaSaya("Abel");

// function submitData() {
//     let name = document.getElementById("input-name").value;
//     let email = document.getElementById("input-email").value;
//     let phone = document.getElementById("input-phone").value;
//     let subject = document.getElementById("input-subject").value;
//     let message = document.getElementById("input-message").value;

//         if (name == "") {
//             return alert("nama harus diisi!")
//         } else if (email == "") {
//             return alert("emailnya woi!")
//         } else if (phone == "") {
//             return alert("masa gk punya nomor")
//         } else if (message == "") {
//             return alert("kata-katanya dulu kanda")
//         }

//         let emailReceiver = "rizkiboediman59@gmail.com"

//         let a = document.createElement("a")
//         a.href = `mailto:${emailReceiver}?subject=${subject}&body=Halo, nama saya ${name}, ${message}, Silahkan kontak saya di nomor ${phone}, terima kasih`
//         a.click()

//     console.log(name);
//     console.log(email);
//     console.log(phone);
//     console.log(subject)
//     console.log(message);

//     let emailer = {
//         name,
//         email,
//         phone,
//         subject,
//         message,
//     }
//     console.log(emailer)
// }

function submitData() {
    let name = document.getElementById("input-name").value;
    let phone = document.getElementById("input-phone").value;
    let subject = document.getElementById("input-subject").value;
    let message = document.getElementById("input-message").value;

    let emailReceiver = "rizkiboediman59@gmail.com"
    let a = document.createElement("a")
    a.href = `mailto:${emailReceiver}?subject=${subject}&body=Halo, nama saya ${name}, ${message}. Silahkan kontak saya di ${phone}, terima kasih`
}



