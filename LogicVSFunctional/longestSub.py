def longestCommonPrefix(myList):
    firstWord = myList[0]
    otherWords = myList[1:]
    commonLetters = firstWord

    for word in otherWords:
        for i in range(0, len(firstWord)):
            if word[i] == commonLetters[i]:
                commonLetters += word[i]
            else:
                commonLetters = commonLetters[:i]
                break
    return commonLetters


def main():
    myList = ["interview",  "interrupt",  "integrate",  "intermediate"]
    print(longestCommonPrefix(myList))


main()
