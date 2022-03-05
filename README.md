# ascii-art-web

The ascii art project, modified to work in a web page.

Created by Nikolay, Remi, and Maya.

A web application written in Go, HTML, and CSS which allows users to generate ascii-art text based on their inputs. The algorithm works by creating a map where the key is a rune (any character in the ascii table (starting from " " until ~) and the value is a slice of strings where each element of the slice is one line of the eight-line ascii character that is read from the banner file. On the client side, if a POST request is sent to the /ascii-art route, the user's input will be accessed by the server and sent to the "mapsimple.go" file in the "functions" folder, where it will be processed. If there are no errors, the "MakeMapSimple" function will return ascii art based on the users input, then sent to an HTML template to be displayed on the front end. If there are errors, the web application will return the correct ones accordingly. I.e. 404 for an unknown route, 400 for a bad request or 500 for an internal server error.

Execute "go run ." in the terminal and then load up http://localhost:8080 in your browser. Upon loading, select a text style you would like your ASCII art to have from the radio buttons, and type the text you would like to turn in to ASCII art in the textbox below. Select the submit button once you are happy with your choices to see your artwork.

To test the code, change into the test directory and run the command "go test" or "go test ." in the terminal. You do not have to run the server to do this. If successful, the result "ok" shall be printed in the terminal.

The currently available styles are: Standard, Shadow, and Tinkertoy.