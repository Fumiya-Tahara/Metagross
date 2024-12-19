function Main(input) {
  const lines = input.trim().split("\n");
  const Q = parseInt(lines[0], 10);
  const results = [];

  function multiplyMatrices(A, B) {
    const C = Array.from({ length: 3 }, () => Array(3).fill(0));
    for (let i = 0; i < 3; i++) {
      for (let j = 0; j < 3; j++) {
        for (let k = 0; k < 3; k++) {
          C[i][j] += A[i][k] * B[k][j];
        }
      }
    }
    return C;
  }

  function matrixPower(matrix, power) {
    let result = [
      [1, 0, 0],
      [0, 1, 0],
      [0, 0, 1]
    ];
    let base = matrix;

    while (power > 0) {
      if (power % 2 === 1) {
        result = multiplyMatrices(result, base);
      }
      base = multiplyMatrices(base, base);
      power = Math.floor(power / 2);
    }

    return result;
  }

  for (let i = 1; i <= Q; i++) {
    const [X, Y, Z, T] = lines[i].split(" ").map(Number);
    const transformationMatrix = [
      [1 - X, Y, 0],
      [0, 1 - Y, Z],
      [X, 0, 1 - Z]
    ];

    const finalMatrix = matrixPower(transformationMatrix, T);

    const a = finalMatrix[0][0] + finalMatrix[0][1] + finalMatrix[0][2];
    const b = finalMatrix[1][0] + finalMatrix[1][1] + finalMatrix[1][2];
    const c = finalMatrix[2][0] + finalMatrix[2][1] + finalMatrix[2][2];

    results.push(`${a.toFixed(7)} ${b.toFixed(7)} ${c.toFixed(7)}`);
  }

  console.log(results.join("\n"));
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));