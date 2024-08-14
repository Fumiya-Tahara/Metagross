function modpow(a, b, m) {
  let p = a, ans = 1n;
  for (let i = 0n; i < 30; i++) {
    if ((b & (1n << i)) != 0) {
      ans *= p;
      ans %= m;
    }
    p *= p;
    p %= m;
  }
  return ans;
}
function division(a, b, m) {
  return (a * modpow(b, m-2n, m)) % m;
}

function main(input) {
  input = input.trim().split('\n');
  let [x,y] = input[0].split(' ').map(val => BigInt(val));

  let mod = 1000000007n;
  if (2n*y-x<0 || 2n*x-y<0) {
    return console.log(0); 
  }
  if ((2n*y-x)%3n!=0 || (2n*x-y)%3n!=0) {
    return console.log(0); 
  }
  let bunshi = 1n;
  let bunbo = 1n;
  let a = (2n*y-x) / 3n;
  let b = (2n*x-y) / 3n;
  for (let i = 1n; i <= a+b; i++) {
    bunshi *= i;
    bunshi %= mod;
  }
  for (let i = 1n; i <= a; i++) {
    bunbo *=i;
    bunbo %= mod;
  }
  for (let i = 1n; i <= b; i++) {
    bunbo *=i;
    bunbo %= mod;
  }
  console.log(division(bunshi, bunbo, mod).toString());
  
  
}
main(require("fs").readFileSync("/dev/stdin", "utf8"));