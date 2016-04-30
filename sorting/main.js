function quickSort(items, left, right) {
	if (items.length > 1) {
		left = left.length < 1 ? 0 : left;
		right = right.length < 1 ? items.length - 1 : right;

		var index = partition(items, left, right);

		if (left < index - 1) {
				quickSort(items, left, index - 1);
		}

		if (index < right) {
				quickSort(items, index, right);
		}
	}

	return items;
}

function swap(items, firstIndex, secondIndex){
    var temp = items[firstIndex];
    items[firstIndex] = items[secondIndex];
    items[secondIndex] = temp;
}

function partition(items, left, right) {
    var pivot   = items[Math.floor((right + left) / 2)],
        i       = left,
        j       = right;

    while (i <= j) {
			while (items[i] < pivot) {
					i++;
			}

			while (items[j] > pivot) {
					j--;
			}

			if (i <= j) {
					swap(items, i, j);
					i++;
					j--;
			}
    }

    return i;
}

let list = ["C", "B", "F", "E", "G", "A", "D"];
console.log(quickSort(list, [], []));
