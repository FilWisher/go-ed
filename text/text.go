/*
Basic Text interface for insertion and deletion

*/

package text

type Text interface {
	Insert(uint, rune) 
	Delete(uint, uint) 
	
	String() []rune
}
