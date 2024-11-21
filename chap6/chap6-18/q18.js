function Main(input) {
  const [A, B] = input.trim().split(" ").map(BigInt);
  const g = gcd(A, B);
  const lcm = (A / g) * B;

  if (lcm > 10n ** 18n) {
    console.log("Large");
  } else {
    console.log(lcm.toString());
  }
}

function gcd(a, b) {
  while (b !== 0n) {
    const tmp = b;
    b = a % b;
    a = tmp;
  }
  return a;
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));