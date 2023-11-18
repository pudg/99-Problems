import {
    SockPairs,
    validParens
} from "./twenty/twenty.js";



function main() {
    console.log(validParens("()[]{}"));
    console.log(SockPairs([1, 2, 1, 2, 1, 3, 2]))
}
main();