package cropper

// Cropper type.
type Cropper struct{}

// New returns a new cropper.
func New(todo string) (*Cropper, error) {
	return &Cropper{}, nil
}
