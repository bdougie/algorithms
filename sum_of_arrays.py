
def main(array):
    sum = 0
    for num in array:
        sum += num
    return (sum == 0)

main([1, 1, -2])

# real	0m0.023s
# user	0m0.011s
# sys	0m0.008s
