console.log('Starting async tests...');

let completed = false;

setTimeout(() => {
    completed = true;
    console.log('Timeout completed');
}, 100);

// Test that setTimeout works
console.assert(!completed, 'Should not complete immediately');
