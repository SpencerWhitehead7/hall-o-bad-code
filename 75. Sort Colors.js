// from leetcode medium

// Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.

// Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

// Note: You are not suppose to use the library's sort function for this problem.

// Example:

// Input: [2,0,2,1,1,0]
// Output: [0,0,1,1,2,2]
// Follow up:

// A rather straight forward solution is a two-pass algorithm using counting sort.
// First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.
// Could you come up with a one-pass algorithm using only constant space?

/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
const sortColorsDisgustinglyTrivial = (nums) => {
  nums.sort((a, b) => a - b);
};

const sortColorsTrivial = (nums) => {
  const counter = { 0: 0, 1: 0, 2: 0 };
  nums.forEach((num) => {
    counter[num]++;
  });
  nums.forEach((num, i) => {
    if (counter[0]) {
      nums[i] = 0;
      counter[0]--;
    } else if (counter[1]) {
      nums[i] = 1;
      counter[1]--;
    } else if (counter[2]) {
      nums[i] = 2;
      counter[2]--;
    }
  });
};

const sortColors = (nums) => {
  for (let i = 0; i < nums.length; i++) {
    const num = nums[i];
    if (num === 0) {
      nums.unshift(nums.splice(i, 1)[0]);
    } else if (num === 2) {
      nums.push(nums.splice(i, 1)[0]);
      let counter = 0;
      while (nums[i] === 2 && i + counter < nums.length) {
        nums.push(nums.splice(i, 1)[0]);
        counter++;
      }
      if (i + counter >= nums.length) break;
      i--;
    }
  }
};

// bro I was wiiiiiling lol
