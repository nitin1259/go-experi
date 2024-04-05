/*

Given an array of asynchronous functions functions, return a new promise promise. Each function in the array accepts no arguments and returns a promise. All the promises should be executed in parallel.
 
promise resolves:
 
When all the promises returned from functions were resolved successfully in parallel. The resolved value of promise should be an array of all the resolved values of promises in the same order as they were in the functions. The promise should resolve when all the asynchronous functions in the array have completed execution in parallel.
promise rejects:
 
When any of the promises returned from functions were rejected. promise should also reject with the reason of the first rejection.
Please solve it without using the built-in Promise.all function.
 
 
Input: functions = [
  () => new Promise(resolve => setTimeout(() => resolve(5), 200)),
  () => new Promise(resolve => setTimeout(() => resolve(10), 100))
]
Output: {"t": 200, "resolved": [5, 10]}

Input: functions = [
    () => new Promise(resolve => setTimeout(() => resolve(1), 200)), 
    () => new Promise((resolve, reject) => setTimeout(() => reject("Error"), 100))
]
Output: {"t": 100, "rejected": "Error"}

*/

function ExecuteAysPromiseFunc(promises) {
  return new Promise((resolve, reject) => {
    // validte the promise
    if (!Array.isArray(promises)) {
      reject(new TypeError("promises is not an array"));
      return;
    }

    // declare output array
    const results = [];

    // counter for Promise reject resolve
    let completedPromises = 0;

    function hadleCompletion(i, val) {
      results[i] = val;
      completedPromises++;

      if (completedPromises == promises.length) {
        resolve({ t: 200, resolved: results.join(",") });
      }
    }

    function handleRejection(reason) {
      reject({ t: 100, rejected: reason });
    }

    // iterate the promise array
    // check for reject / resolve status

    for (let i = 0; i < promises.length; i++) {
      const currPromise = promises[i]();

      if (!(currPromise instanceof Promise)) {
        reject(new TypeError("Not a promise"));
        return;
      }

      currPromise
        .then((val) => hadleCompletion(i, val))
        .catch(handleRejection("Error"));
    }
  });
}

const promise1 = Promise.resolve(1);

ExecuteAysPromiseFunc([
  () => new Promise((resolve) => setTimeout(() => resolve(5), 200)),
  () => new Promise((resolve) => setTimeout(() => resolve(10), 100)),
])
  .then((val) => {
    console.log(val);
  })
  .catch((err) => {
    console.log(err);
  });
