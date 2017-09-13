function sieve(num) {
  let array = new Array(num + 1).fill(true) // create and fill an array to num
  let upperLimit = Math.sqrt(num)
  let circledPrimes = []

  // cross out multiples
  for(var i = 2; i <= upperLimit; i++) {
    if (array[i]) {
      for (var j = i * i; j < num; j += i) {
         array[j] = false;
      }
    }
  }

  // circle primes
  for (var i = 2; i < num; i++) {
    if (array[i]) {
      circledPrimes.push(i)
    }
  }

  return circledPrimes
}

console.log(sieve(100))
// sieve(100)

