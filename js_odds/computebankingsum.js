/**

Input:
computeAmount().lacs(15).crore(5).crore(2).lacs(20).thousand(45).crore(7).value();

Output:
143545000
*/

function calculator() {
  this.amount = 0;
  this.crore = function (amountToBeAdd) {
    this.amount += amountToBeAdd * 10000000;
    return this;
  };

  this.lacs = function (amountToBeAdd) {
    this.amount += amountToBeAdd * 100000;
    return this;
  };

  this.thousand = function (amountToBeAdd) {
    this.amount += amountToBeAdd * 1000;
    return this;
  };

  this.hundred = function (amountToBeAdd) {
    this.amount += amountToBeAdd * Math.pow(10, 2);
    return this;
  };

  this.ten = function (amountToBeAdd) {
    this.amount += amountToBeAdd * Math.pow(10, 1);
    return this;
  };
  this.unit = function (amountToBeAdd) {
    this.amount += amountToBeAdd * Math.pow(10, 0);
    return this;
  };

  this.value = function () {
    return this.amount;
  };
}

const computeAmount = function () {
  return new calculator();
};

console.log(
  computeAmount()
    .lacs(15)
    .crore(5)
    .crore(2)
    .lacs(20)
    .thousand(45)
    .crore(7)
    .value()
);

// write the closure implementation  of above program ??
