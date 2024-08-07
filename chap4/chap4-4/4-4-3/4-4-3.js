const readline = require("readline");

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

rl.question("Enter N: ", (input) => {
  const N = parseInt(input, 10);

  let F = new Array(N + 1).fill(0);
  let answer = 0;


  for (let i = 1; i <= N; i++) {

    for (let j = i; j <= N; j += i) {
      F[j] += 1;
    }
  }

  for (let i = 1; i <= N; i++) {
    answer += i * F[i];
  }

  console.log(answer);

  rl.close();
});

/*Atcoder用
function Main(input) {
  input = input.trim();
  const N = parseInt(input, 10);
	let F = new Array(N + 1).fill(0);
  let answer = 0;

	for (let i = 1; i <= N; i++) {
    for (let j = i; j <= N; j += i) {
      F[j] += 1;
    }
  }

  for (let i = 1; i <= N; i++) {
    answer += i * F[i];
  }

  console.log(answer);
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
*/
