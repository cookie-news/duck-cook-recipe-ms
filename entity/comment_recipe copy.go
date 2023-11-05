package entity

type LikeRecipe struct {
	IdLike string `json:"id,omitempty"`
	Recipe `json:"-"`
}
