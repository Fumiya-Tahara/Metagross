function Main(input) {
  const lines = input.trim().split("\n");
  const [L,R] = lines[0].split(" ").map(Number);

  const limit = Math.floor(Math.sqrt(R)) + 1;
  const isPrime = Array(limit).fill(true);
  isPrime[0,1] = false;

  for (let i = 2; i < limit; i++) {
    if (isPrime[i]) {
      for (let j = i * i; j < limit; j += i) {
        isPrime[j] = false;
      }
    }
  }

  const isPrimeRange = Array(R - L + 1).fill(true);
  for (let i = 2; i < limit; i++) {
    if (isPrime[i]) {
      const start = Math.max(i * i, Math.ceil(L / i) * i);
      for (let j = start; j <= R; j += i) {
        isPrimeRange[j - L] = false;
      }
    }
  }

  let count = 0;
  for (let i = 0; i <= R - L; i++) {
    if (isPrimeRange[i] && (i + L) > 1) {
      count++;
    }
  }

  console.log(count);
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
