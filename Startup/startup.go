/*
* After startup the control-unit enlists all attached slaves.
 */

package startup

/* Slice with saved IDs */
var id []int

/*
* @ brief: Auto-Enlists attached slaves
* @ param: number of attached slaves
* @ return: Hashmap of slave number and id, error
 */
func Setup() error {
	return nil
}

/*
@ brief: Adds Slave to Hashmap
@ param: id
@ return: error
*/

func AddId(new_id int) error {
	id = append(id, new_id)
	return nil
}
