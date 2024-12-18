function main(input) {
  input = input.trim().split('\n');
  let n = parseInt(input[0]);
  let c = input[1];
  
  function ncr(x, y) {
    if (x < 3 && y < 3) {
      if (x < y) return 0;
      if (x == 2 && y == 1) return 2;
      return 1;
    }
    return ncr(Math.floor(x / 3), Math.floor(y / 3)) * ncr(x % 3, y % 3) % 3;
  }

  let ans = 0;
  for (let i = 0; i < n; i++) {
    let code = -1;
    if (c[i] == 'B') code = 0;
    if (c[i] == 'W') code = 1;
    if (c[i] == 'R') code = 2;
    ans += code * ncr(n - 1, i);
    ans %= 3;
  }
  if (n % 2 == 0) ans = (3 - ans) % 3;
  console.log('BWR'[ans]);
}

main(require("fs").readFileSync("/dev/stdin", "utf8"));