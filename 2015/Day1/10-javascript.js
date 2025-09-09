let count = 0;

const path = "./input.txt";
const file = Bun.file(path);
const content = await file.text();

content.split("").forEach((c) => {
  if (c == "(") {
    count += 1;
  } else if (c == ")") {
    count -= 1;
  }
});

console.log(count);
