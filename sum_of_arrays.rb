require 'benchmark'

def sum_fast(array)
  sum = array.inject(0, &:+)
  return true if sum == 0
  return false
end

print Benchmark.measure { sum_fast([1, 1, -2]) }
# puts sum_fast([1, 1, -2])
