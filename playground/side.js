// console.log('hey')

// custom function with return value
function calculateBill() {
    console.log('Running the function')
    const total = 100 * 1.13;
    return total
}

//storing the function value in myTotal variable
const myTotal = calculateBill()

//printing the value
console.log(myTotal)


//closures

function createGreeting(greeting = '') {
    const myGreet = greeting.toUpperCase();
    return function (name){

        return `${myGreet} ${name}`;
    }
}

const sayHello = createGreeting('Hello')