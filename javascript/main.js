
import { 
    removeDups,
    twoSum,
    newtonSquareRoot,
    containsDups,
    validAnagram,
    topKElements,
    validPalind,
    groupAnagrams
} from "./ten/ten.js"


function main() {
    console.log(removeDups([0,0,1,1,1,2,2,3,3,4]))
    console.log(twoSum([2,7,11,15], 18))
    console.log(validPalind("A man, a plan, a canal: Panama"))
    console.log(newtonSquareRoot(327, 0.01))
    console.log(containsDups([1,1,1,3,3,4,3,2,4,2]))
    console.log(validAnagram("anagram", "nagaram"))
    console.log(validAnagram("rat", "car"))
    console.log(topKElements([1,1,1,2,2,3], 2));
    console.log(topKElements([1], 1));
    console.log(groupAnagrams(["eat","tea","tan","ate","nat","bat"]));
}

main()