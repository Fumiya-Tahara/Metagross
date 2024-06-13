//要復習

function Main(input) {
  let inputs = input.trim().split("\n");
  let N = Number(inputs[0]);
  let points = [];
  for (let i = 1; i <= N; i++) {
    let [x, y] = inputs[i].split(" ");
    points.push({x, y});
  }

  points.sort((a, b) => a.x - b.x);

  let ans = Infinity;
  for (let i = 0; i < N; i++) {
    for (let j = i + 1; j < N; j++) {
      if ((points[j].x - points[i].x) ** 2 >= ans) break;

      let dist = (points[i].x - points[j].x) ** 2 + (points[i].y - points[j].y) ** 2;
      ans = Math.min(ans, dist);
    }
  }

  console.log(Math.sqrt(ans));
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));