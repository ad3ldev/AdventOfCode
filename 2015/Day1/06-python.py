def main():
    f = open("input.txt")
    content = f.read()
    count = 0
    for character in content: 
        if character == '(':
            count += 1
        elif character == ')':
            count -=1
    f.close()
    print(count)

if __name__ == '__main__':
    main()
