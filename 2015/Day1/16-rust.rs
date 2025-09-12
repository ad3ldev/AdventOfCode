use std::fs;

fn main() {
    let filename = "input.txt";

    match fs::read_to_string(filename) {
        Ok(content) => {
            let mut count = 0;
            for ch in content.chars() {
                match ch {
                    '(' => count += 1,
                    ')' => count -= 1,
                    _ => {}
                }
            }
            println!("{}", count);
        }
        Err(e) => eprintln!("Error: {}", e),
    }
}
