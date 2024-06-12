function Main(input) {
  let inputs = input.trim().split("\n");
  let a = inputs[0].split(" ").map(Number);
  let b = inputs[1].split(" ").map(Number);
  
  let dist = Math.sqrt((a[0]-b[0])**2 + (a[1]-b[1])**2);
  
  if (dist < a[2]-b[2]){
    console.log(1)
  } else if(dist == a[2]-b[2]){
    console.log(2)
  }else if(dist < a[2]+b[2]){
    console.log(3)
  }else if(dist == a[2]+b[2]){
    console.log(4)
  }else{
    console.log(5)
  }
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));