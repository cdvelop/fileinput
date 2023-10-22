
function prevButtonFileImg(e) {
    MoveScrollFileImg("left", e)
}

function nextButtonFileImg(e) {
    MoveScrollFileImg("right", e)
}

function MoveScrollFileImg(to, e) {

    let move_left = -100
    if (to === "right") {
        move_left = 100
    }

    const tagname = e.target.tagName.toLowerCase()
    // console.log("prev",e.target,"tagname:",tagname)
    let container = e.target.parentNode
    if (tagname === "i") {
        container = container.parentNode
    }

    container = container.querySelector('.scroll-file-img')
    // console.log("container",container)
    container.scrollBy({ left: move_left, behavior: 'smooth' }); // Ajusta el valor según lo que desees

}