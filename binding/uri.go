package binding

type uriBidning struct{}

func (uriBidning) Name() string {
	return "uri"
}
func (uriBidning) BindURI(m map[string][]string, obj any) error {
	if err := mapURI(obj, m); err != nil {
		return err
	}
	// todo validtor
	return nil
}
