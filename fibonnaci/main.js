function fib(n) {
  if ([0,1].includes(n)) {
    return n;
  }

  return fib(n-1) + fib(n -2);
}

for (i = 0; i < 10; i++) {
  console.log(fib(i));
}
