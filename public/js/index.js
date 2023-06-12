let hamburgerIsOpen = false;

function openHamburger() {
    let hamburgerNav = document.getElementById ("hamburger-nav-container");

    // jika pada hamburgerIsOpen ditambah tanda seru jadi !hamburgerIsOpen maka
    // hamburgerIsOpen == false , sebaliknya jika tanpa tanda seru maka
    // hamburgerIsOpen == true

    if (!hamburgerIsOpen) {
        console.log(hamburgerIsOpen);
        hamburgerNav.style.display = "block";
        hamburgerIsOpen = true;
        // disini dilihat kondisinya karena diatas let hamburgerIsOpen = false maka kondisi yang dihasilkan true karena sesuai
    } else {
        console.log(hamburgerIsOpen);
        hamburgerNav.style.display = "none";
        hamburgerIsOpen = false;
    }
}