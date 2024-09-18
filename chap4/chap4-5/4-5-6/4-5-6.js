function Main(input) {
  const lines = input.trim().split("\n");
  const [H, W] = lines[0].split(" ").map(Number);
  const [sx, sy] = lines[1].split(" ").map(Number);
  const [gx, gy] = lines[2].split(" ").map(Number);
  const c = lines.slice(3, 3 + H).map((line) => line.split(""));

  const start = (sx - 1) * W + (sy - 1);
  const goal = (gx - 1) * W + (gy - 1);

  const dist = Array(H * W).fill(-1);
  const G = Array.from({ length: H * W }, () => []);

  for (let i = 0; i < H; i++) {
    for (let j = 0; j < W - 1; j++) {
      const idx1 = i * W + j;
      const idx2 = i * W + (j + 1);
      if (c[i][j] === "." && c[i][j + 1] === ".") {
        G[idx1].push(idx2);
        G[idx2].push(idx1);
      }
    }
  }

  for (let i = 0; i < H - 1; i++) {
    for (let j = 0; j < W; j++) {
      const idx1 = i * W + j;
      const idx2 = (i + 1) * W + j;
      if (c[i][j] === "." && c[i + 1][j] === ".") {
        G[idx1].push(idx2);
        G[idx2].push(idx1);
      }
    }
  }

  const queue = [start];
  dist[start] = 0;

  while (queue.length > 0) {
    const v = queue.shift();
    for (const nv of G[v]) {
      if (dist[nv] === -1) {
        dist[nv] = dist[v] + 1;
        queue.push(nv);
      }
    }
  }

  console.log(dist[goal]);
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
