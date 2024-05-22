//演習問題集用
function Main(input) {
  input = input.split("\n");
  let N = Number(input[0]);
  let A = input[1].split(" ")
  let dp1=new Array( N + 1 ).fill(0);
  let dp2=new Array( N + 1 ).fill(0);;
  
  Power(1);
  
  function Power(i) {
  if (i <= N) {
      console.log(dp1[i-1])
      console.log(dp2[i-1])
      dp1[i] = dp2[i - 1] + Number(A[i - 1]);
      dp2[i] = Math.max(dp1[i - 1], dp2[i - 1]);
      Power(i + 1);
  } else {
    console.log(Math.max(dp1[N], dp2[N]));
  }
}

}