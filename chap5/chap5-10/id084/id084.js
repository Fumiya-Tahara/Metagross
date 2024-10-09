function Main(input) {
  const lines = input.trim().split("\n");
  const numbers = lines[0].split(" ").map(BigInt);

  const a = numbers[0];
  const b = numbers[1];
  const c = numbers[2];

  if (c - a - b < 0n) {
    console.log("No");
  } else if (4n * a * b < (c - a - b) * (c - a - b)) {
    console.log("Yes");
  } else {
    console.log("No");
  }
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
