/* Package shared is used to store the structure of the shared data or structs */
package shared

type Role int

const (
	Candidate = iota
	Leader
	Follower
)
