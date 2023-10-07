mod ten;
mod twenty;
fn main() {
    println!("{}", ten::remove_dups(&mut vec![0,0,1,1,1,2,2,3,3,4]));
    println!("{:?}", ten::two_sum(&mut vec![2,7,11,15], 26));
    println!("{}", ten::valid_palind("A man, a plan, a canal: Panama".to_string()));
    println!("{}", ten::newtons_sqroot(327, 0.00001));
    println!("{}", ten::contains_dups(vec![1,1,1,3,3,4,3,2,4,2]));
    println!("{}", ten::contains_dups_v2(vec![1,1,1,3,3,4,3,2,4,2]));
    println!("{:?}", ten::group_anagrams(vec![
        "eat".to_string(),
        "tea".to_string(),
        "tan".to_string(),
        "ate".to_string(),
        "nat".to_string(),
        "bat".to_string()]));
    println!("{:?}", ten::top_k_freq(vec![1,1,1,2,2,3], 2));
    println!("{:?}", ten::longest_substr_len("abcabcbb".to_string()));
    println!("{:?}", ten::longest_substr_len("pwwkew".to_string()));
    println!("{:?}", ten::longest_substr_len("bbbbb".to_string()));
    println!("{:?}", ten::cypher("abcdefghijklmnopqrstuvwxyz".to_string(), 3));
    println!("{:?}", ten::cypher("There's-a-starman-waiting-in-the-sky".to_string(), 3));
    println!("{}", ten::is_palindrome(121));
    println!("{}", ten::is_palindrome(12));
    println!("{}", twenty::valid_parens("()[]{}".to_string()));
    println!("{}", twenty::valid_parens("(]".to_string()));
}
