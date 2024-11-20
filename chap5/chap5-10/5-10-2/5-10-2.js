function Main(input) {
  const N = BigInt(input.trim());
  const MOD = 1000000007n;

  const sum = ((N * (N + 1n)) / 2n) % MOD;
  const result = (sum * sum) % MOD;

  console.log(result.toString());
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
