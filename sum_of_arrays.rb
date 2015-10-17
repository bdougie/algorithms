require 'benchmark'

# slower example
#
# def sum_fast(array)
#   # speed in real time (0.000008)
#   sum = array.inject(0, &:+)
#   return sum == 0
# end

# faster example
#
def sum_fast(array)
  # speed in real time (0.000004)
  sum = 0
  array.each do |i|
    sum += i
  end

  return sum == 0
end

# puts sum_fast([1, 1, -2])
print Benchmark.measure { sum_fast([1, 1, -2]) }
