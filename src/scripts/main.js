var gasser = document.querySelector(".name")
const og_m_left = 0.5
const m_left_trigger = 0.5
const og_m_top = 0.01
const m_top_trigger = 0.25
//setInterval(hover_x, 720)
setInterval(hover_y, 850)
const index = {
    x: false,
    y: false
}
function hover_x() {
    const rand = Math.random() / 2
    const style = window.getComputedStyle(gasser)
    if (index.x) {
        gasser.style.marginLeft = (og_m_left + rand) + "vw";
        index.x = !index.x
    }
    else {
        gasser.style.marginLeft = (m_left_trigger + rand) + "vw";
        index.x = !index.x
    }
}

function hover_y() {
    const rand = Math.random() / 2
    const style = window.getComputedStyle(gasser)
    if (index.y) {
        gasser.style.marginTop = (og_m_top) + "vw";
        index.y = !index.y
    }
    else {
        gasser.style.marginTop = (m_top_trigger) + "vw";
        index.y = !index.y
    }
}