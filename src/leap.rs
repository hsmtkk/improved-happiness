pub fn is_leap(year: i64) -> bool {
	if year%400 == 0 {
		return true;
	}
	if year%100 == 0 {
		return false;
	}
	if year%4 == 0 {
		return true;
	}
	return false;
}

#[test]
fn test_is_leap(){
    assert!(!is_leap(1900));
    assert!(is_leap(1996));
    assert!(!is_leap(1997));
    assert!(is_leap(2000));
}