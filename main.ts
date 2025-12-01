/**
 * @author Rabira Motuma
 * @version 1.0.0
 * @date 2025-11-27
 * @fileoverview This program converts cents to dollars
 */

// variables
let numberOfCentsAsString: string;
let numberOfCentsAsNumber: number;
let dollars: number;
let cents: number;

// input
numberOfCentsAsString = prompt("Input the cents please:") || "0";

// process
numberOfCentsAsNumber = parseInt(numberOfCentsAsString, 10);
dollars = Math.floor(numberOfCentsAsNumber / 100);     // https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Math/floor
cents = numberOfCentsAsNumber % 100;                   // gets remainder

// output
console.log(`\nThat is ${dollars} dollars and ${cents} cents`)

console.log("\nDone.")
