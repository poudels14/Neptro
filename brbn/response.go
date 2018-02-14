package brbn

type Serializable string

type Response struct {
	status int
	data   Serializable
}
