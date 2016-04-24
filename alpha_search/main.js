function linearSearchIterations(list, value) {
	iterations = 0;

	for (i = 0; i < list.length; i++) {
		iterations = iterations + 1;

		if (list[i] === value) {
			return iterations;
		}
	}
	return iterations
}

function binarySearchIterations(list, value) {
	var low, mid, high, compare;
	low = 0; high = list.length - 1; mid = (low + high) / 2;
	compare = list[mid].localeCompare(value);

	switch (compare) {
		case 0: low = mid;
			break;
		case 1: high = mid - 1;
			break;
		case -1: low = mid + 1;
			break;
		default: return mid;
	}

	list = list.slice(low, high + 1)
	return linearSearchIterations(list, value)
}

var list = ["A", "B", "C", "D", "E", "F", "G"];

console.log("linear", linearSearchIterations(list, "G"));
console.log("binary", binarySearchIterations(list, "G"));

console.log("binary", binarySearchIterations(list, "D"));
