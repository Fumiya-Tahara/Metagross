//2
let result = 2 / Math.E;
console.log(result);

//3
function newton(){
  let r = 2.0;
  let a = 1.0;

  for(let i = 1; i <= 10; i++){
    let x = a;
    let y = Math.exp(a);

    let sA = y;
    let sB = y - sA * x;

    let newA = (r - sB) / sA;
    console.log(i + ":"+ newA);
    a = newA;
  }
  return 0;
}
newton();