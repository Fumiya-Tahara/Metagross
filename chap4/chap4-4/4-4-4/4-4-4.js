function main() {

  let cnt = 0;
  const LIMIT = 23; //
  let current = 0;

  while (current < LIMIT) {
    cnt += 1;
    current += 1.0 / cnt;
  }

  console.log(cnt);
}

main();//5471312310