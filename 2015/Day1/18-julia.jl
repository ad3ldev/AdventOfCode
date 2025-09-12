function count_parens(filename)
    try
        content = read(filename, String)
        count = 0
        
        for char in content
            if char == '('
                count += 1
            elseif char == ')'
                count -= 1
            end
        end
        
        println(count)
    catch e
        println("Error: $e")
    end
end

# Run the function
count_parens("input.txt")
