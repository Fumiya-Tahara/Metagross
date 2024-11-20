function Main(input) {
  const lines = input.trim().split("\n");
  const N = parseInt(lines[0], 10);
  const B = lines[1].split(" ").map(Number);

  let A = new Array(N).fill(0);

  A[0] = B[0];
  A[N-1] = B[N-2];

  for (let i = 1; i < N - 1; i++) {
    A[i] = Math.min(B[i - 1], B[i]);
  }

  const sumA = A.reduce((acc, val) => acc + val, 0);

  console.log(sumA);
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));

//             A [0, 0, 0]
//3 B[2, 5] => A[2, 0, 5]
//          => A[2, 2, 5] => 9