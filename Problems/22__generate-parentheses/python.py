#!/usr/local/bin/python3



# Performance: O(2^(2n))
# Space: O(2^(2n))
# Runtime (n=13) 5.810s
# Runtime (n=14) 20.819s
def generateParenthesisRecursive1(n):
    maxSize = n*2
    def gp(parens, left, right):
        if left+right >= maxSize:
            print(parens)
            return
        if left < n:
            gp(parens+'(', left+1, right)
        if right < left:
            gp(parens+')', left, right+1)
    
    gp('', 0, 0)


# Performance: O(2^(2n))
# Space: O(n)
# Runtime (n=13) 5.851s
# Runtime (n=14) 19.417s
def generateParenthesisRecursive2(n):
    parens = [' '] * (n*2)

    def gp(left, right):
        if left+right >= len(parens):
            print(''.join(parens))
            return
        if left < n:
            parens[left+right] = '('
            gp(left+1, right)
        if right < left:
            parens[left+right] = ')'
            gp(left, right+1)
    
    gp(0, 0)

generateParenthesisRecursive2(13)

# Performance: O(2^(2n))
# Space: O(2^(2n))
# Runtime (n=13) 6.068s
# Runtime (n=14) 18.008s
def generateParenthesisIterative1(n):
    maxSize = n*2
    stack = [('', 0, 0)]

    while len(stack) > 0:
        parens, left, right = stack.pop()
        if len(parens) >= maxSize:
            print(parens)
            continue
        if left < n:
            stack.append((parens+'(', left+1, right))
        if right < left:
            stack.append((parens+')', left, right+1))

