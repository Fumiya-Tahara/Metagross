function Main(input) {
  const lines = input.trim().split("\n");
  const N = parseInt(lines[0], 10);
  const intervals = [];

  for (let i = 1; i <= N; i++) {
    const [l, r] = lines[i].split(" ").map(Number);
    intervals.push({ start: l, end: r });
  }

  intervals.sort((a, b) => a.end - b.end);

  let Current_Time = 0;
  let Answer = 0;

  for (let i = 0; i < N; i++) {
    if (intervals[i].start >= Current_Time) {
      Current_Time = intervals[i].end;
      Answer++;
    }
  }

  console.log(Answer);
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
