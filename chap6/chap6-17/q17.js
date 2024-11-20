function Main(input) {
    const N = input;
    const minLength = Infinity;
    
    for(let a = 1; a*a <= N; a++){
      if(N%a == 0){
        const b = N/a;
        const length = 2 * ( a + b );
        if(length < minLength){
          minlength = length;
        }
      }
    }
    console.log(minlength)
  }
  
  Main(require("fs").readFileSync("/dev/stdin", "utf8"));