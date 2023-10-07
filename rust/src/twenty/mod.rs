use std::collections::{VecDeque, HashSet};

pub fn valid_parens(parens: String) -> bool {
    let mut stack: VecDeque<char> = VecDeque::new();
    let open_parens: HashSet<char> = HashSet::from(['(', '{', '[']);
    for p in parens.chars() {
        if open_parens.contains(&p) {
            stack.push_back(p.clone());
        } else {
            if stack.len() == 0 {
                return false;
            }
            let open_paren = stack.pop_back().unwrap();
            if p == ')' && open_paren != '(' {
                return false;
            } else if p == '}' && open_paren != '{' {
                return false;
            } else if p == ']' && open_paren != '[' {
                return false;
            }
        }
    }

    stack.len() == 0
}