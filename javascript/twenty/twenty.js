export const validParens = (message) => {
    let stack = [];
    const openParens = new Set(['(', '{', '[']);
    for (let i = 0; i < message.length; i++) {
        if (openParens.has(message[i])) {
            stack.push(message[i]);
            console.log(stack);
            console.log(stack.length);
        }
        else {
            if (stack.length == 0) {
                return false;
            }

            let openParen = stack.pop();
            if (message[i] == ')' && openParen != '(') {
                return false;
            } else if (message[i] == '}' && openParen != '{') {
                return false;
            } else if (message[i] == ']' && openParen != '[') {
                return false;
            }
        }
    }
    return stack.length == 0;
}

export const SockPairs = (nums) => {
    let pairs = 0;
    let freqs = {};

    for (let num of nums) {
        if (num in freqs) {
            freqs[num] += 1;
            if (freqs[num] % 2 == 0) {
                pairs += 1;
            }
        } else {
            freqs[num] = 1;
        }
    }
    
    return pairs;
}


export const PlusMinus = (nums) => {
    let freqs = {
        'pos': 0,
        'neg': 0,
        'zero': 0
    };

    for (let num of nums) {
        if (num > 0) {
            freqs['pos'] += 1;
        } else if (num == 0) {
            freqs['zero'] += 1;
        } else {
            freqs['neg'] += 1;
         }
    }

    let totalNums = nums.length;

    for (let key of Object.keys(freqs)) { 
        console.log((freqs[key] / totalNums).toFixed(6));
    }
}

export const MinMaxSum = (nums) => {
    nums.sort();
    let maxSum = nums.slice(-4).reduce((accumulator, currVal) => {
        return accumulator + currVal;
    }, 0);
    let minSum = nums.slice(0, 4).reduce((accumulator, currVal) => {
        return accumulator + currVal;
    }, 0);
    console.log(minSum, maxSum);
}

export const LonelyInt = (nums) => {
    let numsMap = {};
    for (let num of nums) {
        if (num in numsMap) {
            numsMap[num] += 1;
        } else {
            numsMap[num] = 1;
        }
    }

    for (let key of Object.keys(numsMap)) {
        if (numsMap[key] === 1) {
            return key;
        }
    }
}

export const DiagonalDifference = (matrix) => {
    let topDown = 0;
    let botUp = 0;
    let row = matrix.length - 1;

    for (let col in matrix) {
        topDown += matrix[col][col];
        botUp += matrix[row][col];
        row -= 1;
    }
    return Math.abs(topDown - botUp);
}