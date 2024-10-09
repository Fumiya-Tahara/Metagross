function Main(input) {
  const lines = input.trim().split("\n");
  const [a, b, c] = lines[0].split(" ").map(BigInt);

  if (c === 1n) {
    console.log("No");
    return;
  }

  let v = 1n;
  for (let i = 1n; i <= b; i++) {
    if (a / c < v) {
      console.log("Yes");
      return;
    }
    v *= c;
  }

  console.log("No");
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
