package tags

type (
	Generable interface {
		Generate() (string, string)
	}
)

func Tags(generables ...Generable) map[string]string {
	response := make(map[string]string)
	for _, tagGen := range generables {
		t, c := tagGen.Generate()
		response[t] = c
	}

	return response
}
