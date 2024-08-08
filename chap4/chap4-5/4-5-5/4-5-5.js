// https://atcoder.jp/contests/math-and-algorithm/tasks/typical90_bz

//Atcoderでのみ()
function Main(input) {
  const data = input.trim().split("\n");
  const [N, M] = data[0].split(" ").map(Number);

  let graph = [];
  for (let i = 0; i <= N; i++) {
    graph.push([]);
  }

  let G = graph;

  for (let i = 1; i <= M; i++) {
    const [A, B] = data[i].split(" ").map(Number);
    G[A].push(B);
    G[B].push(A);
		// console.log(G);
  }

  let Answer = 0;
  for (let i = 1; i <= N; i++) {
    let cnt = 0;
    for (let j = 0; j < G[i].length; j++) {
      if (G[i][j] < i) cnt += 1;
    }
    if (cnt === 1) Answer += 1;
  }

  console.log(Answer);
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
