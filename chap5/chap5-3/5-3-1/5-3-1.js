function Main(input) {
  const [H, W] = input.trim().split(" ").map(Number);
  if (H === 1 || W === 1) {
    console.log(1);
  } else {
    console.log(Math.floor((H * W + 1) / 2));
  }
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));