def sieve(num)
  # creating an array the size of num
  two_to_num = (2..num).to_a

  selected_prime_nums = []
  multiples = []

  two_to_num.each do |value|
    unless multiples.include?(value)
      selected_prime_nums << value  #circle primes
      multiples.concat((value..num).step(value).to_a) # cross out multiples as possible solutions
    end
  end

  selected_prime_nums
end

print sieve(100)
