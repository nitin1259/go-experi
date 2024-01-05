
/**
 * 
 * The task is to find the length of the longest subsequence in a given array of integers such that all elements of the subsequence are sorted in strictly ascending order. This is called the Longest Increasing Subsequence (LIS) problem.

For example, the length of the LIS for
is since the longest increasing subsequence is

[1, 1 1 1 0 ]
[5,2,7,3,4,8]

Here's a great YouTube video of a lecture from MIT's Open-CourseWare covering the topic. 
 * 
 */

public class LongestIncrSubSeq{


  public static int LenOFLongesSubseq(int[]nums){

    if (nums==null || nums.length==0){
      return 0;
    }

    int n = nums.length;

    int[] table = new int[n];

    Arrays.fill(table, 1); 


    for (int i = 1; i<n; i++){
      for (int j = 0; j<i; j++){
          if (nums[i]> nums[i]){
            table[i] = Math.max(table[i], table[j]+1);
          }
      }
    }

    int maxLen = 0;

    for (int l : table){
      maxLen = Math.max(maxLen, l);
    }

    return maxLen;

  }



}