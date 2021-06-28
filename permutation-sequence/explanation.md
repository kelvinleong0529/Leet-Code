
To make things simpler, we start index everything from 0

Theory:

(1) We begin from the left-most digit

(2) We create a ***list*** , that contains all digit from 1 to ***n***

(2) We split the possible permutations into different groups according to their 'leading' digits

(3) We discover which group (n*th*) the final answer lies within

(4) We remove the n*th* element in the ***list***, updates ***k*** ( ***k=k%(len(list)-1)!*** )

(5) Repeat steps 3 - 5 until ***list*** is empty
	
***Sample Case: n=4, k=15 (0-indexed)***	
	
### Stage 1
	
	List = [1,2,3,4]
	Group 0: 1234,1243,1324,1342,1423,1432 (where k=0~5) 
	Group 1: 2134,2143,2314,2341,2413,2431 (where k=6~11)
	Group 2: 3124,3142,3214,3241,3412,3421 (where k=12~17)
	Group 3: 4123,4132,4213,4231,4312,4321 (where k=18~23)
	
we can confirm that the final answer lies within **Group [int(k/(len(List)-1)!)] = 2**

since we know the answer is in Group 2, meaning the 1st digit of the answer is '3' (equivalent to the value of ***List[2]*** ), answer = **3 _ _ _**

we remove ***List[3]***  , update **k=k%(len(List)-1)!=3**

### Stage 2

	List = [1,2,4]
	Group 0: 124,142 (where k=0~1)
	Group 1: 214,241 (where k=2~3)
	Group 2: 412,421 (where k=4~5)

we can decude the answer lies within **Group [int(k/(len(List)-1)!)] = 1**

we can confirm that the 2nd digit is '2' (equivalent to the value of ***List[1]*** ), answer = **3 2 _ _**

remove ***List[2]*** , update **k = k%(len(List)-1)!=1**

### Stage 3

	List = [1,4]
	Group 0: 14 (where k=0)
	Group 1: 41 (where k=1)

Answer lies within **Group [int(k/(len(List)-1!)] = 1**, 3rd digit is ***List[1]*** , which is '4', answer = **3 2 4** _

remove ***List[1]*** , update **k = k%(len(List)-1)! = 0**
	
### Stage 4
	
	List = [1]
	Group 0: 1 (where k=0)
	
Answer lies within **Group [int(k/(len(List)-1!)] = 0**, 4th digit is ***List[0]*** , which is '1', answer = **3 2 4 1**

After we remove ***List[0]*** , ***list*** becomes empty, terminate the loop
