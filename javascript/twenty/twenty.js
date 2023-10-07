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