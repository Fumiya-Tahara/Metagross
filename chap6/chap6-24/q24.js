function main(input) {
  input = input.trim().split('\n');
  let n = BigInt(input[0]);
  let x = [], y = [];
  for (let i = 0n; i < n; i++) {
    [x[i],y[i]] = input[i+1n].split(' ').map(val => BigInt(val));
  }
  let [a,b] = input[n+1n].split(' ').map(val => BigInt(val));

  let count = 0;
  for (let i = 0n; i < n; i++) {
    let xa = x[i] - a;
    let ya = y[i] - b;
    let xb = x[(i + 1n) % n] - a;
    let yb = y[(i + 1n) % n] - b;
    if (ya > yb) {
      [xa, xb] = [xb, xa];
      [ya, yb] = [yb, ya];
    }
    if (ya<=0 && 0<yb && xa*yb-xb*ya<0) {
      count += 1;
    }
  }
  if (count%2 == 1) {
    console.log('INSIDE');
  } else {
    console.log('OUTSIDE')
  }

  
}
main(require("fs").readFileSync("/dev/stdin", "utf8"));