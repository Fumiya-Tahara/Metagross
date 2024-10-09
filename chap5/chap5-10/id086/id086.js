function Main(input) {
  const lines = input.trim().split("\n");
	const N = lines[0];
	const S = lines[1].split("");

	let depth = 0;

	for (let i = 0; i < N; i++){
		if (S[i] === "("){
			depth++;
		} else if (S[i] === ")") {
      depth--;
			if (depth < 0){
				console.log("No");
				return;
			}
    } else {
      console.log("No");
			return;
    }
	}
	console.log(depth === 0 ? "Yes" : "No");
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
