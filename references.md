# References
## Link of concepts
--- 

I'm going to put some link of discussions and concepts in this file just to register and save.

### Discussion about how pass variable. By reference or by value.
[Golang pass by pointer vs pass by value](https://goinbigdata.com/golang-pass-by-pointer-vs-pass-by-value/)
> "Sometimes the choice of how to pass variable is predetermined by variable type or its usage. Otherwise, it’s highly recommended to pass variables by value."

To read this article I was questioning why using a reference to the struct type variable 
whenever passing it to a function instead of directly passing the variable. So I was reading the forum on the udemy course and the instructor answer this: 

> It depends on the function, but it's often more efficient to pass around pointers than entire structs.

And give the reference of this article. 