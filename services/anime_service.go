package services
import "anime-backend/models"
func GetAnimeList() []models.Anime {
	return []models.Anime{
		{ID: "1", Title: "Naruto", Synopsis: "A story about a young ninja."},
		{ID: "2", Title: "Attack on Titan", Synopsis: "Humanity fights against Titans."},
		{ID: "3", Title: "One Piece", Synopsis: "Pirates seek the legendary treasure."},
		{ID: "4", Title: "My Hero Academia", Synopsis: "Heroes with unique powers fight villains."},
	}
}
func GetAnimeByID(id string) *models.Anime {
	animeList := GetAnimeList()
	for _, anime := range animeList {
		if anime.ID == id {
			return &anime
		}
	}
	return nil
}
