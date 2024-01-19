package com.renovsys;

import java.util.LinkedList;

// Question 1 ->
// custom hash map implementation with get functionality on both key as well as value with O(1)

// Question 2:
// print 1 to 100 with 4 threads in sequence

public class TestPragmaticplay {

    private static final int Max_Count = 100;
    static int count = 1;

    private static final Object lock = new Object();

    public static void main(String[] args) {

        Thread t1 = new Thread(()-> printNumbers(1));
        Thread t2 = new Thread(()-> printNumbers(2));
        Thread t3 = new Thread(()-> printNumbers(3));
        Thread t4 = new Thread(()-> printNumbers(4));

        t1.start();
        t2.start();
        t3.start();
        t4.start();
    }


    private static void printNumbers(int i){
        synchronized (lock) {
            while (count < Max_Count) {
                if (count % 4 == (i - 1)) {
                    System.out.println("Thread Name: "+ i +", count:"+ count);
                    count++;
                    lock.notifyAll();
                    ;
                } else {
                    try {
                        lock.wait();
                    } catch (Exception e) {
                        e.printStackTrace();
                    }
                }
            }
        }
    }
}
