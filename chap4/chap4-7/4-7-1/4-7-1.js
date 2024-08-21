function Matrix() {
  return {
    p: [
      [0n, 0n],
      [0n, 0n],
    ],
  };
}

function Multiplication(A, B) {
  const C = Matrix();
  for (let i = 0; i < 2; i++) {
    for (let k = 0; k < 2; k++) {
      for (let j = 0; j < 2; j++) {
        C.p[i][j] += A.p[i][k] * B.p[k][j];
        C.p[i][j] %= 1000000000n;
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
  if (N === 1n || N === 2n) {
    console.log(1);
    return;
  }

  const A = Matrix();
  A.p[0][0] = 1n;
  A.p[0][1] = 1n;
  A.p[1][0] = 1n;

  const B = Power(A, N - 1n);

  console.log(B.p[0][0].toString());
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
