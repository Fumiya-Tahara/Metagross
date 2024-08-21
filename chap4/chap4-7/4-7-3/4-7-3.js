function Matrix() {
  return {
    p: [
      [0n, 0n, 0n],
      [0n, 0n, 0n],
      [0n, 0n, 0n]
    ]
  };
}

function Multiplication(A, B) {
  const mod = 1000000007n;
  const C = Matrix();
  for (let i = 0; i < 3; i++) {
    for (let k = 0; k < 3; k++) {
      for (let j = 0; j < 3; j++) {
        C.p[i][j] += A.p[i][k] * B.p[k][j];
        C.p[i][j] %= mod;
      }
    }
  }
  return C;
}

function Power(A, n) {
  let P = A;
  let Q;
  let flag = false;
  for (let i = 0n; i < 60n; i++) {
    if ((n & (1n << i)) !== 0n) {
      if (!flag) {
        Q = P;
        flag = true;
      } else {
        Q = Multiplication(Q, P);
      }
    }
    P = Multiplication(P, P);
  }
  return Q;
}

function Main(input) {
  const N = BigInt(input.trim());
  const mod = 1000000007n;

  if (N === 1n || N === 2n) {
    console.log(1);
    return;
  } else if (N === 3n) {
    console.log(2);
    return;
  }

  const A = Matrix();
  A.p[0][0] = 1n;
  A.p[0][1] = 1n;
  A.p[0][2] = 1n;
  A.p[1][0] = 1n;
  A.p[2][1] = 1n;

  const B = Power(A, N - 3n);

  const result = (B.p[0][0] * 2n + B.p[0][1] * 1n + B.p[0][2] * 1n) % mod;
  console.log(result.toString());
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));