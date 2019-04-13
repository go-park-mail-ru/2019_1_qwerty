package helpers

// ErroRouter - Разбор ошибок
// func ErroRouter(w *http.ResponseWriter, resp *models.User, err *error, statusCode int) {
// 	switch *err {
// 	case nil:
// 		json.NewEncoder(w).Encode(*resp)
// 		w.WriteHeader(statusCode)
// 	case models.EUserNE,
// 		models.ENE:
// 		json.NewEncoder(w).Encode(*err)
// 		w.WriteHeader(http.StatusNotFound)
// 	case models.EUserAE:
// 		json.NewEncoder(w).Encode(*resp)
// 		w.WriteHeader(http.StatusConflict)
// 	default:
// 		json.NewEncoder(w).Encode((*err).Error())
// 		w.WriteHeader(http.StatusBadRequest)
// 	}
// }
