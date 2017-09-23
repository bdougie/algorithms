// Still doesnt solve for "somethingLikeABCThis"
// function wordCount(input) {
//   return input.replace(/([a-z])([A-Z])/g, '$1 $2').split(" ").length
// }

function wordCount(input) {
  if (input == "") {return 0}

  let wordCount = 1; // there will at least be 1 one word

  for (var i = 0; i < input.length; i++) {
    if (input[i] == input[i].toUpperCase()) {
      if (input[i - 1] != input[i - 1].toUpperCase()) { // if previous letter is not a cap then add to count, i.e. ABC
        wordCount++
      }
    }
  }

  return wordCount;
}

console.log(wordCount("somethingABCCamelCase"))

