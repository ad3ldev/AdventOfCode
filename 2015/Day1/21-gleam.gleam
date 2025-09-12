import gleam/io
import gleam/result
import gleam/string
import simplifile

pub fn main() {
  let filename = "input.txt"

  case simplifile.read(filename) {
    Ok(content) -> {
      let count =
        content
        |> string.to_graphemes()
        |> count_parens(0)

      io.println(int.to_string(count))
    }
    Error(_) -> io.println("Error: Could not read file")
  }
}

fn count_parens(chars: List(String), acc: Int) -> Int {
  case chars {
    [] -> acc
    ["(" as char, ..rest] -> count_parens(rest, acc + 1)
    [")" as char, ..rest] -> count_parens(rest, acc - 1)
    [_, ..rest] -> count_parens(rest, acc)
  }
}
