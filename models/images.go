package image

type Image struct {
	models.Model
	Title    string `json:"title" bson:"title"`
	Caption  string `json:"caption" bson:"caption"`
	FilePath string `json:"file_path" bson:"file_path"`
}
