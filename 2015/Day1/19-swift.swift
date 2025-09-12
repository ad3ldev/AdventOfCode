import Foundation

func countParens(filename: String) {
    do {
        let content = try String(contentsOfFile: filename, encoding: .utf8)
        var count = 0
        
        for char in content {
            switch char {
            case "(":
                count += 1
            case ")":
                count -= 1
            default:
                break
            }
        }
        
        print(count)
    } catch {
        print("Error: \(error)")
    }
}

// Run the function
countParens(filename: "input.txt")
