mod ten;

fn main() {
    println!("{}", ten::remove_dups(&mut vec![0,0,1,1,1,2,2,3,3,4]));
    println!("{:?}", ten::two_sum(&mut vec![2,7,11,15], 26));
}
