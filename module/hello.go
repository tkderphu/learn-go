
package main
import "fmt"
import "github.com/rs/zerolog/log"
func main() {
	fmt.Println("hello world")
	log.Info().Msg("test")
}