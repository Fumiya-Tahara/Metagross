//問題なさそうだけど時間切れ、、
// function main(input) {
//   lines = input.trim().split("\n");
//   let [n, m] = lines[0].split(" ").map((val) => parseInt(val));
//   let a = [],
//     b = [];
//   for (let i = 0; i < m; i++) {
//     [a[i], b[i]] = lines[i + 1].split(" ").map((val) => parseInt(val));
//   }

//   let dist = new Array(100009).fill(-1);
//   let g = new Array(100009).fill().map(() => []);
//   let q = [];
//   for (let i = 0; i < m; i++) {
//     g[a[i]].push(b[i]);
//     g[b[i]].push(a[i]);
//   }
//   q.push(1);
//   dist[1] = 0;

//   while (q.length != 0) {
//     let pos = q[0];
//     q.shift();
//     for (let i = 0; i < g[pos].length; i++) {
//       let nex = g[pos][i];
//       if (dist[nex] == -1) {
//         dist[nex] = dist[pos] + 1;
//         q.push(nex);
//       }
//     }
//   }
//   for (let i = 1; i <= n; i++) {
//     if (dist[i] == -1) {
//       console.log(120);
//     } else {
//       console.log(Math.min(dist[i], 120));
//     }
//   }
// }
// main(require("fs").readFileSync("/dev/stdin", "utf8"));

const input = require("fs").readFileSync("/dev/stdin", "utf-8").split("\n");

const [N, M] = input[0].split(" ").map((v) => parseInt(v));

const G = Array.from({ length: N }, (_) => []);

for (let i = 1; i <= M; i++) {
  const [A, B] = input[i].split(" ").map((v) => parseInt(v) - 1);
  G[A].push(B);
  G[B].push(A);
}

const Q = {
  d: Array(N),
  h: 0,
  t: 0,
  len: () => Q.t - Q.h,
  put: (v) => (Q.d[Q.t++] = v),
  get: () => (Q.len() ? Q.d[Q.h++] : null),
};

const dist = Array(N).fill(-1);
dist[0] = 0;
Q.put(0);

while (Q.len()) {
  const i = Q.get();

  G[i].forEach((j) => {
    if (dist[j] !== -1) return;

    dist[j] = dist[i] + 1;
    Q.put(j);
  });
}

for (let i = 0; i < N; i++) {
  console.log(dist[i] === -1 ? 120 : Math.min(dist[i], 120));
}
