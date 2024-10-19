# an algo for finding the max number of identical substrings into which a string can be broken?

import math

class solution:
  def solution(cakeStr):
    cakeStrLen = len(cakeStr)
    sliceSizesCandidates = range(1, math.floor(cakeStrLen / 2) + 1)
    for sliceSize in sliceSizesCandidates:
      if cakeStrLen % sliceSize == 0:
        slice = cakeStr[:sliceSize]
        sliceCount = int(cakeStrLen / sliceSize)
        if slice * sliceCount == cakeStr:
          return sliceCount

    return 1

solution = solution()

print(solution.solution("a" * 3)) # 3
print(solution.solution("ab" * 7)) # 7
print(solution.solution("ab" * 13 + "c")) # 1
print(solution.solution("abcdefghi" * 13 + "c")) # 1

print(solution.solution("abccbaabccba")) # 2
print(solution.solution("abcabcabcabc")) # 4
solution.solution("abcabcabcabc")
