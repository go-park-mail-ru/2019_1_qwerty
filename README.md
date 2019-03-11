# 2019_1_qwerty

## Состав команды
* [Бобров Константин](https://github.com/KostyaBobroff);
* [Векшин Роман](https://github.com/BrBrRoman);
* [Захаров Дмитрий](https://github.com/goddeuce1);
* [Ражев Михаил](https://github.com/Lunex08).

 Ментор:
 * [Юрий Байков](https://github.com/OkciD).

 ## [Репозиторий фронтенда](https://github.com/frontend-park-mail-ru/2019_1_qwerty/)


## API

### **/api/user/**
* **/ GET:**
Вернуть информацию о пользователе
```CODE: 200, 404```
* **signup/ POST|OPTIONS:**  
Зарегестрироваться  
```CODE: 200, 404```
* **login/ POST|OPTIONS:**  
Начать сессию   
```CODE: 200, 404```
* **check/ GET:**  
Проверить сессию  
```CODE: 200, 404```
* **logout/ GET:**  
Завершить сессию    
```CODE: 200, 404```
* **avatar/ POST | OPTIONS:**
Загрузить на сервер аватар для текущего пользователя  
```CODE: 200, 403```  
* **update/ POST | OPTIONS:**
Обновление данных пользователя
```CODE: 200, 403```  

### **/api/score**
* **POST:**  
___?score=[number]___ Сохранить результат для текущего пользователя  
```CODE: 200, 404```
* **GET:**  
___?offset=[number]___ Вернуть таблицу результатов начиная с позиции offset (default: 0)  
```CODE: 200, 404```
* **PUT:**  
Неприменимо  
```CODE: 405```
* **DELETE:**  
Неприменимо    
```CODE: 405```
