#!/usr/local/bin/python3.9

# [-2^31, 2^31-1]

def reverseDigitsChecker(num):
    # Grab sign and make num positive
    sign = +1 if num > 0 else -1
    num *= sign
    num = str(num)
    num = list(num)
    num = reversed(num)
    num = ''.join(num)
    num = int(num)
    return sign * num


def log10(x):
    d = 0
    while x >= 10:
        x /= 10
        d += 1
    return d


def reverseDigits(num):
    reversedNum = 0
    # Grab sign and make num positive
    sign = +1 if num > 0 else -1
    num *= sign
    # Evaluate the number of digits in the number
    size = log10(num)
    # Loop through the digits
    for i in range(size+1):
        digit = num % 10
        num = int(num / 10)
        reversedNum += digit * (10 ** size)
        size -= 1
    
    return reversedNum * sign



for i in range(-2**31, 2**31):
    A = reverseDigits(i)
    B =  reverseDigitsChecker(i)
    if A != B:
        print(i, A, B)

    assert A == B



# class Solution:
#     def reverse(self, x: int) -> int:
        