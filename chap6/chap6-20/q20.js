function Main(input) {
  const lines = input.trim().split("\n");
  const N = parseInt(lines[0], 10);
  const C = [];
  const P = [];
  
  for (let i = 1; i <= N; i++) {
    const [c, p] = lines[i].split(" ").map(Number);
    C.push(c);
    P.push(p);
  }

  const Q = parseInt(lines[N + 1], 10);
  const queries = [];
  for (let i = N + 2; i < N + 2 + Q; i++) {
    const [L, R] = lines[i].split(" ").map(Number);
    queries.push([L, R]);
  }

  const sum1 = new Array(N + 1).fill(0);
  const sum2 = new Array(N + 1).fill(0);

  for (let i = 0; i < N; i++) {
    sum1[i + 1] = sum1[i];
    sum2[i + 1] = sum2[i];
    if (C[i] === 1) {
      sum1[i + 1] += P[i];
    } else {
      sum2[i + 1] += P[i];
    }
  }

  const results = [];
  for (const [L, R] of queries) {
    const sumClass1 = sum1[R] - sum1[L - 1];
    const sumClass2 = sum2[R] - sum2[L - 1];
    results.push(`${sumClass1} ${sumClass2}`);
  }

  console.log(results.join("\n"));
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));