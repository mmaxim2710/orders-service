# orders-service

-----------

## Список компонентов


### Сервис авторизации
**API**:
- Авторизация - done
  - входные параметры: login, пароль
  - выходные параметры: идентификатор пользователя, JWT token, refresh token, expiration time
- Создание нового пользователя - done
    - входные параметры: login, пароль, имя, фамилия, email
    - выходные параметры: ID, email, login, first name, last name
- Изменение пользователя - done
  - Входные параметры: ID, имя, фамилия, email
  - Выходные параметры: ID, email, login, first name, last name
- Удаление пользователя - done
  - Входные параметры: идентификатор пользователя
  - Выходные параметры: данные удалённого пользователя
- Поиск пользователя по фильтрам - id
  - входные параметры:  фильтры для поиска (login, маска фамилии, маска имени) **При указании логина поиск будет производиться только по нему**
  - выходные параметры: login, имя, фамилия, email

### Сервис услуг
**API**:
- Создание услуги - done
    - Входные параметры: идентификатор пользователя, название услуги, параметры услуги, дата создания
    - Выходыне параметры: идентификатор услуги
- Изменение услуги - done
    - Входные параметры: идентификатор услуги, идентификатор пользователя, параметры услуги
    - Выходыне параметры: нет
- Удаление услуги - done
    - Входные параметры: идентификатор услуги, идентификатор пользователя
    - Выходыне параметры: нет
- Получение списка услуг по фильтрам: - by serviceID, by userID
    - Входные параметры: фильтры для поиска
    - Выходные параметры: массив с услугами, где для каждого указаны его идентификатор, название, параметры, автор и дата создания

### Сервис заказов
**API**:
- Создание заказа
    - Входные параметры: идентификатор пользователя, массив из идентификаторов услуг, дата создания
    - Выходные параметры: идентификатор заказа
- Изменение заказа:
    - Входные параметры: идентификатор пользователя, идентификатор заказа, массив из идентификаторов услуг
    - Выходные параметры: нет
- Удаление заказа
    - Входнае параметры: идентификатор пользователя, идентификатор заказа
    - Выходные парамтеры: нет
- Получение заказа по идентификатору пользователя
    - Входные параметры: идентификатор пользователя
    - Выходные параметры: массив из идентификаторов заказа
- Получение информации о заказе по идентификатору заказа
    - Входные параметры: идентификатор заказа
    - Выходные параметры: массив из идентификаторов услуг, дата создания