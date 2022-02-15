const p = document.querySelector('p')
const p2 = document.querySelector('h2')
const divs = document.querySelectorAll('div')
console.log(p)
// console.dir(p2)

const test = document.querySelector('.nice')
console.log(test.classList)

function toggleWhenClick(){
    test.classList.toggle('round');

}

test.addEventListener('click', toggleWhenClick)

test.alt="random"
test.width = 200;

// window.addEventListener('load', function(){
//     console.log(test.naturalWidth)
// })

console.log(test.getAttribute('alt'))