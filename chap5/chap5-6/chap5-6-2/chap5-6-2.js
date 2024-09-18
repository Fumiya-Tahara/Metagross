function Main(input) {
  const lines = input.trim().split("\n");
  const N = BigInt(lines[0]);
  const A = lines[1].split(" ").map(BigInt);
  const mod = BigInt(1000000007);

  let result = BigInt(0);
  let ninoruijo = BigInt(1);
  for (let i = 0; i < N; i++) {
    result = (result + A[i] * ninoruijo) % mod;
    ninoruijo = (ninoruijo * BigInt(2)) % mod;
  }
  console.log(result.toString());
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
