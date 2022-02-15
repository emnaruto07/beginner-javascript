const myParagraph = document.createElement('p')
myParagraph.textContent = "I'm a P";
myParagraph.classList.add('special')
console.log(myParagraph)

const myDiv = document.createElement('div');
myDiv.classList.add('wrapper');
console.log(myDiv)

myDiv.appendChild(myParagraph)
document.body.appendChild(myDiv)

const ul = document.createElement('ul')
const li = document.createElement('li')
li.textContent = "three"
ul.appendChild(li)

// myDiv.appendChild(ui)
// myDiv.appendChild(li)
document.body.insertAdjacentElement("afterbegin", ul)