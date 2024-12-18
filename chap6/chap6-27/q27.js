function main(input) {
  input = input.trim().split('\n');
  let N = BigInt(input[0]);

  const MOD = 1000000007n;
  const LIMIT = 100000n;
  let fact = new Array(Number(LIMIT) + 1).fill(1n);
  let invFact = new Array(Number(LIMIT) + 1).fill(1n);

  for (let i = 1n; i <= LIMIT; i++) {
    fact[Number(i)] = (fact[Number(i - 1n)] * i) % MOD;
  }

  // 階乗の逆元の計算
  invFact[Number(LIMIT)] = modpow(fact[Number(LIMIT)], MOD - 2n, MOD);
  for (let i = LIMIT - 1n; i >= 1n; i--) {
    invFact[Number(i)] = (invFact[Number(i + 1n)] * (i + 1n)) % MOD;
  }

  function modpow(a, b, m) {
    let ans = 1n;
    while (b > 0) {
      if (b & 1n) ans = (ans * a) % m;
      a = (a * a) % m;
      b >>= 1n;
    }
    return ans;
  }

  function ncr(n, r) {
    if (r > n) return 0n;
    return (fact[Number(n)] * invFact[Number(r)] % MOD) * invFact[Number(n - r)] % MOD;
  }

  for (let k = 1n; k <= N; k++) {
    let ans = 0n;
    let limit = (N - 1n) / k + 1n;
    for (let j = 1n; j <= limit; j++) {
      ans += ncr(N - (k - 1n) * (j - 1n), j);
      ans %= MOD;
    }
    console.log(ans.toString());
  }
}

main(require("fs").readFileSync("/dev/stdin", "utf8"));