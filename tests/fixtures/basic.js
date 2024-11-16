// Basic functionality test
console.log('Starting basic tests...');

// Test arithmetic
const sum = 2 + 2;
console.assert(sum === 4, 'Basic arithmetic failed');

// Test string operations
const str = 'Hello' + ' ' + 'World';
console.assert(str === 'Hello World', 'String concatenation failed');

// Test arrays
const arr = [1, 2, 3];
console.assert(arr.length === 3, 'Array length test failed');

// Test objects
const obj = { name: 'test' };
console.assert(obj.name === 'test', 'Object property test failed');

console.log('Basic tests completed');