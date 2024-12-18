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

/*
//N - (k - 1n) * (j - 1n)は「k 個以上の差を持つ j 個のボールを選ぶために残るボールの数」?

memo
N = 3 の場合を考えます。

k = 1 の場合
k = 1 の場合、ボールの差は 1 以上です。
limit = (N - 1n) / k + 1n = (3 - 1n) / 1n + 1n = 3n
j = 1 の場合
N - (k - 1n) * (j - 1n) = 3 - (1 - 1) * (1 - 1) = 3 - 0 * 0 = 3
この場合、選ぶボールの数は 3 です。
組み合わせは {1}, {2}, {3} の3通りです。
j = 2 の場合
N - (k - 1n) * (j - 1n) = 3 - (1 - 1) * (2 - 1) = 3 - 0 * 1 = 3
この場合、選ぶボールの数は 3 です。
組み合わせは {1, 2}, {1, 3}, {2, 3} の3通りです。
j = 3 の場合
N - (k - 1n) * (j - 1n) = 3 - (1 - 1) * (3 - 1) = 3 - 0 * 2 = 3
この場合、選ぶボールの数は 3 です。
組み合わせは {1, 2, 3} の1通りです。
k = 2 の場合
k = 2 の場合、ボールの差は 2 以上です。
limit = (N - 1n) / k + 1n = (3 - 1n) / 2n + 1n = 2n
j = 1 の場合
N - (k - 1n) * (j - 1n) = 3 - (2 - 1) * (1 - 1) = 3 - 1 * 0 = 3
この場合、選ぶボールの数は 3 です。
組み合わせは {1}, {2}, {3} の3通りです。
j = 2 の場合
N - (k - 1n) * (j - 1n) = 3 - (2 - 1) * (2 - 1) = 3 - 1 * 1 = 2
この場合、選ぶボールの数は 2 です。
組み合わせは {1, 3} の1通りです。
k = 3 の場合
k = 3 の場合、ボールの差は 3 以上です。
limit = (N - 1n) / k + 1n = (3 - 1n) / 3n + 1n = 1n
j = 1 の場合
N - (k - 1n) * (j - 1n) = 3 - (3 - 1) * (1 - 1) = 3 - 2 * 0 = 3
この場合、選ぶボールの数は 3 です。
組み合わせは {1}, {2}, {3} の3通りです。

*/