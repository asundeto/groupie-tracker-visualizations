const first = document.getElementById("first")
const info = document.getElementById("info")
const exit = document.getElementById("exit")
const desc = document.querySelector('.description')
var cnt = 0
var id = 0
var check = false

function Toogle(slider_all, n) {
    if (n > 0) {
        desc.classList.add('hidden')
    } else {
        desc.classList.remove('hidden')
    }
    if (n+1 < slider_all.length) {
        slider_all[n].classList.toggle('opened')
    } else {
        slider_all[n].classList.toggle('opened2')
    }
    slider_all[n].scrollIntoView()
    slider_all[n].scrollIntoView({ behavior: 'smooth', block: 'nearest', inline: 'center' });   
    slider__img = slider_all[n].querySelector('.slider__img')
    img_crcle = slider__img.querySelector('.img_url')
    img_crcle.classList.toggle('opened')
}

function Remove(slider_all, n) {
    if (n > 0) {
        desc.classList.add('hidden')
    } else {
        desc.classList.remove('hidden')
    }
    slider_all[n].classList.remove('opened')
    if (n+1 == slider_all.length) {
        slider_all[n].classList.remove('opened2')
    }
    slider_all[n].classList.remove('opened2')
    slider__img = slider_all[n].querySelector('.slider__img')
    img_crcle = slider__img.querySelector('.img_url')
    img_crcle.classList.remove('opened')
}

function Members(array) {
    var res = ""
    for (let i = 0; i < array.length; i++) {
        res += String(array[i])
        if (i + 1 != array.length) {
            res += ", "
        }
    }
    return res
}

refreshCSS = () => {
    console.log("start_reff");
    let links = document.getElementsByTagName('link');
    for (let i = 0; i < links.length; i++) {
        if (links[i].getAttribute('rel') == 'stylesheet') {
            let href = links[i].getAttribute('href')
                .split('?')[0];
            let newHref = href + '?version=' +
                new Date().getMilliseconds();
            links[i].setAttribute('href', newHref);
        }
    }
    console.log("END REF");
}

function getArtists() {
    for (let i = 0; i < artists.length; i++) {
        addElement(i)
    }
}

function addElement(id) {
    var slider_item = document.createElement("div");
    slider_item.classList.add('swiper-slide', 'slider__item')
    var slider_img = document.createElement("div")
    slider_img.classList.add('slider__img')
    var img_url = document.createElement("img")
    img_url.classList.add('img_url')
    img_url.src = artists[id].image
    slider_img.appendChild(img_url);
    slider_item.appendChild(slider_img)
    first.appendChild(slider_item)
}

$(document).ready(function() {
    getArtists();
    var slider_all = document.querySelectorAll('.slider__item')
    for (let i = 0; i < slider_all.length; i++) {
        slider_all[i].addEventListener('click', event => {
            if (check == true && id != i) {
                Remove(slider_all, id)
                cnt--
            }
            Toogle(slider_all, i)
            if (cnt % 2 == 0) {
                check = true
                info.style.display = 'flex'
                exit.style.display = 'flex'
            } else {
                check = false
                info.style.display = 'none'
                exit.style.display = 'none'
            }
            id = i
            cnt++
            info.querySelector('.artist_name').innerText = artists[id].name
            info.querySelector('.artist_date').innerText = "Creation date: " + artists[id].creationDate
            info.querySelector('.artist_members').innerText = Members(artists[id].members)
            info.querySelector('.artist_album').innerText = "First album: " + artists[id].firstAlbum
            info.querySelector('.artist_calendar').innerText = artists[id].DatesLocat
        })
    }

    exit.addEventListener('click', function() {
        Remove(slider_all, id)
        info.style.display = 'none'
        exit.style.display = 'none'
        check = false
        cnt = 0
    })
    
    const sliderMain = new Swiper('.slider_main', {
        freeMode: true,
        centeredSlides: true,
        mousewheel: true,
        speed: 10000,
        parallax: true,
        breakpoints: {
            0: {
                slidesPerView: 2.5,
                spaceBetween: 20
            },
            100: {
                slidesPerView: 4,
                spaceBetween: 60
            }
        }
    })

    const sliderBg = new Swiper('.slider_bg', {
        centeredSlides: true,
        parallax: true,
        spaceBetween: 60,
        slidesPerView: 3.5
    })

    sliderMain.controller.control = sliderBg

    sliderMain.on('slideChange', e => {
        if (sliderMain.activeIndex > 0) {
            desc.classList.add('hidden')
        } else {
            desc.classList.remove('hidden')
            slider_all[0].scrollIntoView()
            slider_all[0].scrollIntoView({ behavior: 'smooth', block: 'nearest', inline: 'center' });
        }
    })
    
    refreshCSS()
});