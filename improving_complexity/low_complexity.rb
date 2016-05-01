require 'benchmark'
require 'pry'

def sort(*arrays)
  flat_array = make_flatter!(arrays)
  quick_sort(flat_array)
end

def make_flatter!(arr, result=[])
  arr.each { |a| result << a }
  result[0]
end

def quick_sort(array)
  return array if array.length <= 1

  pivot_index = (array.length / 2).to_i
  pivot_value = array[pivot_index]
  array.delete_at(pivot_index)

  lesser = []
  greater = []

  array.each do |x|
    if x <= pivot_value
      lesser << x
    else
      greater << x
    end
  end

  return quick_sort(lesser) + [pivot_value] + quick_sort(greater)
end

max = 10000
size = 10000

a = Array.new(size) { rand(max) }

print Benchmark.measure { sort(a) }
