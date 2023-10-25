
function moveScrollFileImg(t) {
    // console.log("TARGET:",t)
    let move_left = -100
    if (t.name === "next_img") {
        move_left = 100
    }

    // const tagname = t.tagName.toLowerCase()
    // console.log("tagname:",tagname)
    let container = t.parentNode
    // if (tagname === "i") {
    //     container = container.parentNode
    // }

    container = container.querySelector('.scroll-file_img')
    // console.log("container",container)
    container.scrollBy({ left: move_left, behavior: 'smooth' }); // Ajusta el valor según lo que desees

}