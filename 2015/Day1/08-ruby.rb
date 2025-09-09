content = File.read("input.txt")
count = 0
content.split("").each do |i|
  if i == '(' then
    count+=1
  elsif i == ')' then
    count -=1
  end
end
puts count
