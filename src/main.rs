mod leap;

fn main() {
    let years = [1900, 1996, 1997, 2000];
    for year in years.iter() {
        if leap::is_leap(*year){
            println!("{} is leap year", year);
        } else {
            println!("{} is NOT leap year", year);
        }
    }
}
