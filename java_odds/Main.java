
/*
Write a function, add,  which adds two large numbers together and returns their sum:
// Returns sum of two numbers
String add(String a, String b)
* Inputs to the add function are two numbers, represented as strings
* Output of the add function is a single string representing the sum of the inputs

* Only positive integers need to be supported
* Input numbers can be very large, with 100 digits or more. This negates the ability to convert the entirety of the strings to integers and simply add them

*/

// Online Java Compiler
// Use this editor to write, compile and run your Java code online

class HelloWorld {
    public static void main(String[] args) {
        
    String num1 = "701,408,733";
    String num2 = "433,494,437";
    
    num1 = removeComas(num1);
     num2 = removeComas(num2);
    System.out.println(addComas(add(num1, num2 )));
    }
    
     public static String removeComas(String num) {
        return num.replace(",", "");
    }
    
    /*     assertAdd("1,234", "1,234", "2,468");
 
            assertAdd("701,408,733", "433,494,437", "1,134,903,170");
            assertAdd("9,234", "1,234", "10,468");
            
            assertAdd("999", "1", "1,000");

            assertAdd("1,999", "1", "1,000");
            assertAdd("1,999", "0,001", "1,000");

            234,123,432,423,423,134,123
    
    */
    
    
    
    public static String add(String num1, String num2){
        
        int len1 = num1.length();
        int len2 = num2.length();
        
        StringBuilder result = new StringBuilder();
        int carry = 0 ;
        
        int max = Math.max(num1.length(), num2.length());
        
        
        while(num1.length() < max){
            num1 = "0"+num1;
        }
        
        while(num2.length() < max){
            num2 = "0"+num2;
        }
        
        for (int i = max-1; i>=0; i--){
            
            int digit1 = num1.charAt(i)- '0';
            int digit2 = num2.charAt(i)- '0';
            int sum = digit1+digit2+carry;
            carry = sum/10;
            
            result.insert(0, sum%10);
        }
        
        if (carry > 0){
            result.insert(0, carry);
        }
        return result.toString();
    }
    
    
    public static String addComas(String num){
        StringBuilder result = new StringBuilder();
        int len = num.length();
        int count = 0;
        for (int i = len-1; i>=0; i--){
            if (count == 3){
                result.insert(0, ",");
                count =0 ;
            }
            result.insert(0, num.charAt(i)- '0');
            count++;
        }
        
        return result.toString();
    }
}