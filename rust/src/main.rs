mod ten;

fn main() {
    println!("{}", ten::remove_dups(&mut vec![0,0,1,1,1,2,2,3,3,4]));
    println!("{:?}", ten::two_sum(&mut vec![2,7,11,15], 26));
<<<<<<< Updated upstream
    println!("{}", ten::valid_palind("A man, a plan, a canal: Panama".to_string()));
=======
    println!("{}", ten::newtons_sqroot(327, 0.00001));
    println!("{}", ten::contains_dups(vec![1,1,1,3,3,4,3,2,4,2]));
    println!("{}", ten::contains_dups_v2(vec![1,1,1,3,3,4,3,2,4,2]));
>>>>>>> Stashed changes
}
