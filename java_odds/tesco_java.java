/*

Problem Statement  
Tesco has a fleet of vehicles to deliver orders to the customer. Assigning the right set of orders to different sized vehicles is crucial 
for efficient delivery of orders. Different vehicle can fit different container sizes. 
 
Given c containers, along with their volumes [l,b,h], catalogue of product with its volume requirement (l,b,h) and an order 
with p products and its quantity. 
 
Example: 
Containers:  
SMALL -> id=1, length=10, breadth=20, height=30 
MEDIUM -> id=2, length=50, breadth=60, height=70 
LARGE -> id=3, length=100, breadth=200, height=300 
 
Product: 
productId=1, length=2, breadth=4, height=10 
productId=2, length=10, breadth=30, height=4 
productId=3, length=5, breadth=6, height=7 
 
Order: 
productId=1, quantity=3 
productId=3, quantity=7 
 
 
Determine if that order fits in any of the given c containers and return the ID of the container that can be used. 
For the above sample of example SMALL container with id=1 should be returned. 
has context menu


container size fixed - 3
products -> n products
Order -> list<products> -> right container for that order

*/

import java.util.*;
import java.lang.*;
import java.io.*;

// The main method must be in a class named "Main".

class Dimensions{

    int len;
    int wid;
    int hgt;
    double vol;

    public Dimensions(int len, int wid, int hgt){
        this.len = len;
        this.wid = wid;
        this.hgt = hgt;
        this.vol = len*wid*hgt;
    }
}

class Container extends Dimensions{
    String contianerId;

    public Container(String id, int len, int wid, int hgt){
        super(len, wid, hgt);
        this.contianerId = id;
        
    }
    
}

class Product extends Dimensions{
    String productId;

    public Product(String id, int len, int wid, int hgt){
        super(len, wid, hgt);
        this.productId = id;
        
    }
    
}


public class Main {

    static List<Container> containers = new ArrayList<>();
        
    public static void main(String[] args) {
        System.out.println("Hello world!");
        containers.add(new Container("3", 100, 200,300));
        containers.add(new Container("2", 50, 60,70));
        containers.add(new Container("1", 10, 20,30));

        Arrays.sort(containers, Comparator.comparingDouble(c -> c.contianerId)); Omlog(m)

        List<Product> orders = new ArrayList<>();

        orders.add(new Product("1", 2, 4, 10 ));
            orders.add(new Product("2", 10, 30, 4 ));
            orders.add(new Product("3", 5, 6, 7 ));
            orders.add(new Product("4", 20, 40, 2 ));

        System.out.println(findtheRightContainerForOrder(orders));


    }
    public static String findtheRightContainerForOrder(List<Product> orders){ // N + 3 ~ O(N)

        int maxVol = 0;
        String foundContainerID = "";
        for (Product p : orders){
            maxVol += p.vol;
        }

        for (Container c : containers){
            if (maxVol <= c.vol){
                foundContainerID = c.contianerId;
                    break;
            }
        }

        return foundContainerID;
    }
}