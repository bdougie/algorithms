require 'benchmark'

# def sum_fast(array)
#   # speed in real time (0.000008)
#   sum = array.inject(0, &:+)
#   return true if sum == 0
#   return false
# end

def sum_fast(array)
  # speed in real time (0.000005)
  sum = 0
  array.each do |i|
    sum += i
  end
  return true if sum == 0
  return false
end

# puts sum_fast([1, 1, -2])
print Benchmark.measure { sum_fast([1, 1, -2]) }
