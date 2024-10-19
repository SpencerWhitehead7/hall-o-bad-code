# I remember absolutely no details of this, but I think this was the first time I really felt the difference 
# between a fast implementation and a slow implementation and the importance of memory allocation/layout
# all specific memories of writing it and indeed of the underlying python
# are gone, but the shock at the difference in speed remains

import math
import time

def sieve_rec(number):
  def filterer(numbers, startIndex):
    if startIndex == len(numbers) - 1:
      return list(map(lambda x: x[1], numbers))
    else:
      selector = numbers[startIndex][1]
      filtered = list(filter(lambda x: x[1] % selector != 0 or x[1] == selector, numbers))
      print("finished a pass")
      return filterer(filtered, startIndex + 1)

  numbers = list(enumerate(range(2, number + 1)))
  return filterer(numbers, 0)

def sieve_iter(number):
  numbers = list(enumerate(range(2, number + 1)))
  start_index = 0
  while(True):
    if start_index == len(numbers) - 1:
      return list(map(lambda x: x[1], numbers))
    else:
      selector = numbers[start_index][1]
      numbers = list(filter(lambda x: x[1] % selector != 0 or x[1] == selector, numbers))
      start_index += 1

def sieve_iter_2(number):
  numbers = list(range(0, number + 1))
  numbers[1] = 0
  start_index = 2
  while(True):
    if start_index == math.ceil(math.sqrt(len(numbers))):
      return list(filter(lambda x: x, numbers))
    else:
      selector = start_index * start_index
      while(selector < len(numbers)):
        numbers[selector] = 0
        selector += start_index
      start_index += 1

def prime_factorization(number):
  res = []
  primes = sieve_iter_2(math.ceil(math.sqrt(number)))
  while(number != 1):
    for x in primes:
      if number // x == number / x:
        res.append(x)
        number /= x
        break
    else:
      return [number]
  return res

# start = time.time()
# print(prime_factorization(600851475143)) # [71, 839, 1471, 6857]
# print(prime_factorization(6008514751430)) # ???
# print(prime_factorization(2*3*5*7*11*13*17*19*23*29*31))
# print(prime_factorization(29))
# end = time.time()
# print(end - start)
# print(math.ceil(math.sqrt(600851475143)))

# truePrimes = set([2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97])
# myPrimes = sieve_iter_2(100)
# myPrimesSet = set(myPrimes)
# print(len(myPrimes) == len(myPrimesSet) and len(truePrimes ^ myPrimesSet) == 0)
