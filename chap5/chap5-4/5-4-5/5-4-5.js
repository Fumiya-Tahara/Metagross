function gcd(a, b) {
  while (b !== 0) {
    const temp = b;
    b = a % b;
    a = temp;
  }
  return a;
}

function lcm(a, b) {
  return (a * b) / gcd(a, b);
}

function countMultiples(N, K, V) {
  let count = 0;

  for (let i = 1; i < 1 << K; i++) {
    let bits = i;
    let lcmValue = 1;
    let bitsCount = 0;

    for (let j = 0; j < K; j++) {
      if (bits & (1 << j)) {
        lcmValue = lcm(lcmValue, V[j]);
        bitsCount++;
      }
    }

    if (bitsCount % 2 === 1) {
      count += Math.floor(N / lcmValue);
    } else {
      count -= Math.floor(N / lcmValue);
    }
  }

  return count;
}

function Main(input) {
  const lines = input.trim().split("\n");
  const [N, K] = lines[0].split(" ").map(Number);
  const V = lines[1].split(" ").map(Number);

  const result = countMultiples(N, K, V);
  console.log(result);
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
