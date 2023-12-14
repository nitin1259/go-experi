function flattenArrayByLevel(arr, level) {
  let result = [];

  function flattenHelper(currentArray, currentLevel) {
    for (let i = 0; i < currentArray.length; i++) {
      if (Array.isArray(currentArray[i]) && currentLevel < level) {
        // Recursively flatten nested arrays up to the specified level
        result = flattenHelper(currentArray[i], currentLevel + 1);
      } else {
        // Add non-array elements to the result
        result.push(currentArray[i]);
      }
    }
    return result;
  }

  return flattenHelper(arr, 1);
}

const input = [1, 2, 3, [4, 5, 6], [7, 8, [9, 10, 11], 12], [13, 14, 15]];
const output = flattenArrayByLevel(input, 2);

console.log(output);
