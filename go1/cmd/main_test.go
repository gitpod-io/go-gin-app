package main

import (
	"fmt"
	"testing"
)


func TestMain(t *testing.T)  {
	*phrase = "div"
	*urls = "climaticarus.ru"
	*file = ""
	c := worker(*phrase, *urls, *file)
	fmt.Println(c)


	if result[worker(*phrase, *urls, *file) != 422{
		t.Error("Найдено совпадений", test1.count)
	}else{
		fmt.Println("Все ОК")
	}
}