local file = io.open("input.txt", "r")
if file == nil then
	return 0
end
local content = file:read("*a")
local count = 0
for c in content:gmatch(".") do
	if c == "(" then
		count = count + 1
	elseif c == ")" then
		count = count - 1
	end
end
file:close()
print(count)
