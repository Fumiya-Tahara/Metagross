function Main(input) {
	const lines = input.trim().split("\n");
	const N = BigInt(lines[0], 10);

  const sum = ((N - 1n) * N) / 2n;
  console.log(sum.toString(10));
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
