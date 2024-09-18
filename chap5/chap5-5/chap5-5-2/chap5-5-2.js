//O(N^3)
function Main(input) {
  const lines = input.trim().split("\n");
  const N = parseInt(lines[0], 10);
  const conditions = lines.slice(1).map((line) => line.split(" ").map(Number));
	//conditions = 条件
  let maxSum = -Infinity;

  for (let i = 0; i < N; i++) {
    for (let j = i + 1; j < N; j++) {
      const [a1, b1, c1] = conditions[i];
      const [a2, b2, c2] = conditions[j];

      const denominator = a1 * b2 - a2 * b1;
			//denominator = 分母
      if (denominator === 0) continue;

      const x = (c1 * b2 - c2 * b1) / denominator;
      const y = (a1 * c2 - a2 * c1) / denominator;

      let valid = true;
      for (let k = 0; k < N; k++) {
        const [a, b, c] = conditions[k];
        if (a * x + b * y > c) {
          valid = false;
          break;
        }
      }

      if (valid) {
        maxSum = Math.max(maxSum, x + y);
      }
    }
  }

  console.log(maxSum);
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
