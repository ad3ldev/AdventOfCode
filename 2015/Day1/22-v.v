import os

fn main() {
    filename := 'input.txt'
    
    content := os.read_file(filename) or {
        println('Error: ${err}')
        return
    }
    
    mut count := 0
    for c in content {
        match c {
            `(` { count++ }
            `)` { count-- }
            else {}
        }
    }
    
    println(count)
}
