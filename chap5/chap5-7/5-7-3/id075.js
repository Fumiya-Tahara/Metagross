//https://atcoder.jp/contests/math-and-algorithm/tasks/math_and_algorithm_bi

function Main(input) {
  input = input.trim().split("\n");
  let n = parseInt(input[0]);
  let a = input[1].split(" ").map((val) => BigInt(val));

  let mod = 10n ** 9n + 7n;

  function modpow(a, b) {
    let p = a;
    let ans = 1n;
    for (let i = 0n; i < 30; i++) {
      if ((b & (1n << i)) != 0) {
        ans *= p;
        ans %= mod;
      }
      p *= p;
      p %= mod;
    }
    return ans;
  }

  function division(a, b) {
    return (a * modpow(b, mod - 2n)) % mod;
  }

  let ans = 0n;
  let fact = Array(200009);

  function ncr(n, r) {
    return division(fact[n], (fact[r] * fact[n - r]) % mod);
  }

  fact[0] = 1n;

  for (let i = 1n; i < 200009; i++) {
    fact[i] = (i * fact[i - 1n]) % mod;
  }

  for (let i = 0; i < n; i++) {
    ans += a[i] * ncr(n - 1, i);
  }

  ans %= mod;
  console.log(ans.toString());
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
