require 'benchmark'
require 'pry'

def sort(*arrays)
  flat_array = make_flatter!(arrays)
  mark_sort(flat_array)
end

def make_flatter!(arr, result=[])
  arr.each { |a| result << a }
  result[0]
end

def mark_sort(array)
  array_max = array.max
  array_min = array.min
  markings = [0] * (array_max - array_min + 1)
  array.each do |a|
    markings[a - array_min] += 1
  end

  res = []
  markings.length.times do |i|
    markings[i].times do
      res << i + array_min;
    end
  end

  res
end

max = 10000
size = 10000

a = Array.new(size) { rand(max) }

print Benchmark.measure { sort(a) }
