
# GO-RELOADED
* `ayessenb` 


## First project in go lang testing period

I created the program which receives as arguments the name of a file containing a text that needs some modifications (the input) and the name of the file the modified text should be placed in (the output). Next is a list of possible modifications that my program should execute:

* go.mod - defines a module, which is a collection of Go packages. It contains the data about the module, such as its name, version, and dependencies on other modules.
* go.sum - contain a list of all the downloaded modules, their versions, and their hash.
* main.go - go-reloaded main function which executes all the functions related all
list of modifications.

#### Modificators:


* Every instance of (hex) should replace the word before with the decimal version of the word (in this case the word will always be a hexadecimal number, in case of not valid number it should not change string and remove the "(hex)" command). (Ex: "1E (hex) files were added" -> "30 files were added") (Ex: "hello (hex) files were added" -> "hello files were added" and the ERROR message in the console will be printed: "ERROR: invalid hex") 

* Every instance of (bin) should replace the word before with the decimal version of the word (in this case the word will always be a binary number, in case of not valid number it should not change string and remove the "(bin)" command). (Ex: "It has been 10 (bin) years" -> "It has been 2 years") (Ex: "It has been good (bin) years" -> "It has been good years" and the ERROR message int he console will be printed: "ERROR: invalid bin" )

* Every instance of (up) converts the word before with the Uppercase version of it. (Ex: "Ready, set, go (up) !" -> "Ready, set, GO !") this command will work within single quotes, brackets () and even can affect the strings outside these quotes and brackets. (Ex: "hello ((up) how are    you? ((up, 3)))" -> "HELLO( HOW ARE YOU?())")

* Every instance of (low) converts the word before with the Lowercase version of it. (Ex: "I should stop SHOUTING (low)" -> "I should stop shouting")

* Every instance of (cap) converts the word before with the capitalized version of it. (Ex: "Welcome to the Brooklyn bridge (cap)" -> "Welcome to the Brooklyn Bridge")
        
    
* For (low), (up), (cap) if a number appears next to it, like so: (low, \<number\>) it turns the previously specified number of words in lowercase, uppercase or capitalized accordingly. (Ex: "This is so exciting (up, 2)" -> "This is SO EXCITING")
    * if the numeric number is negative - it gives the ERROR message and keeps the command as it is
    * if the numeric number is bigger than the number of string before the command - it applies to all and removes the command
    * if the numeric number is 0 - it doesn't change anything and removed the command
    * if anything else if provided instead of numeric number and closing bracket it will give and ERROR message and will keep it as it was provided (Ex: "Hello, whats up? (  up, blahblah   )" ->  "Hello, whats up? (up, blahblah)")

* Every instance of the punctuations ., ,, !, ?, : and ; should be close to the previous word and with space apart from the next one. (Ex: "I was sitting over there ,and then BAMM !!" -> "I was sitting over there, and then BAMM!!").
    * Except if there are groups of punctuation like: ... or !?. In this case the program should format the text as in the following example: "I was thinking ... You were right" -> "I was thinking... You were right".

    * The punctuation mark ' will always be found with another instance of it and they should be placed to the right and left of the word in the middle of them, without any spaces. (Ex: "I am exactly how they describe me: ' awesome '" -> "I am exactly how they describe me: 'awesome'")

    * If there are more than one word between the two ' ' marks, the program should place the marks next to the corresponding words (Ex: "As Elton John said: ' I am the most well-known homosexual in the world '" -> "As Elton John said: 'I am the most well-known homosexual in the world'")

    * Every instance of a should be turned into an if the next word begins with a vowel (a, e, i, o, u) or a h. (Ex: "There it was. A amazing rock!" -> "There it was. An amazing rock!").


#### Improved skills:
* go lang programming skills 
* usage of string, strconv libraries of Go
* conversion of different numbers
* usage of many functions and files

## Usage/Examples
Cloning storage to your host
```CMD/Terminal 
git clone git@github.com:compygirl/go-reloaded.git
```
Go to the downloaded repository:

```CMD/Terminal 
cd go-reloaded
```
Run a program:
```CMD/Terminal 
go run main.go sample.txt result.txt
```
Output:
```
inside the result.txt files
additionally the ERROR messages on consolse
```

## Feedback

If you liked our project, we would be grateful if you could add `Star` to the repository.

Alem Student
25.04.2023.