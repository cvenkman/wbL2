# WB Tech: level # 2 (Golang)
## Как делать задания:
Никаких устных решений — только код. Одно решение — один файл с хорошо откомментированным кодом. Каждое решение или невозможность решения надо объяснить.

Разрешается и приветствуется использование любых справочных ресурсов, привлечение сторонних экспертов и т.д. и т.п.


Основной критерий оценки — четкое понимание «как это работает». Некоторые задачи можно решить несколькими способами, в этом случае требуется привести максимально возможное количество вариантов и обосновать наиболее оптимальный из них, если таковой имеется.

Можно задавать вопросы, как по условию задач, так и об их решении.

Для задач на разработку (/develop) все тест-кейсы должны быть оформлены в виде отдельных тестов.




develop

               01
               02    дополнительно - escape
         03          sort - дополнительные флаги -h -M
               04
               05
     06              cut - сделать тесты
07                   Or channel
     08              shell - fork/exec и nc
               09    wget
               10    telnet
11                   HTTP сервер для работы с календарем


https://habr.com/ru/post/457970/

https://stackoverflow.com/questions/6324167/do-browsers-send-r-n-or-n-or-does-it-depend-on-the-browser


1. что такое сокет?

2. post get delete...

3. telnet

4. способы закрытия соединения и горутин

5. потоки и процессы



Если взять очень обобщенно работу сервера, то получается такая последовательность действий:
1. Создать сокет
2. Привязать сокет к сетевому интерфейсу
3. Прослушивать сокет, привязанный к определенному сетевому интерфейсу
4. Принимать входящие соединения
5. Реагировать на события происходящие на сокетах