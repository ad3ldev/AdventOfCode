defmodule ParenCounter do
  def count_parens(filename) do
    case File.read(filename) do
      {:ok, content} ->
        count = String.graphemes(content)
                |> Enum.reduce(0, fn char, acc ->
                  case char do
                    "(" -> acc + 1
                    ")" -> acc - 1
                    _ -> acc
                  end
                end)
        IO.puts(count)
      
      {:error, reason} ->
        IO.puts("Error: #{reason}")
    end
  end
end

# Run the function
ParenCounter.count_parens("input.txt")
