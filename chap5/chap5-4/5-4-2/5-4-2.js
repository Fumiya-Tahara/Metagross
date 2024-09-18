function Main(input) {
  const lines = input.trim().split("\n");
  const [H, W] = lines[0].split(" ").map(Number);
  const A = lines.slice(1).map((line) => line.split(" ").map(Number));

  const rowSum = new Array(H).fill(0);
  const colSum = new Array(W).fill(0);

  for (let i = 0; i < H; i++) {
    for (let j = 0; j < W; j++) {
      rowSum[i] += A[i][j];
      colSum[j] += A[i][j];
    }
  }

  const result = [];
  for (let i = 0; i < H; i++) {
    const rowResult = [];
    for (let j = 0; j < W; j++) {
      rowResult.push(rowSum[i] + colSum[j] - A[i][j]);
    }
    result.push(rowResult.join(" "));
  }

  console.log(result.join("\n"));
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
