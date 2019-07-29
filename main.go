package main

func sum(a,b int)int{
	return a+b
}

func CallNoDefer(){
	sum(10,20)
}

func CallDefer(){
	defer sum(10,20)
}