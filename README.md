# WB Tech: level # 2 (Golang)
## Как делать задания:
Никаких устных решений — только код. Одно решение — один файл с хорошо откомментированным кодом. Каждое решение или невозможность решения надо объяснить.

Разрешается и приветствуется использование любых справочных ресурсов, привлечение сторонних экспертов и т.д. и т.п.


Основной критерий оценки — четкое понимание «как это работает». Некоторые задачи можно решить несколькими способами, в этом случае требуется привести максимально возможное количество вариантов и обосновать наиболее оптимальный из них, если таковой имеется.

Можно задавать вопросы, как по условию задач, так и об их решении.

Для задач на разработку (/develop) все тест-кейсы должны быть оформлены в виде отдельных тестов.




develop

            01
            02
    03
            04  анаграмы
            05
06
07
        08
09
10
11

### Паттерны

выделить прееменные составляющие и инкапсулируйте их, чтобы позднее их можно было изменять или расщирять без воздействия на постоянные составяляющие.

все паттерны обеспечивают возможность изменений некоторой части системы независимо от других частей.



### Прадигмы
Способ программирования, независящий от языка.
Совокупность идей и понятий, определяющих стиль написания компьютерных программ (подход к программированию).

* Структурное программирование

Любая программа, которая строится без использования оператора goto, состоит из трёх базовых управляющих конструкций: последовательность, ветвление, цикл; кроме того, используются подпрограммы. При этом разработка программы ведётся пошагово, методом «сверху вниз».

* Императивное программирование

Что мы делаем

C#, Ruby, Java, C++, Python

Программный код в императивном стиле организован как последовательность отдельных команд, инструкций, описывающих логику работы программы. Читая такой код, можно понять, каким образом будет меняться состояние приложения в тот или иной момент — в зависимости от того, какие фрагменты кода будут запущены.

Подвиды:

  - Процедурное программирование

Алгоритм выполнения программы представлен как последовательность инструкций, которые организованы в специальные блоки кода, процедуры (подпрограммы), которые можно вызывать много раз из любой точки программы.

Легко разобраться, с ростом проекта сложно поддерживать и масштабировать.

  - Объектно-ориентированное программирование

Разделение на объекты и классы.

* объект — это элементарная сущность, имеющая свойства (атрибуты) и поведение (методы, они же — бывшие процедуры);
* класс — это тип, шаблон, определяющий структуру, набор атрибутов и методов для объектов одного типа — то есть, экземпляров класса;
класс может наследовать атрибуты и методы своего родительского класса и иметь при этом свои собственные. Так формируется иерархия классов, она позволяет моделировать предметную область на разных уровнях абстракции и детализации, решая задачу по частям;
* полиморфизм — это механизм, позволяющий использовать одну и ту же функцию (метод) для данных разных типов (классов);
* инкапсуляция — это механизм, позволяющий скрывать некоторые детали реализации внутри класса. Часто сторонним сущностям, которые работают с объектом ни к чему знать нюансы реализации его класса и иметь доступ к каким-то его атрибутам и методам. Поэтому часто разработчик создает для класса интерфейс, которые отвечает за взаимодействие с внешними сущностями, открывая специально выбранные для этого атрибуты и методы.

* Декларативное

Что мы хотим

LISP, SQL, Haskell, Scala

Подвиды:

  - Функциональное программирование
  - Логическое программирование



