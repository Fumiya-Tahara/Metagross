function Main(input) {
  const lines = input.trim().split("\n");
  const line = lines[0].split(" ");

  const A = BigInt(line[0]);
  const B = BigInt(line[1]);
  const C = BigInt(line[2]);
  const MOD = 998244353n;

  const sumA = ((A * (A + 1n)) / 2n) % MOD;
  const sumB = ((B * (B + 1n)) / 2n) % MOD;
  const sumC = ((C * (C + 1n)) / 2n) % MOD;

  const result = (((sumA * sumB) % MOD) * sumC) % MOD;

  console.log(result.toString());
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
