# Микро воркшоп

Репозиторий микро воркшопа по итераторам в Go 1.23

## Материалы

- [Пример итератора](example/example.go)
- Кратко: https://tutorialedge.net/golang/go-123-iterators-tutorial/
- Немного душновато, но стоит почитать: https://bitfieldconsulting.com/posts/iterators

## Упражнение

Приложение выводит в консоль список из 10 пользователей.

Нужно:
1. Перевести метод `printAllUsers` в [файле](objective.go) на итератор.
2. Перевести сделанный итератор на `iter.Pull`

### Запуск

```shell
    make up # Запустить БД

    make run # Заупустить приложение

    make down # Остановить БД
```
