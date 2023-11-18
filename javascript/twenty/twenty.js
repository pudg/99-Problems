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