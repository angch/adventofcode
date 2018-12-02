Q2 Part 2:
-----------------------------------------

So this time you want to find the "highest matched chars" between two strings.
The order of the chars are important.

   0 1 2 3 4
  ----------
0 |a b c d e -> row1
1 |f g h i j -> row2
2 |k l m n o
3 |p q r s t
4 |f g u i j
5 |a x c y e
6 |w v x y z

We start from row1 and then keep changing row2.
When row1, and row2 are same you can skip it.
We cache the length of the highest pair of rows.
As new higher rows are matched, we override old ones.

Note 1:
Since equality is bi-directional ways. match(row1, row2)===match(row2, row1),
when we are checking (row1,row2) we can cache (row2,row) and skip it when we 
meet it in the subsequent loop.

Note 2:
We can also do this question by directly eliminating the chars. But we choose
this way as an exercise.


Q2 Part 1:
-----------------------------------------
We want to find out chars with the count 2, and the count of 3
Just cache previous chars, and check if new ones are matching previous ones.

Note: If you terminate too early when checking the count, you will miss later 
chars in the string and cause false positives in the final result.

In the code, I'm flipping the array(hash) to eliminate redundant chars 
on the final section.
Ex;

a => 2
b => 2
c => 3

to

2 => b
3 => c


Q1 Part 2:
-----------------------------------------
Still the same idea like Q1 Part 1. 
This time you don't depend on the length of the input.
Cache previous frequencies. 
Wrap the summation around a infinite while loop, check if the frequence exist after summation.


Q1 Part 1:
-----------------------------------------
The frequency is calculating by arithmetic operation. 
You can just sum them even if its negative.
The frequency will fluctuate between >+0, 0, and <-0
