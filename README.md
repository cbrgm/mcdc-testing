# Certified Tester SoSe 18 Praktikum III

## Aufgabenstellung
Die Basis für MC/DC oder auch andere Mehrfachbedingungsüberdeckungen sind Wahrheitswerttabellen, die die
Logik der Bedingung auf Basis der atomaren Bedingungen darstellen. Die Wahrheitswerttabellen werden reduziert
auf boolesche Werte für die atomaren Bedingungen (im Beispiel A1, A2, A3) – unabhängig von ihrer konkreten
Ausdrücken im Code – und dem Resultatwert (im Beispiel B).

### Teil 1
Zwei Eingabevektoren (x1..xn) und (y1..yn) sind Nachbarn, wenn sie sich genau in einer Komponente i
unterscheiden, dh wenn es genau ein i in {1..n} gibt so dass xi = 0 und yi = 1 oder xi = 1 und yi = 0.
Erweitern Sie Ihre Klasse um zwei Methoden:

* isNeighbour soll für 2 Eingabevektoren prüfen ob sie Nachbarn sind
* getNeighbours soll für einen Eingabevektor (x1...xn) alle Nachbarn liefern

### Teil 2
Ergänzen Sie Ihr Programm um eine Methode zur Ermittlung der benötigten Testfälle (als Vektoren der Eingabeparameter mit erwartetem Ergebniswert) für das MC/DC Kriterium
Demonstrieren Sie die Funktion der Methode durch Anwendung auf die Tabelle aus der Folien (s.o.) für Minimal Bestimmte Mehrfachbedingungsüberdeckung

### Teil 4
Ergänzen Sie die Unit-Tests so, dass die Coverage bei mindestens 90% beträgt.

## Installation

```
go build
./mcdc --file <path to matrix file>
```

Input:

```
A1,A2,A3,B

0,0,0,1
1,0,0,0
0,1,0,0
1,1,0,0
0,0,1,1
1,0,1,0
0,1,1,1
1,1,1,0
```

Sample Output:

```
chris@pc:~/code/src/github.com/cbrgm/mcdc$ ./mcdc --file ./data/table2.txt                                                                                     
Testcase #0 : {[true false true false]}
Testcase #1 : {[false false true true]}
Testcase #2 : {[true true true false]}
Testcase #3 : {[true false false false]}
Testcase #4 : {[false true false false]}
Testcase #5 : {[false false false true]}
Testcase #6 : {[false true true true]}
```
