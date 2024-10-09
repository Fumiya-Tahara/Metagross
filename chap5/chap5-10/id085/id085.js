//for~4だとTLE

function Main(input) {
  const lines = input.trim().split("\n");
  const numbers = lines[0].split(" ").map(BigInt);

  const N = numbers[0];
  const X = numbers[1];
  const Y = numbers[2];

	for(let a=1n; a<=N; a++){
		for(let b=a; b<=N; b++){
			for(let c=b; c<=N; c++){
				const d = X - a - b - c;
        if (d >= c && d <= N && a * b * c * d === Y) {
          console.log("Yes");
          return;
        }
			}
		}
	}
	console.log("No");
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
