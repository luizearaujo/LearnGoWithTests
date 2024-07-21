package inteiros

import "testing"
import "fmt"

func TestAdicionar(t *testing.T){
	
	soma := Adicionar(2, 2)
	esperado := 4
	
	if soma != esperado {
		t.Errorf("esperado '%d', resultado '%d'", esperado, soma)
	}
}

func ExampleAdiciona() {
	soma := Adicionar(1, 5)
	fmt.Println(soma)
	// Output: 6
}
