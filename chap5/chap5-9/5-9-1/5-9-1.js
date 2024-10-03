function Main(input) {
  const lines = input.trim().split("\n");
  let N = BigInt(lines[0], 10);

  let count = 0;
  count += Number(N / 10000n);
  N %= 10000n;
  count += Number(N / 5000n);
  N %= 5000n;
  count += Number(N / 1000n);
  N %= 1000n;

  console.log(count);
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
