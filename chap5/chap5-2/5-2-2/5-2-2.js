function Main(input) {
  const N = parseInt(input.trim(), 10);
	let ans = "";
	(N % 3 === 0) ? ans = "First" : ans = "Second";
  console.log(dp[N] ? "First" : "Second");
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
