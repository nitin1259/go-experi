/**
 * 

Write a program to Find k closest elements to a given value.

Input: K = 4, X = 35

arr[] = {7,8,9,12, 16, 22, 30, 35, 39, 42, 45, 48, 50, 53, 55, 56}

Output: 30 39 42 45
 */

function findKclosetElement(arr, k, target) {
  let left = 0;
  let right = arr.length - 1;

  while (left < right) {
    let mid = Math.floor((left + right) / 2);
    if (target > arr[mid]) {
      left = mid + 1;
    } else {
      right = mid;
    }
  }

  let leftPtr = left - 1;
  let rightPtr = left;

  while (k > 0) {
    if (leftPtr < 0) {
      rightPtr++;
    } else if (rightPtr >= arr.length) {
      leftPtr--;
    } else {
      rightPtr++;
    }
    k--;
  }

  return arr.slice(leftPtr, rightPtr);
}

const arr = [
  7, 8, 10, 12, 16, 22, 30, 31, 35, 37, 39, 42, 45, 48, 50, 53, 55, 56,
];
const k = 4;

const X = 35;

console.log(findKclosetElement(arr, k, X));
