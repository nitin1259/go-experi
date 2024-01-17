package com.renovsys;

import java.util.Arrays;

public class ShiftNegativePositive {
    public static void main(String[] args) {
        // Input: -12, 11, -13, -5, 6, -7, 5, -3, -6
        // Output: -12 -13 -5 -7 -3 -6 11 6 5
        Integer [] input = {-12, 11, -13, -5, 6, -7, 5, -3, -6};
        shiftNegativePositive(input);
        System.out.println(input);

        for (Integer i : input){
            System.out.print(i);
        }
    }

    public static void shiftNegativePositive(Integer []input){
        Arrays.sort(input, (a, b)->{
            // int a = Integer.parseInt(A)
            if (a < 0 && b < 0) {
                return 0;
            } else if (a < 0 && b >= 0) {
                return -1;
            } else if (a >= 0 && b < 0){
                return 1;
            }else{
                return 0;
            }
        });
    }

}